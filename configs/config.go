package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Host          string   `mapstructure:"HOST"`
	Port          string   `mapstructure:"PORT"`
	DBPrimary     string   `mapstructure:"DB_PRIMARY"`
	DBRepl1       string   `mapstructure:"DB_REPL1"`
	DBRepl2       string   `mapstructure:"DB_REPL2"`
	MigrationURL  string   `mapstructure:"MIGRATION_URL"`
	KafkaBrokers  []string `mapstructure:"KAFKA_BROKERS"`
	KafkaTopic    string   `mapstructure:"KAFKA_TOPIC"`
	ConsumerGroup string   `mapstructure:"CONSUMER_GROUP"`
}

func LoadConfig() (config *Config, err error) {
	v := viper.New()

	v.SetConfigFile(".env")

	v.AutomaticEnv()

	v.SetDefault("HOST", "0.0.0.0")
	v.SetDefault("PORT", ":8080")
	v.SetDefault("DB_PRIMARY", "postgresql://postgres:qwerty@localhost:5432/chatapp?sslmode=disable")
	v.SetDefault("DB_REPL1", "postgresql://postgres:qwerty@localhost:5433/chatapp?sslmode=disable")
	v.SetDefault("DB_REPL2", "postgresql://postgres:qwerty@localhost:5434/chatapp?sslmode=disable")
	v.SetDefault("MIGRATION_URL", "file:///db/migrations")
	v.SetDefault("KAFKA_BROKERS", "localhost:9093,localhost:9094,localhost:9095")
	v.SetDefault("KAFKA_TOPIC", "chatapp")
	v.SetDefault("CONSUMER_GROUP", "test-group")

	err = v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = v.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
