package config

import "github.com/spf13/viper"

type (
	Env struct {
		App      *App
		Database *Database
	}

	App struct {
		Port int
	}

	Database struct {
		Host string
		Port int
		User string
		Pass string
		Name string
	}
)

func NewEnv() (*Env, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	env := &Env{
		&App{
			Port: viper.GetInt("app.port"),
		},
		&Database{
			Host: viper.GetString("database.host"),
			Port: viper.GetInt("database.port"),
			User: viper.GetString("database.user"),
			Pass: viper.GetString("database.pass"),
			Name: viper.GetString("database.name"),
		},
	}

	return env, nil
}
