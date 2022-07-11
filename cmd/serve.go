package cmd

import (
	"fmt"
	"github.com/go-funcards/authz-service/internal/authz/db"
	"github.com/go-funcards/authz-service/internal/config"
	service "github.com/go-funcards/authz-service/internal/v1"
	"github.com/go-funcards/authz-service/proto/v1"
	"github.com/go-funcards/grpc-server"
	"github.com/go-funcards/validate"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve Authorization Service gRPC",
	Long:  "Serve Authorization Service gRPC",
	Run:   executeServeCommand,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func executeServeCommand(cmd *cobra.Command, _ []string) {
	ctx := cmd.Context()

	cfg, err := config.GetConfig(globalFlags.ConfigFile)
	if err != nil {
		log.Fatal(err)
	}

	logger, err := cfg.Log.BuildLogger(cfg.Debug)
	if err != nil {
		panic(err)
	}

	logger.Info(fmt.Sprintf("starting: %s", use))
	logger.Info(fmt.Sprintf("version: %s", version))

	validate.Default.RegisterStructRules(cfg.Rules, []any{
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

	logger.Debug("initializing mongodb")
	mongoDB, err := cfg.MongoDB.GetDatabase(ctx)
	if err != nil {
		panic(err)
	}

	logger.Debug("initializing definition storage")
	def, err := db.NewDefStorage(ctx, mongoDB, logger)
	if err != nil {
		panic(err)
	}

	logger.Debug("initializing rule storage")
	rule, err := db.NewRuleStorage(ctx, mongoDB, logger)
	if err != nil {
		panic(err)
	}

	logger.Debug("initializing subject storage")
	sub, err := db.NewSubStorage(ctx, mongoDB, logger)
	if err != nil {
		panic(err)
	}

	logger.Debug("create casbin factory")
	factory := cfg.Casbin.EnforcerFactory(def, rule, sub, nil)

	register := func(server *grpc.Server) {
		v1.RegisterAuthorizationCheckerServer(server, service.NewCheckerService(factory))
		v1.RegisterDefinitionServer(server, service.NewDefService(def))
		v1.RegisterRuleServer(server, service.NewRuleService(rule))
		v1.RegisterSubjectServer(server, service.NewSubService(sub))
	}

	grpcserver.Start(
		ctx,
		cfg.Server.Listen.Listener(logger),
		register,
		logger,
		grpc.ChainUnaryInterceptor(grpcserver.ValidatorUnaryServerInterceptor(validate.Default)),
	)
}
