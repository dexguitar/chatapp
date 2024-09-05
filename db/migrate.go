package db

import (
	"fmt"

	"github.com/dexguitar/chatapp/configs"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

func Migrate(c *configs.Config) error {
	op := "db.Migrate"

	dbConn, err := NewDatabase(c)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = dbConn.Ping()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer dbConn.Close()

	driver, err := postgres.WithInstance(dbConn, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	m, err := migrate.NewWithDatabaseInstance(c.MigrationURL, "chatapp", driver)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
