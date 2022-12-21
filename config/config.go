package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Port int `env:"PORT" envDefault:"7777"`

	CouchDbAddr     string `env:"COUCHDB_ADDR" envDefault:"http://127.0.0.1:5984"`
	CouchDBUser     string `env:"COUCHDB_USER" envDefault:"admin"`        // To avoid but for the sake of the simplicity
	CouchDBPassword string `env:"COUCHDB_PASSWORD" envDefault:"password"` // To avoid but for the sake of the simplicity
}

func NewConfig() (Config, error) {
	config := Config{}
	if err := env.Parse(&config); err != nil {
		return config, err
	}
	return config, nil
}
