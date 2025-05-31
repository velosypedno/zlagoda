package ioc

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/velosypedno/zlagoda/internal/config"
	"github.com/velosypedno/zlagoda/internal/handlers"
	"github.com/velosypedno/zlagoda/internal/repos"
	"github.com/velosypedno/zlagoda/internal/services"
)

type HandlerContainer struct {
	CategoryCreatePOSTHandler   gin.HandlerFunc
	CategoryRetrieveGETHandler  gin.HandlerFunc
	CategoriesListGETHandler    gin.HandlerFunc
	CategoryDeleteDELETEHandler gin.HandlerFunc
	CategoryUpdatePATCHHandler  gin.HandlerFunc
}

func BuildHandlerContainer(c *config.Config) *HandlerContainer {
	db, err := sql.Open(c.DB_DRIVER, c.DB_DSN)
	if err != nil {
		log.Fatal(err)
	}
	categoryRepo := repos.NewCategoryRepo(db)
	categoryService := services.NewCategoryService(categoryRepo)

	return &HandlerContainer{
		CategoryCreatePOSTHandler:   handlers.NewCategoryCreatePOSTHandler(categoryService),
		CategoryRetrieveGETHandler:  handlers.NewCategoryRetrieveGETHandler(categoryService),
		CategoriesListGETHandler:    handlers.NewCategoryListGETHandler(categoryService),
		CategoryDeleteDELETEHandler: handlers.NewCategoryDeleteDELETEHandler(categoryService),
		CategoryUpdatePATCHHandler:  handlers.NewCategoryUpdatePATCHHandler(categoryService),
	}
}
