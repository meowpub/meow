package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/monzo/typhon"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/meowpub/meow/api"
	"github.com/meowpub/meow/config"
	"github.com/meowpub/meow/config/secrets"
	"github.com/meowpub/meow/lib"
)

// serveCmd represents the api command
var serveCmd = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"s"},
	Short:   "Runs a web worker",
	Long:    `Runs a web worker.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		L := zap.L().Named("serve")

		addr := viper.GetString("api.addr")
		domains := viper.GetStringSlice("api.domain")

		// We need to know what domains are local, because it changes some semantics.
		if len(domains) == 0 {
			return errors.Errorf("no domains configured: specify -d/--domain, eg. `meow s -d localhost`")
		}

		// Derive subkeys from the master secret; this has to be done before use.
		if err := secrets.Init(L.Named("secrets"), config.Secret()); err != nil {
			return errors.Wrap(err, "secrets.Init")
		}

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

		// Connect to the database + redis, disconnect when the context exits.
		db, err := openDB(L.Named("db"))
		if err != nil {
			return err
		}
		defer func() { lib.Report(ctx, db.Close(), "couldn't cleanly close database connection") }()

		r, err := openRedis(L.Named("redis"))
		if err != nil {
			return err
		}
		defer func() { lib.Report(ctx, r.Close(), "couldn't cleanly close redis connection") }()

		// Listen!
		svc := api.Service.
			Filter(api.StoresFilter(db, r)).
			Filter(typhon.ExpirationFilter).
			Filter(api.PanicFilter).
			Filter(api.ErrorFilter).
			Filter(typhon.H2cFilter)
		srv, err := typhon.Listen(svc, addr)
		if err != nil {
			return err
		}
		L.Info("Listening!", zap.String("addr", addr), zap.Strings("domains", domains))

		// Stop that when the context exits, eg. on Ctrl+C.
		<-ctx.Done()
		cancelCtx, cancelCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelCancel()
		srv.Stop(cancelCtx)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringSliceP("domain", "d", nil, "domain(s) to serve on (eg. \"example.com\")")
	serveCmd.Flags().StringP("addr", "a", "localhost:8000", "address to listen on")
	bindPFlags("api", serveCmd.Flags())
}
