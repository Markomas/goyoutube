package app

type Config struct {
	Port int `mapstructure:"port" default:"8080"`
}
