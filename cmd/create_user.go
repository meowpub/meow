package cmd

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/meowpub/meow/config"
	"github.com/meowpub/meow/jsonld"
	"github.com/meowpub/meow/models"
	"github.com/meowpub/meow/models/entities"
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

		stores := models.NewStores(tx, nil, "")
		store := entities.NewStore(stores.Entities())
		userStore := stores.Users()

		ctx := context.Background()
		ctx = models.WithStores(ctx, stores)
		ctx = entities.WithStore(ctx, store)

		person_, err := store.NewEntity(ctx, "person", id)
		if err != nil {
			return err
		}
		person := person_.(*entities.Person)

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
