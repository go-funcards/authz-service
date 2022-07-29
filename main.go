package main

import (
	"context"
	"flag"
	"github.com/go-funcards/authz-service/internal/authz/db"
	"github.com/go-funcards/authz-service/internal/config"
	server "github.com/go-funcards/authz-service/internal/v1"
	"github.com/go-funcards/authz-service/proto/v1"
	"github.com/go-funcards/grpc-server"
	"github.com/go-funcards/grpc-server/grpc_middleware/recovery"
	"github.com/go-funcards/mongodb"
	"github.com/go-funcards/validate"
	"github.com/jwreagor/grpc-zerolog"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"io"
	"net"
	"os"
	"time"
)

//go:generate sh genproto.sh

const (
	envConfigFile = "CONFIG_FILE"
	envLogLevel   = "LOG_LEVEL"
	envLogPretty  = "LOG_PRETTY"
)

var (
	version     string
	configFile  string
	logLevelStr string
	logLevel    zerolog.Level
	logOutput   io.Writer
)

func init() {
	flag.StringVar(&configFile, "c", "config.yaml", "application config path")
	flag.StringVar(&logLevelStr, "log-level", "info", "application log level")
	flag.Parse()

	if os.Getenv(envConfigFile) != "" {
		configFile = os.Getenv(envConfigFile)
	}

	if os.Getenv(envLogLevel) != "" {
		logLevelStr = os.Getenv(envLogLevel)
	}
	logLevel, _ = zerolog.ParseLevel(logLevelStr)
	if zerolog.NoLevel == logLevel {
		logLevel = zerolog.InfoLevel
	}

	logOutput = os.Stdout
	if os.Getenv(envLogPretty) != "" {
		logOutput = zerolog.ConsoleWriter{Out: logOutput}
	}

	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.TimestampFieldName = "timestamp"
	zerolog.LevelFieldName = "severity"
}

func main() {
	ctx := context.Background()

	log := zerolog.
		New(logOutput).
		Level(logLevel).
		With().
		Caller().
		Timestamp().
		Str("system", "grpc").
		Str("span.kind", "server").
		Str("server.name", os.Args[0]).
		Str("server.version", version).
		Logger()

	grpclog.SetLoggerV2(grpczerolog.New(log))

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

	log.Debug().Msg("initializing mongodb")
	mongoDB := mongodb.GetDB(ctx, cfg.MongoDB.URI, log)

	log.Debug().Msg("initializing definition storage")
	def := db.NewDefStorage(ctx, mongoDB, log)

	log.Debug().Msg("initializing rule storage")
	rule := db.NewRuleStorage(ctx, mongoDB, log)

	log.Debug().Msg("initializing subject storage")
	sub := db.NewSubStorage(ctx, mongoDB, log)

	log.Debug().Msg("create casbin factory")
	factory := cfg.Casbin.EnforcerFactory(def, rule, sub)

	register := func(srv *grpc.Server) {
		v1.RegisterAuthorizationCheckerServer(srv, server.NewCheckerServer(factory))
		v1.RegisterDefinitionServer(srv, server.NewDefServer(def))
		v1.RegisterRuleServer(srv, server.NewRuleServer(rule))
		v1.RegisterSubjectServer(srv, server.NewSubServer(sub))
	}

	lis, err := net.Listen("tcp", cfg.GRPC.Addr)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create tcp listener")
	}

	log.Info().Msgf("bind application to addr: %s", lis.Addr().(*net.TCPAddr).String())

	grpcserver.Start(ctx, lis, register, log, grpc.ChainUnaryInterceptor(
		mongodb.ErrorUnaryServerInterceptor(),
		validate.DefaultValidatorUnaryServerInterceptor(),
		grpc_recovery.UnaryServerInterceptor(),
	))
}
