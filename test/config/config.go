package config

import (
	"errors"
	"os"
)

type Config struct {
	Mongo struct {
		Username string `json:"username"`
		Password string `json:"password"`
		DBName   string `json:"db_name"`
		Host     string `json:"host"`
		Port     string `json:"port"`
	}
}

func ENV() (Config, error) {
	username, ok := os.LookupEnv("MONGODB_USERNAME")
	if !ok {
		return Config{}, errors.New("username env not found")
	}

	return Config{Mongo: struct {
		Username string `json:"username"`
		Password string `json:"password"`
		DBName   string `json:"db_name"`
		Host     string `json:"host"`
		Port     string `json:"port"`
	}(struct {
		Username string
		Password string
		DBName   string
		Host     string
		Port     string
	}{Username: username, Password: "", DBName: "", Host: "", Port: ""})}, nil
}
