package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/MHSaeedkia/store-by-gin-redis-postgress/internal/module"
	"github.com/MHSaeedkia/store-by-gin-redis-postgress/internal/product"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func productIsExist(products []module.Products, name string) bool {
	for _, product := range products {
		if product.ProductName == name {
			return true
		}
	}
	return false
}

func Insert(c *gin.Context) {
	name := strings.ToLower(c.PostForm("name"))
	count, err := strconv.Atoi(c.PostForm("count"))
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		return
	}
	price, err := strconv.ParseFloat(c.PostForm("price"), 64)

	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		return
	}
	if name == "" || c.PostForm("count") == "" || c.PostForm("price") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "name or price of count is empty , paarameter should not be empty",
		})
		return
	}
	if price == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "price could not be 0",
		})
		return
	}
	err, products := product.GetAllProduct()
	if productIsExist(products, name) {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": fmt.Sprintf("product %s is exist in store , you can update it by end point /updateBN and /updadeBI", name),
		})
		return
	}
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		return
	}

	err, lastId = product.InsertProduct(lastId, name, price, count)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": fmt.Sprint(fmt.Errorf(err.Error())),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("product %s added to store", name),
	})
}

func UpdateByName(c *gin.Context) {
	name := strings.ToLower(c.PostForm("name"))
	count, err := strconv.Atoi(c.PostForm("count"))
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		return
	}
	price, err := strconv.ParseFloat(c.PostForm("price"), 64)

	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		return
	}
	if name == "" || c.PostForm("count") == "" || c.PostForm("price") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "name or price of count is empty , paarameter should not be empty",
		})
		return
	}
	if price == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "price could not be 0",
		})
		return
	}
	err = product.UpdateProductByName(name, price, count)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": fmt.Sprint(fmt.Errorf(err.Error())),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("product %s is updated", name),
	})
}

func UpdateById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		return
	}
	name := strings.ToLower(c.PostForm("name"))
	count, err := strconv.Atoi(c.PostForm("count"))
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		return
	}
	price, err := strconv.ParseFloat(c.PostForm("price"), 64)

	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		return
	}
	if name == "" || c.PostForm("count") == "" || c.PostForm("price") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "name or price of count is empty , paarameter should not be empty",
		})
		return
	}
	if price == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "price could not be 0",
		})
		return
	}
	err = product.UpdateProductById(id, name, price, count)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": fmt.Sprint(fmt.Errorf(err.Error())),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("product %s is updated", name),
	})
}

func GetByName(c *gin.Context) {
	name := strings.ToLower(c.Param("name"))
	response, err := product.GetProductByName(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": fmt.Sprint(fmt.Errorf(err.Error())),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			fmt.Sprintf("Response : %s", response.Source): response.Data,
		})
	}
}

func GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": fmt.Sprint(fmt.Errorf(err.Error())),
		})
		return
	}
	response, err := product.GetProductById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": fmt.Sprint(fmt.Errorf(err.Error())),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			fmt.Sprintf("Response : %s", response.Source): response.Data,
		})
	}
}

func RemoveByName(c *gin.Context) {
	name := strings.ToLower(c.Param("name"))
	err := product.RemoveProductByName(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": fmt.Sprint(fmt.Errorf(err.Error())),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("product %s is deleted", name),
		})
	}
}

func RemoveById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": fmt.Sprint(fmt.Errorf(err.Error())),
		})
		return
	}
	err = product.RemoveProductById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": fmt.Sprint(fmt.Errorf(err.Error())),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("product with id %v is deleted", id),
		})
	}
}

var (
	loginVar bool
	lastId   int
)

func Login(c *gin.Context) {
	loginVar = true
	c.JSON(http.StatusOK, gin.H{
		"message": "Authorized",
	})
	id, err := product.GetLastId()
	if err != nil {
		fmt.Errorf(err.Error())
	}
	lastId = id
	fmt.Println(lastId)
}

var secretKey = []byte("secret-key")

func createToken(userName string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": userName,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func checkUser(username, password string) (string, bool) {
	if username == "Chek" && password == "123456" {
		tokenString, err := createToken(username)
		if err != nil {
			fmt.Println(err)
			return "", false
		}
		return tokenString, true
	} else {
		return "", false
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.String() == "/login" {
			username := c.Request.Header.Get("username")
			password := c.Request.Header.Get("password")
			tokenString, valid := checkUser(username, password)
			if !valid {
				c.JSON(http.StatusUnauthorized, gin.H{"err": "Unauthorized"})
				c.Abort()
				return
			}
			err := verifyToken(tokenString)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"err": "Unauthorized"})
				c.Abort()
				return
			}
		} else if c.Request.URL.String() != "/login" {
			if loginVar {
				fmt.Println("Status OK")
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "First you have to login by /login end point",
				})
				c.Abort()
			}
		}
		c.Next()
	}
}
