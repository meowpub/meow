package cmd

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/liclac/meow/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// serveCmd represents the api command
var serveCmd = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"s"},
	Short:   "Runs a web worker",
	Long:    `Runs a web worker.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		addr := viper.GetString("api.addr")

		// Create a context that is cancelled on Ctrl+C.
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		go func() {
			sigC := make(chan os.Signal, 1)
			signal.Notify(sigC, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
			<-sigC
			signal.Stop(sigC)
			cancel()
		}()

		// Build a server.
		srv := &http.Server{Handler: server.New()}

		// Listen!
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			return err
		}
		zap.L().Info("listening on...", zap.String("addr", lis.Addr().String()))

		// Run until the context exits.
		errC := make(chan error, 1)
		go func() { errC <- srv.Serve(lis) }()

		select {
		case err := <-errC:
			return err
		case <-ctx.Done():
			shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			return srv.Shutdown(shutdownCtx)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringP("addr", "a", "localhost:8000", "address to listen on")
	bindPFlags("api", serveCmd.Flags())
}
