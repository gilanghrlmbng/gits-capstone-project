package config

import (
	"os"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/joho/godotenv"
)

type Config struct {
	ServicePort    string
	Database       DatabaseConfig
	DatabaseInit   DatabaseInit
	Secret         string
	Port           string `env:"PORT,default=80"`
	Email          string `env:"EMAIL,required"`
	PasswordEmail  string `env:"PASSWORD_EMAIL,required"`
	FirebaseApiKey string `env:"FIREBASE_API_KEY,required"`
}

type DatabaseConfig struct {
	Host     string `env:"DATABASE_HOST,default=localhost"`
	Port     string `env:"DATABASE_PORT,default=5432"`
	Username string `env:"DATABASE_USERNAME,required"`
	Password string `env:"DATABASE_PASSWORD,required"`
	Name     string `env:"DATABASE_NAME,required"`
}

type DatabaseInit struct {
	ResetTable bool `env:"RESET_TABLES, default=false"`
	SeedTable  bool `env:"SEED_TABLES, default=false"`
}

func GetConfig(e *echo.Echo) Config {
	err := godotenv.Load()
	if err != nil {
		e.Logger.Error(err)
	}

	seedTable, err := strconv.ParseBool(os.Getenv("SEED_TABLES"))
	if err != nil {
		e.Logger.Error(err)
	}
	resetTable, err := strconv.ParseBool(os.Getenv("RESET_TABLES"))
	if err != nil {
		e.Logger.Error(err)
	}

	databaseInit := DatabaseInit{
		ResetTable: resetTable,
		SeedTable:  seedTable,
	}

	return Config{
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		DatabaseInit:   databaseInit,
		Secret:         os.Getenv("SECRET"),
		Port:           os.Getenv("PORT"),
		Email:          os.Getenv("EMAIL"),
		PasswordEmail:  os.Getenv("PASSWORD_EMAIL"),
		FirebaseApiKey: os.Getenv("FIREBASE_API_KEY"),
	}
}

func GetConfigs(c echo.Context) Config {
	err := godotenv.Load()
	if err != nil {
		c.Logger().Error(err)
	}

	seedTable, err := strconv.ParseBool(os.Getenv("SEED_TABLES"))
	if err != nil {
		c.Logger().Error(err)
	}
	resetTable, err := strconv.ParseBool(os.Getenv("RESET_TABLES"))
	if err != nil {
		c.Logger().Error(err)
	}

	databaseInit := DatabaseInit{
		ResetTable: resetTable,
		SeedTable:  seedTable,
	}

	return Config{
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		DatabaseInit:   databaseInit,
		Secret:         os.Getenv("SECRET"),
		Port:           os.Getenv("PORT"),
		Email:          os.Getenv("EMAIL"),
		PasswordEmail:  os.Getenv("PASSWORD_EMAIL"),
		FirebaseApiKey: os.Getenv("FIREBASE_API_KEY"),
	}
}
