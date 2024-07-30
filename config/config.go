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
		Host           string
		Port           int
		Username       string
		Password       string
		CollectionName string
		DatabaseName   string
	}
)

func GetConfig() Config {

	// executable, err := os.Executable()
	// if err != nil {
	// 	panic(err)
	// }
	// executableDir := filepath.Dir(executable)
	// configPath := filepath.Join(executableDir, "../config.yml")

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
			Host:           viper.GetString("database.host"),
			Port:           viper.GetInt("database.port"),
			Username:       viper.GetString("database.user"),
			Password:       viper.GetString("database.password"),
			CollectionName: viper.GetString("database.collection_name"),
			DatabaseName:   viper.GetString("database.database_name"),
		},
	}
}
