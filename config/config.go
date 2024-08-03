package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App App
		Db  Db
	}

	App struct {
		Port int
	}

	Db struct {
		Postgres Postgres
		Mongo    Mongo
	}

	Postgres struct {
		Host     string
		Port     int
		Username string
		Password string
		DBName   string
		SSLMode  string
		TimeZone string
	}

	Mongo struct {
		Host           string
		Port           int
		Username       string
		Password       string
		CollectionName string
		DatabaseName   string
	}
)

func GetConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %v ", err))
	}

	return Config{
		App: App{
			Port: viper.GetInt("app.server.port"),
		},
		Db: Db{
			Postgres: Postgres{
				Host:     viper.GetString("postgres.host"),
				Port:     viper.GetInt("postgres.port"),
				Username: viper.GetString("postgres.user"),
				Password: viper.GetString("postgres.password"),
				DBName:   viper.GetString("postgres.dbname"),
				SSLMode:  viper.GetString("postgres.sslmode"),
				TimeZone: viper.GetString("postgres.timezone"),
			},
			Mongo: Mongo{
				Host:           viper.GetString("mongo.host"),
				Port:           viper.GetInt("mongo.port"),
				Username:       viper.GetString("mongo.user"),
				Password:       viper.GetString("mongo.password"),
				CollectionName: viper.GetString("mongo.collection_name"),
				DatabaseName:   viper.GetString("mongo.database_name"),
			},
		},
	}
}
