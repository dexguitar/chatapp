package cmd

import (
	"fmt"

	"github.com/dexguitar/chatapp/configs"
	"github.com/dexguitar/chatapp/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/spf13/cobra"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattes/migrate/source/file"
)

var migrateCmd = func(c *configs.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "migrate",
		Long:  `migrations`,
		RunE:  runMigrate(c),
		Args:  cobra.ExactArgs(1),
	}
}

func runMigrate(c *configs.Config) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		dbConn, err := db.NewDatabase(c.Postgres)
		if err != nil {
			return fmt.Errorf("could not initialize database connection: %s", err)
		}
		err = dbConn.GetDB().Ping()
		if err != nil {
			return fmt.Errorf("failed to ping DB: %s", err.Error())
		}
		defer dbConn.Close()

		driver, err := postgres.WithInstance(dbConn.GetDB(), &postgres.Config{})
		if err != nil {
			return fmt.Errorf("failed to create DB driver: %s", err.Error())
		}

		m, err := migrate.NewWithDatabaseInstance(c.MigrationURL, "postgres", driver)

		if args[0] == "up" {
			err = m.Up()
			if err != nil {
				return fmt.Errorf("failed to run migration: %s", err.Error())
			}
		} else if args[0] == "down" {
			err = m.Down()
			if err != nil {
				return fmt.Errorf("failed to run migration: %s", err.Error())
			}
		} else {
			return fmt.Errorf("invalid migrate arg, should be `up` or `down`")
		}

		return nil
	}
}
