package cmd

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/liclac/meow/config"
	"github.com/liclac/meow/jsonld"
	"github.com/liclac/meow/models"
	"github.com/liclac/meow/models/entities"
)

// createUserCmd represents the createUser command
var createUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Creates a user",
	Long:  `Creates a user`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id := viper.GetString("create.user.id")
		preferredUsername := viper.GetString("create.user.username")
		email := viper.GetString("create.user.email")
		password := viper.GetString("create.user.password")

		db, err := gorm.Open("postgres", config.DB())
		if err != nil {
			return err
		}
		tx := db.Begin()

		rawStore := models.NewEntityStore(tx)
		store := entities.NewStore(rawStore)
		userStore := models.NewUserStore(tx)

		person, err := entities.NewPerson(store, id)
		if err != nil {
			return err
		}

		if preferredUsername != "" {
			person.PreferredUsername = jsonld.ToString(preferredUsername)
			person.Name = jsonld.ToLangString(preferredUsername)
		} else {
			person.Name = jsonld.ToLangString("A new user")
		}
		person.Summary = jsonld.ToLangString("A newly created Meow🐱 user")

		if err := store.Save(person); err != nil {
			return err
		}

		usr, err := models.NewUser(person.GetSnowflake(), email, password)
		if err != nil {
			return err
		}

		dump(person)
		dump(usr)

		if err := userStore.Save(usr); err != nil {
			return err
		}

		return tx.Commit().Error
	},
}

func init() {
	createCmd.AddCommand(createUserCmd)

	createUserCmd.Flags().StringP("id", "i", "", "user's uri id (required)")
	createUserCmd.Flags().StringP("email", "e", "", "user's email (required)")
	createUserCmd.Flags().StringP("password", "p", "", "user's password (required)")
	createUserCmd.Flags().StringP("username", "u", "", "user's username (optional)")

	bindPFlags("create.user", createUserCmd.Flags())
}
