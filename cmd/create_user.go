package cmd

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/meowpub/meow/config"
	"github.com/meowpub/meow/models"
)

var createUserCmd = &cobra.Command{
	Use:   "user id email",
	Short: "Create a user",
	Long:  "Create a user.",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		L := zap.L().Named("create.user")

		db, err := openDB(L)
		if err != nil {
			return err
		}
		stores := models.NewStores(db, nil, config.RedisKeyspace())

		id := args[0]
		email := args[1]
		noPasswd := viper.GetBool("create.user.no-passwd")

		// Look up the user's profile; it should've been created with `meow gen entity | meow ingest' already.
		profile, err := stores.Entities().GetByID(id)
		if err != nil {
			if models.IsNotFound(err) {
				return errors.Wrap(err, "please create the user's profile with 'meow gen entity as:Person' and 'meow ingest' first")
			}
			return errors.Wrap(err, "couldn't look up profile")
		}

		// Read the password from stdin.
		var password string
		if !noPasswd {
			fmt.Fprint(os.Stderr, "Password: ")
			pass, err := terminal.ReadPassword(int(os.Stdin.Fd()))
			if err != nil {
				return errors.Wrap(err, "couldn't read password from stdin")
			}
			password = string(pass)
		}

		// Create a user!
		user, err := models.NewUser(profile.ID, email, password)
		if err != nil {
			return err
		}
		return stores.Users().Save(user)
	},
}

func init() {
	createCmd.AddCommand(createUserCmd)

	createUserCmd.Flags().BoolP("no-passwd", "N", false, "don't set a password; disallows login")
	bindPFlags("create.user", createUserCmd.Flags())
}
