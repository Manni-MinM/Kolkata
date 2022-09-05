package config

import (
    "github.com/spf13/viper"
)

type Config struct {
    Redis     RedisConfig     `mapstructure:"redis"`
    Server    ServerConfig    `mapstructure:"server"`
}

type ServerConfig struct {
    Port    int    `mapstructure:"port"`
}

type RedisConfig struct {
    Port        int       `mapstructure:"port"`
    Addr        string    `mapstructure:"address"`
    Password    string    `mapstructure:"password"`
    DB          string    `mapstructure:"db"`
}

func LoadConfig() (*Config, error) {
    viper.AddConfigPath("./configs")
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")

    err := viper.ReadInConfig()
    if err != nil {
        return &Config{}, err
    }

    var conf Config
    err = viper.Unmarshal(&conf)
    if err != nil {
        return &Config{}, err
    }

    return &conf, nil
}

