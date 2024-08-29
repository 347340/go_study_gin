package config

import (
    "encoding/json"
    "os"
)

type DatabaseConfig struct {
    Host     string `json:"host"`
    User     string `json:"user"`
    Password string `json:"password"`
    Dbname   string `json:"dbname"`
    Port     int    `json:"port"`
    SSLMode  string `json:"sslmode"`
}

type Config struct {
    Database DatabaseConfig `json:"database"`
}

func LoadConfig() (Config, error) {
    var config Config
    file, err := os.Open("config.json")
    if err != nil {
        return config, err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    err = decoder.Decode(&config)
    return config, err
}
