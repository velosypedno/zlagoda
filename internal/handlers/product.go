package handlers

import (
	"log"
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
		log.Printf("[ProductCreatePOST] Starting product creation request")
		
		var req struct {
			Name            string `json:"name"            binding:"required"`
			Characteristics string `json:"characteristics" binding:"required"`
			CategoryID      int    `json:"category_id"     binding:"required"`
		}
		
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("[ProductCreatePOST] BindJSON error: %v", err)
			log.Printf("[ProductCreatePOST] Request validation failed: %+v", req)
			log.Printf("[ProductCreatePOST] Content-Type: %s", c.GetHeader("Content-Type"))
			log.Printf("[ProductCreatePOST] Content-Length: %s", c.GetHeader("Content-Length"))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		// Log the parsed request data
		log.Printf("[ProductCreatePOST] Parsed request data: Name=%s, Characteristics=%s, CategoryID=%d", 
			req.Name, req.Characteristics, req.CategoryID)

		// Validate name
		if req.Name == "" {
			log.Printf("[ProductCreatePOST] Empty product name")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product name cannot be empty"})
			return
		}

		// Validate characteristics
		if req.Characteristics == "" {
			log.Printf("[ProductCreatePOST] Empty product characteristics")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product characteristics cannot be empty"})
			return
		}

		// Validate category ID
		if req.CategoryID <= 0 {
			log.Printf("[ProductCreatePOST] Invalid category ID: %d (must be positive)", req.CategoryID)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Category ID must be a positive integer"})
			return
		}

		model := models.ProductCreate{
			Name:            req.Name,
			Characteristics: req.Characteristics,
			CategoryID:      req.CategoryID,
		}
		
		log.Printf("[ProductCreatePOST] Calling service.CreateProduct with model: %+v", model)
		
		id, err := service.CreateProduct(model)
		if err != nil {
			log.Printf("[ProductCreatePOST] Service error: %v", err)
			log.Printf("[ProductCreatePOST] Service error details - Name: %s, CategoryID: %d, Error: %s", 
				req.Name, req.CategoryID, err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product: " + err.Error()})
			return
		}
		
		log.Printf("[ProductCreatePOST] Successfully created product with ID: %d", id)
		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

type productReader interface {
	GetProductByID(id int) (models.ProductRetrieve, error)
	GetProducts() ([]models.ProductRetrieve, error)
	GetProductsByCategory(categoryID int) ([]models.ProductRetrieve, error)
	GetProductsByName(name string) ([]models.ProductRetrieve, error)
}

func NewProductRetrieveGETHandler(service productReader) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Printf("[ProductRetrieveGET] Invalid ID parameter: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		p, err := service.GetProductByID(id)
		if err != nil {
			log.Printf("[ProductRetrieveGET] Service error for ID %d: %v", id, err)
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
			log.Printf("[ProductsListGET] Service error: %v", err)
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

func NewProductsByCategoryGETHandler(service productReader) gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryIDStr := c.Param("category_id")
		categoryID, err := strconv.Atoi(categoryIDStr)
		if err != nil {
			log.Printf("[ProductsByCategoryGET] Invalid category ID: %s", categoryIDStr)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
			return
		}

		items, err := service.GetProductsByCategory(categoryID)
		if err != nil {
			log.Printf("[ProductsByCategoryGET] Service error for category ID %d: %v", categoryID, err)
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

func NewProductsByNameGETHandler(service productReader) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			log.Printf("[ProductsByNameGET] Missing name parameter")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing name parameter"})
			return
		}

		items, err := service.GetProductsByName(name)
		if err != nil {
			log.Printf("[ProductsByNameGET] Service error for name '%s': %v", name, err)
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
			log.Printf("[ProductUpdatePATCH] Invalid ID parameter: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var req struct {
			Name            *string `json:"name"`
			Characteristics *string `json:"characteristics"`
			CategoryID      *int    `json:"category_id"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("[ProductUpdatePATCH] BindJSON error for ID %d: %v", id, err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		model := models.ProductUpdate{
			Name:            req.Name,
			Characteristics: req.Characteristics,
			CategoryID:      req.CategoryID,
		}

		if err := service.UpdateProduct(id, model); err != nil {
			log.Printf("[ProductUpdatePATCH] Service error for ID %d: %v", id, err)
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
			log.Printf("[ProductDeleteDELETE] Invalid ID parameter: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		if err := service.DeleteProduct(id); err != nil {
			log.Printf("[ProductDeleteDELETE] Service error for ID %d: %v", id, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product: " + err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
	}
}
