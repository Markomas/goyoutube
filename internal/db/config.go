package db

type Config struct {
	Dsn string `mapstructure:"dsn" default:"tmp/sqlite.db"`
}
