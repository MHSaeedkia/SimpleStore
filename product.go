package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

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

type Products struct {
	ProductId    int     `json:"product_id"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
	ProductCount int     `json:"product_count"`
}

type JsonResponse struct {
	Data   Products `json:"data"`
	Source string   `json:"source"`
}

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

func getProductByName(name string) (*JsonResponse, error) {
	db, redisClient = GetDB()

	cachedProducts, err := redisClient.Get("products").Bytes()
	response := JsonResponse{}

	if err != nil {
		dbProducts, err := fetchFromDb(db)
		if err != nil {
			return nil, err
		}

		cachedProducts, err = json.Marshal(dbProducts)
		if err != nil {
			return nil, err
		}

		err = redisClient.Set("products", cachedProducts, 20*time.Second).Err()
		if err != nil {
			return nil, err
		}
		for _, product := range dbProducts {
			if product.ProductName == name {
				response = JsonResponse{Data: product, Source: "PostgreSQL"}
				return &response, err
			}
		}
	}
	products := []Products{}
	err = json.Unmarshal(cachedProducts, &products)
	if err != nil {
		return nil, err
	}
	for _, product := range products {
		if product.ProductName == name {
			response = JsonResponse{Data: product, Source: "Redis Cache"}
			return &response, nil
		}
	}

	return nil, fmt.Errorf("this product is not exist")

}

func getProductById(id int) (*JsonResponse, error) {
	db, redisClient = GetDB()

	cachedProducts, err := redisClient.Get("products").Bytes()
	response := JsonResponse{}

	if err != nil {
		dbProducts, err := fetchFromDb(db)
		if err != nil {
			return nil, err
		}

		cachedProducts, err = json.Marshal(dbProducts)
		if err != nil {
			return nil, err
		}

		err = redisClient.Set("products", cachedProducts, 5*time.Second).Err()
		if err != nil {
			return nil, err
		}
		for _, product := range dbProducts {
			if product.ProductId == id {
				response = JsonResponse{Data: product, Source: "PostgreSQL"}
				return &response, err
			}
		}
	}
	products := []Products{}
	err = json.Unmarshal(cachedProducts, &products)
	if err != nil {
		return nil, err
	}
	for _, product := range products {
		if product.ProductId == id {
			response = JsonResponse{Data: product, Source: "Redis Cache"}
			return &response, nil
		}
	}

	return nil, fmt.Errorf("product with this id is not exist")

}

func fetchFromDb(db *gorm.DB) ([]Products, error) {
	products := []Products{}
	db.Find(&products)
	if products != nil {
		return products, nil
	}
	return nil, fmt.Errorf("this product is not exist in your primary database")
}

func insertProduct(id int, name string, price float64, count int) (error, int) {
	db, _ = GetDB()

	product := Products{
		ProductId:    id + 1,
		ProductName:  name,
		ProductPrice: price,
		ProductCount: count,
	}

	err := db.Create(&product).Error
	if err != nil {
		return err, id
	}

	err = db.Last(&product).Error
	if err != nil {
		return err, id
	}
	id = product.ProductId

	return nil, id + 1
}

func updateProductByName(name string, price float64, count int) error {
	db, _ = GetDB()

	product := Products{
		ProductName:  name,
		ProductPrice: price,
		ProductCount: count,
	}

	err := db.Model(&Products{}).Where("product_name= ?", name).Updates(product).Error
	if err != nil {
		return err
	}

	return nil
}

func updateProductById(id int, name string, price float64, count int) error {
	db, _ = GetDB()

	product := Products{
		ProductName:  name,
		ProductPrice: price,
		ProductCount: count,
	}

	err := db.Model(&Products{}).Where("product_id= ?", id).Updates(product).Error
	if err != nil {
		return err
	}

	return nil
}

func removeProductByName(name string) error {
	db, _ = GetDB()

	err := db.Where("product_name= ?", name).Delete(&Products{}).Error
	if err != nil {
		return err
	}

	return nil
}

func removeProductById(id int) error {
	db, _ = GetDB()

	err := db.Where("product_id= ?", id).Delete(&Products{}).Error
	if err != nil {
		return err
	}

	return nil
}

func getLastId() (int, error) {
	db, _ = GetDB()

	product := Products{}

	err := db.Last(&product).Error
	if err != nil {
		return -1, nil
	}

	id := product.ProductId

	return id, nil

}

func getAllProduct() (error, []Products) {
	db, _ = GetDB()

	products := []Products{}

	err := db.Find(&products).Error
	if err != nil {
		return err, nil
	}
	return nil, products

}
