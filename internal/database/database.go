package database

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db          *gorm.DB
	redisClient *redis.Client
	onec        sync.Once
)

func GetDB() (*gorm.DB, *redis.Client) {
	onec.Do(func() {
		initializePostgress()
		initializeRedis()
	})
	return db, redisClient
}

func initializePostgress() {
	err := godotenv.Load("dbPostgress.env")
	if err != nil {
		fmt.Println("Err loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	db = connection
}

func initializeRedis() {
	err := godotenv.Load("dbRedis.env")
	if err != nil {
		fmt.Println("Err loading .env file")
	}

	dbAddr := os.Getenv("DB_ADDR")
	dbPassword := os.Getenv("DB_PASSWORDD")
	db, _ := strconv.Atoi(os.Getenv("DB"))

	client := redis.NewClient(&redis.Options{
		Addr:     dbAddr,
		Password: dbPassword,
		DB:       db,
	})

	redisClient = client
}
