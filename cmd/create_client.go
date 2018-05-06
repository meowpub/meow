package cmd

import (
	"github.com/bwmarrin/snowflake"
	"github.com/jinzhu/gorm"
	"github.com/liclac/meow/config"
	"github.com/liclac/meow/models"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// createClientCmd represents the createClient command
var createClientCmd = &cobra.Command{
	Use:   "client",
	Short: "Creates an OAuth client application",
	Long:  `Creates an OAuth client application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Collect flags upfront.
		name := viper.GetString("create.client.name")
		desc := viper.GetString("create.client.description")
		ownerID := snowflake.ID(viper.GetInt64("create.client.owner-id"))
		redirect := viper.GetString("create.client.redirect-uri")

		// Create a client object.
		clData := models.ClientUserData{
			Name:        name,
			Description: desc,
		}
		if ownerID != 0 {
			clData.OwnerID = &ownerID
		}
		cl, err := models.NewClient(clData, redirect)
		if err != nil {
			return err
		}
		dump(cl)

		// Connect to the database.
		db, err := gorm.Open("postgres", config.DB())
		if err != nil {
			return err
		}

		// Insert the objects inside a transaction.
		tx := db.Begin()
		if err := models.NewClientStore(db).Create(cl); err != nil {
			return err
		}
		return tx.Commit().Error
	},
}

func init() {
	createCmd.AddCommand(createClientCmd)

	createClientCmd.Flags().StringP("name", "n", "", "human-readable name (required)")
	createClientCmd.Flags().StringP("description", "d", "", "human-readable description")
	createClientCmd.Flags().Int64P("owner-id", "O", 0, "ID of the owning user")
	createClientCmd.Flags().StringP("redirect-uri", "r", "", "redirect URI (required)")
	bindPFlags("create.client", createClientCmd.Flags())
}
