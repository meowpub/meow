package cmd

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns"
	"github.com/meowpub/meow/lib"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// genEntityCmd represents the generate entity command
var genEntityCmd = &cobra.Command{
	Use:   "entity id ns:Type [ns:Type...]",
	Short: "Generate JSON-LD entities",
	Long:  `Generate JSON-LD entities from the commandline.`,
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]
		shortTypes := args[1:]

		// URLs must be normalised before going in the database.
		u, err := url.Parse(id)
		if err != nil {
			return err
		}
		nu := lib.NormalizeURL(*u)
		id = nu.String()

		// Types are passed as "ns:Type", eg. "as:Note".
		longTypes := make([]string, len(shortTypes))
		for i, short := range shortTypes {
			long, err := resolveShortType(short)
			if err != nil {
				return err
			}
			longTypes[i] = long
		}

		obj := &ld.Object{V: map[string]interface{}{
			"@id":   id,
			"@type": longTypes,
		}}

		// All attributes we're aware of are exposed as flags; as:content -> --as-content=...
		// These default to being strings, but can be explicitly typed with "type=value".
		for _, n := range ns.Namespaces {
			for _, p := range n.Props {
				raw := viper.GetStringSlice("create.entity." + n.Short + "-" + p.Short)
				if len(raw) == 0 {
					continue
				}
				vs := make([]interface{}, len(raw))
				for i, s := range raw {
					v, err := parseTypedValue(s)
					if err != nil {
						return errors.Wrap(err, n.ID)
					}
					vs[i] = v
				}
				obj.V[p.ID] = vs
			}
		}

		dump(obj)

		return nil
	},
}

func parseTypedValue(s string) (interface{}, error) {
	parts := strings.SplitN(s, "=", 2)
	typ := ""
	val := parts[0]
	if len(parts) > 1 {
		typ = parts[0]
		val = parts[1]
	}

	switch typ {
	case "", "s", "str":
		return ld.Value(val), nil
	case "i", "int":
		i, err := strconv.ParseInt(val, 10, 64)
		return ld.Value(i), err
	case "id":
		return ld.ID(val), nil
	default:
		return nil, errors.Errorf("unknown type: %s", typ)
	}
}

func resolveShortType(s string) (string, error) {
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		return "", errors.Errorf("type must be in ns.Type form (eg. as.Note): %s", s)
	}
	nName := parts[0]
	tName := parts[1]
	for _, n := range ns.Namespaces {
		if n.Short != nName {
			continue
		}
		for _, t := range n.Types {
			if t.Short == tName {
				return t.ID, nil
			}
		}
		return "", errors.Errorf("unknown type: %s%s", n.ID, tName)
	}
	return "", errors.Errorf("unknown namespace: %s", nName)
}

func init() {
	genCmd.AddCommand(genEntityCmd)

	// Every dang attribute we know about is a commandline flag.
	// This is dumb, but I can't think of anything better that doesn't involve throwing
	// multityping under a bus or typing full attribute names, which is just cursed.
	for _, n := range ns.Namespaces {
		for _, p := range n.Props {
			genEntityCmd.Flags().StringSlice(n.Short+"-"+p.Short, nil, p.Comment)
		}
	}
	bindPFlags("create.entity", genEntityCmd.Flags())
}
