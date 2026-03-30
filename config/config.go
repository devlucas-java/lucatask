package config

import (
	"fmt"

	"github.com/devlucas-java/lucatask/internal/domain"
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type conf struct {
	Port           string `mapstructure:"PORT"`
	DB_Host        string `mapstructure:"DB_HOST"`
	DB_Port        string `mapstructure:"DB_PORT"`
	DB_User        string `mapstructure:"DB_USER"`
	DB_Password    string `mapstructure:"DB_PASSWORD"`
	DB_Name        string `mapstructure:"DB_NAME"`
	DB_Driver      string `mapstructure:"DB_DRIVER"`
	JWT_Secret     string `mapstructure:"JWT_SECRET"`
	JWT_Expires_In int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth      *jwtauth.JWTAuth
}

var cfg *conf

func GetConfig() *conf {
	return cfg
}

func InitConfig() *conf {

	cfg = &conf{}

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.SetConfigName("default-config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(cfg)
	if err != nil {
		panic(err)
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWT_Secret), nil)

	return cfg
}

func InitDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB_User,
		cfg.DB_Password,
		cfg.DB_Host,
		cfg.DB_Port,
		cfg.DB_Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&domain.Task{})

	return db
}
