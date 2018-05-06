package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/alecthomas/chroma/quick"
	"github.com/fatih/color"
	"github.com/liclac/meow/config"
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

func dump(v interface{}) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %v", err)
		return
	}
	if !color.NoColor {
		var buf bytes.Buffer
		if err := quick.Highlight(&buf, string(data), "json", "terminal", config.HighlightStyle()); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "ERROR: %v", err)
		}
		data = buf.Bytes()
	}
	_, _ = fmt.Println(string(data))
}
