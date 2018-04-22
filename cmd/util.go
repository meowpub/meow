package cmd

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/multierr"
)

func bindPFlags(pfx string, flags *pflag.FlagSet) error {
	var errs []error
	flags.VisitAll(func(flag *pflag.Flag) {
		key := flag.Name
		if pfx != "" {
			key = pfx + "." + key
		}
		if err := viper.BindPFlag(key, flag); err != nil {
			errs = append(errs, err)
		}
	})
	return multierr.Combine(errs...)
}
