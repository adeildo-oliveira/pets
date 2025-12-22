package config

import (
	"log"

	"github.com/spf13/viper"
)

type Server struct {
	Config   Config         `mapstructure:"server"`
	DataBase DatabaseConfig `mapstructure:"db_config"`
}

type Config struct {
	ServerPort string `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"db_host"`
	Port     string `mapstructure:"db_port"`
	User     string `mapstructure:"db_user"`
	Password string `mapstructure:"db_password"`
	Name     string `mapstructure:"db_name"`
}

func LoadConfig() *Server {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Erro ao ler o arquivo de configuração: %v", err)
	}

	var cfg Server
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Erro ao fazer unmarshal da configuração: %v", err)
	}
	return &cfg
}
