package product

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/MHSaeedkia/store-by-gin-redis-postgress/internal/database"
	"github.com/MHSaeedkia/store-by-gin-redis-postgress/internal/module"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	db          *gorm.DB
	redisClient *redis.Client
	onec        sync.Once
)

func GetProductByName(name string) (*module.JsonResponse, error) {
	db, redisClient = database.GetDB()

	cachedProducts, err := redisClient.Get("products").Bytes()
	response := module.JsonResponse{}

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
				response = module.JsonResponse{Data: product, Source: "PostgreSQL"}
				return &response, err
			}
		}
	}
	products := []module.Products{}
	err = json.Unmarshal(cachedProducts, &products)
	if err != nil {
		return nil, err
	}
	for _, product := range products {
		if product.ProductName == name {
			response = module.JsonResponse{Data: product, Source: "Redis Cache"}
			return &response, nil
		}
	}

	return nil, fmt.Errorf("this product is not exist")

}

func GetProductById(id int) (*module.JsonResponse, error) {
	db, redisClient = database.GetDB()

	cachedProducts, err := redisClient.Get("products").Bytes()
	response := module.JsonResponse{}

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
				response = module.JsonResponse{Data: product, Source: "PostgreSQL"}
				return &response, err
			}
		}
	}
	products := []module.Products{}
	err = json.Unmarshal(cachedProducts, &products)
	if err != nil {
		return nil, err
	}
	for _, product := range products {
		if product.ProductId == id {
			response = module.JsonResponse{Data: product, Source: "Redis Cache"}
			return &response, nil
		}
	}

	return nil, fmt.Errorf("product with this id is not exist")

}

func fetchFromDb(db *gorm.DB) ([]module.Products, error) {
	products := []module.Products{}
	db.Find(&products)
	if products != nil {
		return products, nil
	}
	return nil, fmt.Errorf("this product is not exist in your primary database")
}

func InsertProduct(id int, name string, price float64, count int) (error, int) {
	db, _ = database.GetDB()

	product := module.Products{
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

func UpdateProductByName(name string, price float64, count int) error {
	db, _ = database.GetDB()

	product := module.Products{
		ProductName:  name,
		ProductPrice: price,
		ProductCount: count,
	}

	err := db.Model(&module.Products{}).Where("product_name= ?", name).Updates(product).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateProductById(id int, name string, price float64, count int) error {
	db, _ = database.GetDB()

	product := module.Products{
		ProductName:  name,
		ProductPrice: price,
		ProductCount: count,
	}

	err := db.Model(&module.Products{}).Where("product_id= ?", id).Updates(product).Error
	if err != nil {
		return err
	}

	return nil
}

func RemoveProductByName(name string) error {
	db, _ = database.GetDB()

	err := db.Where("product_name= ?", name).Delete(&module.Products{}).Error
	if err != nil {
		return err
	}

	return nil
}

func RemoveProductById(id int) error {
	db, _ = database.GetDB()

	err := db.Where("product_id= ?", id).Delete(&module.Products{}).Error
	if err != nil {
		return err
	}

	return nil
}

func GetLastId() (int, error) {
	db, _ = database.GetDB()

	product := module.Products{}

	err := db.Last(&product).Error
	if err != nil {
		return -1, nil
	}

	id := product.ProductId

	return id, nil

}

func GetAllProduct() (error, []module.Products) {
	db, _ = database.GetDB()

	products := []module.Products{}

	err := db.Find(&products).Error
	if err != nil {
		return err, nil
	}
	return nil, products

}
