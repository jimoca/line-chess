package conf

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

var ENV = os.Getenv("ENV")
var loadConfigOnce sync.Once

type Config struct {
	ENV string

	// server
	HOST string
	PORT int
	URL  string

	// jwt
	JWT_ISSUER string
	JWT_SECRET string

	// db
	DB_HOST_PORT string
	// DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_DATABASE string
	// DB_CONN_STR string

	// service
	ISP_ML_SERVICE_HOST string
	ISP_ML_SERVICE_PORT int
	ISP_ML_SERVICE_URL  string
}

var config Config

func Conf() Config {
	loadConfigOnce.Do(loadConfig)
	return config
}

func dev(config *Config) {
	log.Println("Using dev environment, loading config...")
	config.DB_HOST_PORT = "localhost:6379"
	// config.DB_PORT = 6379
	// config.DB_USER = "dev_user"
	config.DB_PASSWORD = "!rvqii"
	// config.DB_DATABASE = "dev_database"

	config.JWT_ISSUER = "dev_issuer"
	config.JWT_SECRET = "dev_secret"

	config.ISP_ML_SERVICE_HOST = "localhost"
	config.ISP_ML_SERVICE_PORT = 50051
}

func test(config *Config) {
	log.Println("Using test environment, loading config...")
	panic("config: not implemented")
}

func prod(config *Config) {
	log.Println("Using prod environment, loading config...")
	gin.SetMode(gin.ReleaseMode)
	panic("config: not implemented")
}

func loadConfig() {
	config = Config{
		ENV:  ENV,
		HOST: "localhost",
		PORT: 8080,
	}

	switch ENV {
	case "test":
		test(&config)
	case "prod":
		prod(&config)
	default:
		config.ENV = "dev"
		dev(&config)
	}

	if os.Getenv("HOST") != "" {
		config.HOST = os.Getenv("HOST")
	}
	if os.Getenv("PORT") != "" {
		config.PORT, _ = strconv.Atoi(os.Getenv("PORT"))
	}

	config.URL = fmt.Sprintf("%s:%d", config.HOST, config.PORT)
	// config.DB_CONN_STR = fmt.Sprintf(
	// 	"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
	// 	config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_PASSWORD, config.DB_DATABASE)
	config.ISP_ML_SERVICE_URL = fmt.Sprintf("%s:%d", config.ISP_ML_SERVICE_HOST, config.ISP_ML_SERVICE_PORT)
}
