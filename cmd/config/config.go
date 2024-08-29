package config

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
	"os"
	"strconv"

	entity "github.com/eron97/bff-golang.git/src/model/entitys"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Cfg        *Config
	PrivateKey *rsa.PrivateKey
	Logger     *zap.Logger
)

type Config struct {
	DBDriver      string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	WebServerPort string
	JWTSecret     string
	JWTExpiresIn  int
}

func NewConfig() *Config {
	return Cfg
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Warning: Error loading .env file: %v", err)
	}

	Cfg = &Config{
		DBDriver:      os.Getenv("DB_DRIVER"),
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		WebServerPort: os.Getenv("WEB_SERVER_PORT"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
	}

	Cfg.JWTExpiresIn, _ = strconv.Atoi(os.Getenv("JWT_EXPIRES_IN"))

	rng := rand.Reader
	PrivateKey, err = rsa.GenerateKey(rng, 2048)
	if err != nil {
		log.Fatalf("Error generating RSA key: %v", err)
	}

	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "ts"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""

	Logger, err = config.Build()
	if err != nil {
		log.Fatalf("Error initializing zap logger: %v", err)
	}
	zap.ReplaceGlobals(Logger)
}

func NewDatabaseConnection() (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	switch Cfg.DBDriver {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			Cfg.DBUser,
			Cfg.DBPassword,
			Cfg.DBHost,
			Cfg.DBPort,
			Cfg.DBName)

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	case "sqlite":
		db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		if err != nil {
			return nil, err
		}

		db.AutoMigrate(&entity.CreateUser{})

	default:
		return nil, fmt.Errorf("unsupported database driver: %s", Cfg.DBDriver)
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}
