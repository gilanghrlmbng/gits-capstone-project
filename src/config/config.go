package config

import (
	"os"

	"github.com/labstack/echo/v4"

	"github.com/joho/godotenv"
)

type Config struct {
	ServicePort string
	Database    DatabaseConfig
	Secret      string
	Port        string `env:"PORT,default=4132"`
}

type DatabaseConfig struct {
	Host     string `env:"DATABASE_HOST,default=localhost"`
	Port     string `env:"DATABASE_PORT,default=5432"`
	Username string `env:"DATABASE_USERNAME,required"`
	Password string `env:"DATABASE_PASSWORD,required"`
	Name     string `env:"DATABASE_NAME,required"`
}

func GetConfig(e *echo.Echo) Config {
	err := godotenv.Load()
	if err != nil {
		e.Logger.Error(err)
	}

	return Config{
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		Secret: os.Getenv("SECRET"),
		Port:   os.Getenv("PORT"),
	}
}

func GetConfigs(c echo.Context) Config {
	err := godotenv.Load()
	if err != nil {
		c.Logger().Error(err)
	}

	return Config{
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		Secret: os.Getenv("SECRET"),
		Port:   os.Getenv("PORT"),
	}
}
