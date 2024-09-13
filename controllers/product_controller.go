package controllers

import (
	"context"
	"go-inventory/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateProduct handles POST requests to create a new product
func CreateProduct(c *gin.Context) {
	collection := utils.GetMongoCollection()
	var product bson.M
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Insert the product into the collection
	_, err := collection.InsertOne(context.TODO(), product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully"})
}

// GetProductByID handles GET requests to retrieve a product by ID
func GetProductByID(c *gin.Context) {
	collection := utils.GetMongoCollection()
	id := c.Param("id")

	// Convert the ID from the request to an ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Find the product by ID
	var product bson.M
	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, product)
}

// GetAllProducts handles GET requests to retrieve all products
func GetAllProducts(c *gin.Context) {
	collection := utils.GetMongoCollection()
	var products []bson.M
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err = cursor.All(context.TODO(), &products); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// UpdateProduct handles PUT requests to update an existing product
func UpdateProduct(c *gin.Context) {
	collection := utils.GetMongoCollection()
	id := c.Param("id")
	var updateData bson.M

	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Convert the ID from the request to an ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Update the product in the collection
	updateResult, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": updateData},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if updateResult.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

// DeleteProduct handles DELETE requests to remove a product by ID
func DeleteProduct(c *gin.Context) {
	collection := utils.GetMongoCollection()
	id := c.Param("id")

	// Convert the ID from the request to an ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Delete the product from the collection
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if deleteResult.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
