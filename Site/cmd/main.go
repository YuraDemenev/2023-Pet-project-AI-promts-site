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

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	err := initConfig()
	if err != nil {
		logrus.Fatalf("initializing config error: %s", err.Error())
	}
	//Get url
	url := viper.GetString("url")

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

	//Connect to Redis
	//Read data from cofig.yml and send as RedisConfig to NewPostgresDB
	//and send n minutes as "data live"
	redisCache := cache.NewRedisCache(cache.RedisConfig{
		Host: viper.GetString("redis.host"),
		DB:   viper.GetInt("redis.db"),
	})

	redis := cache.NewCache(redisCache, time.Minute*30)
	//Initialize repository (service through which
	//will run functions interacting with the database)
	repos := repository.NewRepository(db, redis)

	//Initialize service (it will be used to launch the
	//Intermediate functions and interaction with the database will be launched through it
	//through repository)
	services := service.NewService(repos)

	//Initialize hadlers (paths on which requests can be sent)
	handlers := handler.NewHandler(services, redis, url)
	//Initialize the server
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
