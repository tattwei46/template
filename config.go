package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Service struct {
	Host string
	Port int64
}

type MongoConnection struct {
	URL string
}

var serviceInfo map[string]Service
var mongoInfo map[string]MongoConnection

var config = viper.New()

var serviceKey = "service"
var mongoKey = "mongodb"

func setupConfig() {
	// Read config.toml
	config.SetConfigName("config")
	pwd, _ := os.Getwd()
	// Read config.toml in cur dir
	config.AddConfigPath(pwd)
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	getServiceConnection(serviceKey)
	getMongoConnection(mongoKey)

}

func getServiceConnection(key string) {
	serviceInfo = make(map[string]Service)
	// Get service info
	service := config.GetStringMap(key)
	if len(service) == 0 {
		log.Fatalf("Service info not found in config.toml")
	}

	host := service["host"].(string)
	port := service["port"].(int64)

	serviceInfo[serviceKey] = Service{
		Host: host,
		Port: port,
	}
}

func getMongoConnection(key string) {
	k := fmt.Sprintf("%s.url", key)
	con := config.Get(k)

	if con == nil {
		log.Fatalf("Mongodb not found in config.toml")
	}

	conString := con.(string)
	mongoInfo[key] = MongoConnection{
		URL: conString,
	}
}
