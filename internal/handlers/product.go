package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/velosypedno/zlagoda/internal/models"
)

type productCreator interface {
	CreateProduct(p models.ProductCreate) (int, error)
}

func NewProductCreatePOSTHandler(service productCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Name            string `json:"name"            binding:"required"`
			Characteristics string `json:"characteristics" binding:"required"`
			CategoryID      int    `json:"category_id"     binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		model := models.ProductCreate{
			Name:            req.Name,
			Characteristics: req.Characteristics,
			CategoryID:      req.CategoryID,
		}
		id, err := service.CreateProduct(model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product: " + err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

type productReader interface {
	GetProductByID(id int) (models.ProductRetrieve, error)
	GetProducts() ([]models.ProductRetrieve, error)
}

func NewProductRetrieveGETHandler(service productReader) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		p, err := service.GetProductByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":              p.ID,
			"name":            p.Name,
			"characteristics": p.Characteristics,
			"category_id":     p.CategoryID,
		})
	}
}

func NewProductsListGETHandler(service productReader) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := service.GetProducts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products: " + err.Error()})
			return
		}

		resp := make([]gin.H, 0, len(items))
		for _, p := range items {
			resp = append(resp, gin.H{
				"id":              p.ID,
				"name":            p.Name,
				"characteristics": p.Characteristics,
				"category_id":     p.CategoryID,
			})
		}
		c.JSON(http.StatusOK, resp)
	}
}

type productUpdater interface {
	UpdateProduct(id int, model models.ProductUpdate) error
}

func NewProductUpdatePATCHHandler(service productUpdater) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var req struct {
			Name            *string `json:"name"`
			Characteristics *string `json:"characteristics"`
			CategoryID      *int    `json:"category_id"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		model := models.ProductUpdate{
			Name:            req.Name,
			Characteristics: req.Characteristics,
			CategoryID:      req.CategoryID,
		}

		if err := service.UpdateProduct(id, model); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product: " + err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
	}
}

type productRemover interface {
	DeleteProduct(id int) error
}

func NewProductDeleteDELETEHandler(service productRemover) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		if err := service.DeleteProduct(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product: " + err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
	}
}
