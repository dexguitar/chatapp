package db

import (
	"database/sql"
	"fmt"

	"github.com/dexguitar/chatapp/configs"
	psql "github.com/snaffi/pg-helper"
)

func NewDatabase(c *configs.Config) (*sql.DB, error) {
	op := "db.NewDatabase"

	db, err := sql.Open("postgres", "postgresql://postgres:qwerty@localhost:5433/chatapp?sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return db, nil
}

func NewPostgresDB(appConfig *configs.Config) (psql.DB, error) {
	op := "db.NewPostgresDB"

	poolConf := newPoolConfig(appConfig)
	replicaSet := psql.WithRoundRobinReplicaSet(poolConf["repl1"], poolConf["repl2"])
	DB, err := psql.NewConnectionPool(poolConf["primary"], replicaSet)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return DB, nil
}

func newPoolConfig(appConfig *configs.Config) map[string]psql.Config {
	m := make(map[string]psql.Config)
	m["primary"] = psql.Config{
		MaxConnections:    100,
		HealthCheckPeriod: 100,
		DSN:               appConfig.DBPrimary,
	}
	m["repl1"] = psql.Config{
		MaxConnections:    100,
		HealthCheckPeriod: 100,
		DSN:               appConfig.DBRepl1,
	}
	m["repl2"] = psql.Config{
		MaxConnections:    100,
		HealthCheckPeriod: 100,
		DSN:               appConfig.DBRepl1,
	}

	return m
}
