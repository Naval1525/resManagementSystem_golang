package controller

import (
	"context"
	"resManagement/database"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := menuCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve menus"})
			return
		}
		var allMenus []bson.M
		if err = result.All(ctx, &allMenus); err != nil {
			c.JSON(500, gin.H{"error": "Failed to decode menus"})
			return
		}
		c.JSON(200, gin.H{"menus": allMenus})

	}
}
func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Get Users"})
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Create Food"})
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Update Food"})
	}
}
