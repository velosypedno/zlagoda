package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/velosypedno/zlagoda/internal/models"
)

type categoryCreator interface {
	CreateCategory(c models.CategoryCreate) (int, error)
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
		id, err := service.CreateCategory(model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category: " + err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

type categoryReader interface {
	GetCategoryByID(id int) (models.CategoryRetrieve, error)
	GetCategories() ([]models.CategoryRetrieve, error)
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

func NewCategoryListGETHandler(service categoryReader) gin.HandlerFunc {
	type responseItem struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	return func(c *gin.Context) {
		categories, err := service.GetCategories()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve categories: " + err.Error()})
			return
		}

		var resp []responseItem
		for _, cat := range categories {
			resp = append(resp, responseItem{
				ID:   cat.ID,
				Name: cat.Name,
			})
		}

		c.JSON(http.StatusOK, resp)
	}
}

type categoryRemover interface {
	DeleteCategory(id int) error
}

func NewCategoryDeleteDELETEHandler(service categoryRemover) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		err = service.DeleteCategory(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
	}
}

type categoryUpdater interface {
	UpdateCategory(id int, model models.CategoryUpdate) error
	GetCategoryByID(id int) (models.CategoryRetrieve, error)
}

func NewCategoryUpdatePATCHHandler(service categoryUpdater) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		type request struct {
			Name *string `json:"name"`
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		categoryCurrentState, err := service.GetCategoryByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found: " + err.Error()})
			return
		}
		if req.Name == nil {
			req.Name = &categoryCurrentState.Name
		}

		model := models.CategoryUpdate{
			Name: req.Name,
		}

		err = service.UpdateCategory(id, model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
	}
}
