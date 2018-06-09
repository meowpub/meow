package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/meowpub/meow/config"
	"github.com/meowpub/meow/config/secrets"
)

var rootCmd = &cobra.Command{
	Use:   "meow",
	Short: "A cuter ActivityPub server",
	Long:  `A cuter ActivityPub server.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Set up the global logger.
		lconf := zap.NewDevelopmentConfig()
		if config.IsProd() {
			lconf = zap.NewProductionConfig()
		}
		l, err := lconf.Build()
		if err != nil {
			return err
		}
		zap.RedirectStdLog(l)
		zap.ReplaceGlobals(l)

		// Disable colours globally if asked.
		if config.NoColour() {
			color.NoColor = true
		}

		// Derive subkeys from the master secret; this has to be done before use.
		if err := secrets.Init(l.Named("secrets"), config.Secret()); err != nil {
			return errors.Wrap(err, "secrets.Init")
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolP("prod", "P", false, "run in production mode")
	rootCmd.PersistentFlags().StringP("secret", "S", "", "secret key (base64, min 64 bytes)")
	rootCmd.PersistentFlags().Int64("node-id", 0, "node id used for snowflake generation")
	rootCmd.PersistentFlags().String("db", "postgres:///meow?sslmode=disable", "database connection uri")
	rootCmd.PersistentFlags().String("redis", "redis://localhost:6379/0", "redis connection uri")
	rootCmd.PersistentFlags().String("redis-keyspace", "meow", "prepended (+ a ':') to all redis keys")

	rootCmd.PersistentFlags().String("highlight-style", "vim", "style used for syntax highlighting")
	rootCmd.PersistentFlags().Bool("no-colour", false, "disable all colours in the output")

	viper.BindPFlags(rootCmd.PersistentFlags())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetEnvPrefix("meow")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.AutomaticEnv()
}
