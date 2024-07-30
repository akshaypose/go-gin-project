package handlers

import (
	"encoding/json"
	"go-gin-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Products struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
}

type APIResponse struct {
	ProductList []Products `json:"products"`
}

func GetAllProducts(c *gin.Context) {
	body, err := utils.MakeCurlCall("GET", "https://dummyjson.com/products")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to make API Call"})
		return
	}

	var data APIResponse
	if err := json.Unmarshal(body, &data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to unmarshal JSON"})
		return
	}

	if data.ProductList == nil {
		c.JSON(http.StatusOK, gin.H{"Data": "No Products Found."})
		return
	}
	var listOfProducts []Products = data.ProductList
	c.JSON(http.StatusOK, gin.H{"Data": listOfProducts})
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	body, err := utils.MakeCurlCall("GET", "https://dummyjson.com/products/"+id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to make API Call"})
		return
	}

	var data Products
	if err := json.Unmarshal(body, &data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to unmarshal JSON"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})

}
