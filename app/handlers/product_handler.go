package handlers

import (
	"pcshop/app/models"
	"pcshop/config"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
    var products []models.Product
    config.DB.Find(&products)
    c.JSON(200, products)
}

func CreateProduct(c *gin.Context) {
    var product models.Product
    c.BindJSON(&product)
    config.DB.Create(&product)
    c.JSON(201, product)
}