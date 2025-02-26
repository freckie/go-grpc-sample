package run

import (
	"context"
	"net"

	_ "database/sql"

	"frec.kr/tdoo/cmd/tdoo/config"
	l "frec.kr/tdoo/log"
	"frec.kr/tdoo/pkg/v1/gen/tdoo"
	"frec.kr/tdoo/pkg/v1/gen/tdoo/orm"
	"frec.kr/tdoo/pkg/v1/handler"
	"frec.kr/tdoo/pkg/v1/repository"
	"frec.kr/tdoo/pkg/v1/service"
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

		l.Setup(cfg.Logger)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		db, err := orm.Open(cfg.DB.Driver, cfg.DB.Source)
		if err != nil {
			l.Log.Errorf("failed to connect db: %v", err)
			return err
		}
		if err := db.Schema.Create(ctx); err != nil {
			l.Log.Errorf("failed to create schema: %v", err)
			return err
		}

		repo := repository.NewRepository(db)
		svc := service.NewService(repo)
		handler := handler.NewHandler(svc)

		srv := grpc.NewServer()
		tdoo.RegisterTaskServiceServer(srv, handler.Task())
		tdoo.RegisterUserServiceServer(srv, handler.User())

		listener, err := net.Listen("tcp", cfg.Addr)
		if err != nil {
			l.Log.Errorf("failed to listen: %v", err)
			return err
		}

		l.Log.Info("gRPC server is listening on [%s]", listener.Addr().String())
		if err := srv.Serve(listener); err != nil {
			l.Log.Errorf("failed to serve: %v", err)
			return err
		}

		return nil
	},
}

func init() {
	RunCmd.Flags().StringVar(&cfgFile, "config", "config.yaml", "Path to configuration file")
}
