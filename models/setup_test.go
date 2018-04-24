package models

import (
	"os"
	"path"

	"github.com/jinzhu/gorm"
	"github.com/mattes/migrate"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

// TestDB is the database used for testing.
var TestDB *gorm.DB

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	// Override the test DB URI with TEST_DB_URI!
	// -- DO NOT USE A PRODUCTION DATABASE; IT WILL BE WIPED --
	uri := os.Getenv("MEOW_TEST_DB")
	if uri == "" {
		uri = "postgres:///meow_test?sslmode=disable"
	}

	// Connect to the database!
	db, err := gorm.Open("postgres", uri)
	must(err)
	TestDB = db.LogMode(true).Debug()

	// Nuke the entire default schema.
	must(db.Exec(`DROP SCHEMA IF EXISTS public CASCADE`).Error)
	must(db.Exec(`CREATE SCHEMA public`).Error)

	// Read migrations...
	wd, err := os.Getwd()
	must(err)
	migr, err := migrate.New("file://"+path.Join(wd, "..", "migrations"), uri)
	must(err)

	// Migrate up, down, then back up to verify that both ways are working.
	must(migr.Up())
	must(migr.Down())
	must(migr.Up())

	// Finish!
	srcerr, dberr := migr.Close()
	must(srcerr)
	must(dberr)
}
