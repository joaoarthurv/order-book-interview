package config

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type Configuration struct {
	AppName string
	Server  struct {
		Host struct {
			Port string `mapstructure:"port"`
		} `mapstructure:"host"`
	}
	PostgresConfig PostgresConfig `mapstructure:"postgres"`
	AmazonConfig   AmazonConfig   `mapstructure:"amazon"`
}

type AmazonConfig struct {
	Region               string `mapstructure:"region"`
	Endpoint             string `mapstructure:"endpoint"`
	OrdersExecutedSnsUrl string `mapstructure:"snsUrl"`
	OrdersExecutedSnsArn string `mapstructure:"snsArn"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Port     int    `mapstructure:"port"`
	Conn     Conn   `mapstructure:"conn"`
}

type Conn struct {
	Min      int    `mapstructure:"min"`
	Max      int    `mapstructure:"max"`
	Lifetime string `mapstructure:"lifetime"`
	IdleTime string `mapstructure:"idletime"`
}

func (dbc *PostgresConfig) CreateWritePool() (*pgxpool.Pool, error) {
	return dbc.CreatePool(dbc)
}

func (dbc *PostgresConfig) CreateReadPool() (*pgxpool.Pool, error) {
	dbc.ChangeHostToReadOnly()
	return dbc.CreatePool(dbc)
}

func (dbc *PostgresConfig) ChangeHostToReadOnly() {
	dbc.Host = strings.Replace(dbc.Host, "rw1", "ro1", 1)
}

func (dbc *PostgresConfig) CreatePool(config *PostgresConfig) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(fmt.Sprintf("host=%s user=%s dbname=%s "+
		"password=%s sslmode=%s pool_min_conns=%d pool_max_conns=%d "+
		"pool_max_conn_lifetime=%s pool_max_conn_idle_time=%s",
		config.Host,
		config.Username,
		config.Database,
		config.Password,
		"disable",
		config.Conn.Min,
		config.Conn.Max,
		config.Conn.Lifetime,
		config.Conn.IdleTime,
	))

	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(
		context.Background(),
		cfg,
	)

	if pool == nil {
		log.Fatal("error while select buy orders")
		return nil, errors.New("database not connected")
	}

	return pool, err
}

var configuration = &Configuration{
	AppName: "orders-management-service",
}

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

	log.Printf("Environment: " + environment)

	if err := vip.ReadInConfig(); err != nil {
		log.Fatal(ctx, err, "Fail to start config")
	}

	if err := vip.Unmarshal(configuration); err != nil {
		log.Fatal(ctx, err, "Fail to unmarshal config")
	}
}

func config() Configuration {
	return *configuration
}
