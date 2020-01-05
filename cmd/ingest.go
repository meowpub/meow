package cmd

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/models"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var ingestCmd = &cobra.Command{
	Use:   "ingest",
	Short: "Ingest JSON from stdin",
	Long:  "Ingest JSON from stdin",
	RunE: func(cmd *cobra.Command, args []string) error {
		L := zap.L().Named("ingest")

		// Read object from stdin.
		indata, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		var kv map[string]interface{}
		if err := json.Unmarshal(indata, &kv); err != nil {
			return err
		}
		kv, err = ld.Expand(context.Background(), kv, "")
		if err != nil {
			return err
		}
		obj := ld.Object{V: kv}
		dump(obj)

		// Entities ingested this way are always ObjectEntities.
		// (Add a flag to say otherwise if that ever becomes useful.)
		e := &models.Entity{
			ID:   lib.GenSnowflake(),
			Kind: models.ObjectEntity,
			Obj:  &obj,
		}

		// Connect to the database.
		db, err := openDB(L)
		if err != nil {
			return err
		}

		// Insert the objects inside a transaction.
		tx := db.Begin()
		if err := models.NewEntityStore(db).Save(e); err != nil {
			return err
		}
		return tx.Commit().Error
	},
}

func init() {
	rootCmd.AddCommand(ingestCmd)
}
