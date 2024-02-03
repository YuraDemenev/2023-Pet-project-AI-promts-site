package main

import (
	"site/pkg/cache"
	site "site/pkg/elements"
	"site/pkg/handler"
	"site/pkg/repository"
	"site/pkg/service"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const IpAdress = "imagepromts.ru"

// var (
// 	postCache cache.CacheImages = cache.NewRedisCache(IpAdress, 0, time.Hour)
// )

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	err := initConfig()
	if err != nil {
		logrus.Fatalf("initializing config error: %s", err.Error())
	}

	// if err := godotenv.Load(); err != nil {
	// 	logrus.Fatalf("error loading env variables: %s", err.Error())
	// }

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		UserName: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		// Password: os.Getenv("DB_PASSWORD"),
		Password: viper.GetString("db.password"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	postCache := cache.NewRedisCache("94.228.121.176:6379", 0, time.Hour)

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, postCache)
	server := new(site.Server)

	err = server.Run(viper.GetString("port"), handlers.InitRoutes())
	if err != nil {
		logrus.Fatalf("error while runnig server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("../config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
