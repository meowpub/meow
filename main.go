package main

import (
	"github.com/liclac/meow/cmd"

	// Autoload a .env file as environment variables.
	_ "github.com/joho/godotenv/autoload"

	// Load database drivers.
	_ "github.com/mattes/migrate/database/postgres"

	// Why do I have to load support for local files for migrations.
	_ "github.com/mattes/migrate/source/file"
)

func main() {
	cmd.Execute()
}
