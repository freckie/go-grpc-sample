package run

import (
	"context"
	"net"

	_ "database/sql"

	"frec.kr/tdoo/cmd/tdoo/config"
	"frec.kr/tdoo/pkg/v1/gen/tdoo"
	"frec.kr/tdoo/pkg/v1/gen/tdoo/orm"
	"frec.kr/tdoo/pkg/v1/handler"
	"frec.kr/tdoo/pkg/v1/middleware"
	"frec.kr/tdoo/pkg/v1/repository"
	"frec.kr/tdoo/pkg/v1/service"
	"frec.kr/tdoo/trace"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var cfgFile string

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Start the server",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig(cfgFile)
		if err != nil {
			return err
		}
		if err = cfg.Evaluate(); err != nil {
			return err
		}

		l := trace.NewLogger(cfg.Logger)
		ctx := trace.WithLogger(context.Background(), l)

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		// db setup
		db, err := orm.Open(cfg.DB.Driver, cfg.DB.Source)
		if err != nil {
			l.Errorf("failed to connect db: %v", err)
			return err
		}
		if err := db.Schema.Create(ctx); err != nil {
			l.Errorf("failed to create schema: %v", err)
			return err
		}

		// dependency injection
		repo := repository.NewRepository(db)
		svc := service.NewService(repo)
		handler := handler.NewHandler(svc)

		// grpc server initialization
		opts := []grpc.ServerOption{
			grpc.ChainUnaryInterceptor(middleware.UnaryLogger(l)),
			grpc.ChainStreamInterceptor(middleware.StreamLogger(l)),
		}
		srv := grpc.NewServer(opts...)
		tdoo.RegisterTaskServiceServer(srv, handler.Task())
		tdoo.RegisterUserServiceServer(srv, handler.User())

		// serve
		listener, err := net.Listen("tcp", cfg.Addr)
		if err != nil {
			l.Errorf("failed to listen: %v", err)
			return err
		}
		l.Infof("gRPC server is listening on [%s]", listener.Addr().String())
		if err := srv.Serve(listener); err != nil {
			l.Errorf("failed to serve: %v", err)
			return err
		}

		return nil
	},
}

func init() {
	RunCmd.Flags().StringVar(&cfgFile, "config", "config.yaml", "path to configuration file")
}
