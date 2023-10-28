package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Storage struct {
	Username string `json:"username,omitempty" env:"DB_Username"`
	Password string `json:"password,omitempty" env:"DB_Password"`
	Host     string `json:"host,omitempty" env:"DB_Host"`
	Port     string `json:"port,omitempty" env:"DB_Port"`
	Name     string `json:"db___name,omitempty" env:"DB_Name"`
}

type Config struct {
	DB Storage
}

func Config_load() Config {
	var cfg Config
	err := cleanenv.ReadConfig("/home/max/GolandProjects/pets/avito-grpc/.env", &cfg)
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return cfg
}
