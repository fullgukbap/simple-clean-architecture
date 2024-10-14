package config

import (
	"os"

	"github.com/pelletier/go-toml"
)

type Config struct {
	Database struct {
		User            string
		Password        string
		Host            string
		Port            string
		Database        string
		MaxIdleConns    int
		MaxOpenConns    int
		ConnMaxLifeTime string
	}
}

func NewConfig(path string) (*Config, error) {

	config := new(Config)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	err = toml.NewDecoder(file).Decode(config)
	if err != nil {
		return nil, err
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	return config, nil
}
