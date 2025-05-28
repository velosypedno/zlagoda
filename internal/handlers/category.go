package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/velosypedno/zlagoda/internal/models"
)

type categoryCreator interface {
	CreateCategory(c models.CategoryCreate) error
}

func NewCategoryCreatePOSTHandler(service categoryCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		type request struct {
			Name string `json:"name" binding:"required"`
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}
		model := models.CategoryCreate{
			Name: req.Name,
		}
		if err := service.CreateCategory(model); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category: " + err.Error()})
			return
		}

		c.Status(http.StatusCreated)
	}
}

type categoryReader interface {
	GetCategoryByID(id int) (models.CategoryRetrieve, error)
}

func NewCategoryRetrieveGETHandler(service categoryReader) gin.HandlerFunc {
	return func(c *gin.Context) {
		type response struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		category, err := service.GetCategoryByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found: " + err.Error()})
			return
		}

		resp := response{
			ID:   category.ID,
			Name: category.Name,
		}

		c.JSON(http.StatusOK, resp)
	}
}
