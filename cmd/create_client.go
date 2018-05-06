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
		name := viper.GetString("create.client.name")
		desc := viper.GetString("create.client.description")
		ownerID := snowflake.ID(viper.GetInt64("create.client.owner-id"))
		redirect := viper.GetString("create.client.redirect-uri")

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

		db, err := gorm.Open("postgres", config.DB())
		if err != nil {
			return err
		}
		return models.NewClientStore(db).Create(cl)
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
