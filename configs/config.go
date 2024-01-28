package configs

import "github.com/spf13/viper"

// Config app envs
type Config struct {
	Host         string `mapstructure:"HOST"`
	Port         string `mapstructure:"PORT"`
	Postgres     string `mapstructure:"POSTGRES"`
	MigrationURL string `mapstructure:"MIGRATION_URL"`
}

func LoadConfig(path string) (config *Config, err error) {
	v := viper.New()

	v.SetConfigFile(path)

	v.AutomaticEnv()

	v.SetDefault("HOST", defaultHost)
	v.SetDefault("PORT", defaultPort)
	v.SetDefault("POSTGRES", defaultPgString)
	v.SetDefault("MIGRATION_URL", defaultMigrationUrl)

	err = v.ReadInConfig()
	if err != nil {
		return
	}

	err = v.Unmarshal(&config)
	return
}
