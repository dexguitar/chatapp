package cmd

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dexguitar/chatapp/db"

	"github.com/dexguitar/chatapp/configs"
	"github.com/spf13/cobra"
)

type application struct {
	config *configs.Config
	router http.Handler
}

func newApplication(config *configs.Config, router http.Handler) *application {
	return &application{
		config: config,
		router: router,
	}
}

var runServerCmd = func() *cobra.Command {
	return &cobra.Command{
		Use:   "runserver",
		Short: "Runs a server",
		Long:  `Runs a server on specified host and port (first and second argument)`,
		RunE:  runServer(),
	}
}

func runServer() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		op := "cmd.runServer"

		app, err := InitApplication()
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		err = db.Migrate(app.config)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		srv := &http.Server{
			Addr:         app.config.Port,
			Handler:      app.router,
			ReadTimeout:  time.Second * 10,
			WriteTimeout: time.Second * 10,
		}

		go func() {
			slog.Info("✅ CHATAPP STARTED ✅", "address", app.config.Host+app.config.Port)
			if err = srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				slog.Error("FATAL: server listen", "error", err)
				os.Exit(1)
			}
		}()

		ctx := context.Background()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, os.Interrupt, syscall.SIGINT, syscall.SIGQUIT)
		<-quit

		slog.Warn("ChatApp shutting down")

		if err = srv.Shutdown(ctx); err != nil {
			return err
		}

		return nil
	}
}
