package config

import (
	"context"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Configuration struct {
	Server struct {
		Host struct {
			Port string `mapstructure:"port"`
		} `mapstructure:"host"`
	}
	AmazonConfig AmazonConfig `mapstructure:"amazon"`
}

type AmazonConfig struct {
	Region               string `mapstructure:"region"`
	Endpoint             string `mapstructure:"endpoint"`
	OrdersExecutedSqsUrl string `mapstructure:"sqsUrl"`
}

var configuration = &Configuration{}

func init() {
	environment := os.Getenv("PROFILE")
	if environment == "" {
		environment = "local"
	}
	ctx := context.Background()
	vip := viper.New()
	vip.SetConfigType("yml")
	vip.AddConfigPath("internal/config/environment")
	vip.AddConfigPath("../../../internal/config/environment")
	vip.SetConfigName(environment)

	log.Println(ctx, "Environment: "+environment)

	if err := vip.ReadInConfig(); err != nil {
		log.Fatal(ctx, err, "Fail to start config")
	}

	if err := vip.Unmarshal(configuration); err != nil {
		log.Fatal(ctx, err, "Fail to unmarshal config")
	}
}

func Config() Configuration {
	return *configuration
}
