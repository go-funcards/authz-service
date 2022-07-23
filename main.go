package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-funcards/authz-service/internal/authz/db"
	"github.com/go-funcards/authz-service/internal/config"
	server "github.com/go-funcards/authz-service/internal/v1"
	"github.com/go-funcards/authz-service/proto/v1"
	"github.com/go-funcards/grpc-server"
	"github.com/go-funcards/logger"
	"github.com/go-funcards/mongodb"
	"github.com/go-funcards/validate"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
)

//go:generate sh genproto.sh

const envConfigFile = "CONFIG_FILE"

var (
	version    string
	buildDate  string
	buildTime  string
	configFile string
)

func init() {
	configFile = os.Getenv(envConfigFile)
	if configFile == "" {
		flag.StringVar(&configFile, "c", "config.yaml", "application config path")
		flag.Parse()
	}
}

func main() {
	ctx := context.Background()

	log := logger.GetLog().WithFields(logrus.Fields{
		"service": os.Args[0],
		"version": fmt.Sprintf("%s.%s.%s", version, buildDate, buildTime),
	})

	cfg := config.GetConfig(configFile, log)

	validate.Default.RegisterStructRules(cfg.Validation.Rules, []any{
		v1.IsGrantedRequest{},
		v1.SaveDefsRequest_Def{},
		v1.SaveDefsRequest{},
		v1.DeleteDefsRequest{},
		v1.SaveRulesRequest_Rule{},
		v1.SaveRulesRequest{},
		v1.SaveSubRequest_Ref{},
		v1.SaveSubRequest{},
		v1.DeleteSubRequest{},
		v1.DeleteRefRequest{},
		v1.SubRequest{},
	}...)

	log.Debug("initializing mongodb")
	mongoDB := mongodb.GetDB(ctx, cfg.MongoDB.URI, log)

	log.Debug("initializing definition storage")
	def := db.NewDefStorage(ctx, mongoDB, log)

	log.Debug("initializing rule storage")
	rule := db.NewRuleStorage(ctx, mongoDB, log)

	log.Debug("initializing subject storage")
	sub := db.NewSubStorage(ctx, mongoDB, log)

	log.Debug("create casbin factory")
	factory := cfg.Casbin.EnforcerFactory(def, rule, sub)

	register := func(srv *grpc.Server) {
		v1.RegisterAuthorizationCheckerServer(srv, server.NewCheckerServer(factory))
		v1.RegisterDefinitionServer(srv, server.NewDefServer(def))
		v1.RegisterRuleServer(srv, server.NewRuleServer(rule))
		v1.RegisterSubjectServer(srv, server.NewSubServer(sub))
	}

	lis, err := net.Listen("tcp", cfg.GRPC.Addr)
	if err != nil {
		log.WithField("error", err).Fatal("failed to create tcp listener")
	}

	log.Infof("bind application to addr: %s", lis.Addr().(*net.TCPAddr).String())

	grpcserver.Start(ctx, lis, register, log, grpc.ChainUnaryInterceptor(
		grpc_recovery.UnaryServerInterceptor(),
		grpc_logrus.UnaryServerInterceptor(log),
		mongodb.ErrorUnaryServerInterceptor(),
		validate.DefaultValidatorUnaryServerInterceptor(),
	))
}
