package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/velosypedno/zlagoda/internal/models"
	"github.com/velosypedno/zlagoda/internal/utils"
)

type storeProductCreator interface {
	CreateStoreProduct(sp models.StoreProductCreate) (string, error)
}

func NewStoreProductCreatePOSTHandler(service storeProductCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		type request struct {
			UPC                string  `json:"upc" binding:"required,len=12"`
			UPCProm            *string `json:"upc_prom" binding:"omitempty,len=12"`
			ProductID          int     `json:"product_id" binding:"required"`
			SellingPrice       float64 `json:"selling_price" binding:"required,gte=0"`
			ProductsNumber     int     `json:"products_number" binding:"required,gte=0"`
			PromotionalProduct bool    `json:"promotional_product"`
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		if !utils.IsAmountValid(req.SellingPrice) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: invalid selling price"})
			return
		}

		model := models.StoreProductCreate{
			UPC:                req.UPC,
			UPCProm:            req.UPCProm,
			ProductID:          req.ProductID,
			SellingPrice:       req.SellingPrice,
			ProductsNumber:     req.ProductsNumber,
			PromotionalProduct: req.PromotionalProduct,
		}

		upc, err := service.CreateStoreProduct(model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create store product: " + err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"upc": upc})
	}
}

type storeProductReader interface {
	GetStoreProductByUPC(upc string) (models.StoreProductRetrieve, error)
	GetStoreProducts() ([]models.StoreProductRetrieve, error)
	GetStoreProductsWithDetails() ([]models.StoreProductWithDetails, error)
	GetStoreProductsByProductID(productID int) ([]models.StoreProductRetrieve, error)
	GetPromotionalProducts() ([]models.StoreProductRetrieve, error)
}

func NewStoreProductRetrieveGETHandler(service storeProductReader) gin.HandlerFunc {
	return func(c *gin.Context) {
		type response struct {
			UPC                string  `json:"upc"`
			UPCProm            *string `json:"upc_prom"`
			ProductID          int     `json:"product_id"`
			SellingPrice       float64 `json:"selling_price"`
			ProductsNumber     int     `json:"products_number"`
			PromotionalProduct bool    `json:"promotional_product"`
		}

		upc := c.Param("upc")
		if len(upc) != 12 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UPC format"})
			return
		}

		storeProduct, err := service.GetStoreProductByUPC(upc)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Store product not found: " + err.Error()})
			return
		}

		resp := response{
			UPC:                storeProduct.UPC,
			UPCProm:            storeProduct.UPCProm,
			ProductID:          storeProduct.ProductID,
			SellingPrice:       storeProduct.SellingPrice,
			ProductsNumber:     storeProduct.ProductsNumber,
			PromotionalProduct: storeProduct.PromotionalProduct,
		}

		c.JSON(http.StatusOK, resp)
	}
}

func NewStoreProductsListGETHandler(service storeProductReader) gin.HandlerFunc {
	type responseItem struct {
		UPC                string  `json:"upc"`
		UPCProm            *string `json:"upc_prom"`
		ProductID          int     `json:"product_id"`
		SellingPrice       float64 `json:"selling_price"`
		ProductsNumber     int     `json:"products_number"`
		PromotionalProduct bool    `json:"promotional_product"`
	}

	return func(c *gin.Context) {
		storeProducts, err := service.GetStoreProducts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve store products: " + err.Error()})
			return
		}

		var resp []responseItem
		for _, sp := range storeProducts {
			resp = append(resp, responseItem{
				UPC:                sp.UPC,
				UPCProm:            sp.UPCProm,
				ProductID:          sp.ProductID,
				SellingPrice:       sp.SellingPrice,
				ProductsNumber:     sp.ProductsNumber,
				PromotionalProduct: sp.PromotionalProduct,
			})
		}

		c.JSON(http.StatusOK, resp)
	}
}

func NewStoreProductsWithDetailsListGETHandler(service storeProductReader) gin.HandlerFunc {
	return func(c *gin.Context) {
		storeProducts, err := service.GetStoreProductsWithDetails()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve store products with details: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, storeProducts)
	}
}

func NewStoreProductsByProductIDGETHandler(service storeProductReader) gin.HandlerFunc {
	type responseItem struct {
		UPC                string  `json:"upc"`
		UPCProm            *string `json:"upc_prom"`
		ProductID          int     `json:"product_id"`
		SellingPrice       float64 `json:"selling_price"`
		ProductsNumber     int     `json:"products_number"`
		PromotionalProduct bool    `json:"promotional_product"`
	}

	return func(c *gin.Context) {
		productIDStr := c.Param("product_id")
		productID, err := strconv.Atoi(productIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}

		storeProducts, err := service.GetStoreProductsByProductID(productID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve store products: " + err.Error()})
			return
		}

		var resp []responseItem
		for _, sp := range storeProducts {
			resp = append(resp, responseItem{
				UPC:                sp.UPC,
				UPCProm:            sp.UPCProm,
				ProductID:          sp.ProductID,
				SellingPrice:       sp.SellingPrice,
				ProductsNumber:     sp.ProductsNumber,
				PromotionalProduct: sp.PromotionalProduct,
			})
		}

		c.JSON(http.StatusOK, resp)
	}
}

func NewPromotionalProductsGETHandler(service storeProductReader) gin.HandlerFunc {
	type responseItem struct {
		UPC                string  `json:"upc"`
		UPCProm            *string `json:"upc_prom"`
		ProductID          int     `json:"product_id"`
		SellingPrice       float64 `json:"selling_price"`
		ProductsNumber     int     `json:"products_number"`
		PromotionalProduct bool    `json:"promotional_product"`
	}

	return func(c *gin.Context) {
		storeProducts, err := service.GetPromotionalProducts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve promotional products: " + err.Error()})
			return
		}

		var resp []responseItem
		for _, sp := range storeProducts {
			resp = append(resp, responseItem{
				UPC:                sp.UPC,
				UPCProm:            sp.UPCProm,
				ProductID:          sp.ProductID,
				SellingPrice:       sp.SellingPrice,
				ProductsNumber:     sp.ProductsNumber,
				PromotionalProduct: sp.PromotionalProduct,
			})
		}

		c.JSON(http.StatusOK, resp)
	}
}

type storeProductUpdater interface {
	UpdateStoreProduct(upc string, sp models.StoreProductUpdate) error
	GetStoreProductByUPC(upc string) (models.StoreProductRetrieve, error)
}

func NewStoreProductUpdatePATCHHandler(service storeProductUpdater) gin.HandlerFunc {
	return func(c *gin.Context) {
		upc := c.Param("upc")
		if len(upc) != 12 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UPC format"})
			return
		}

		type request struct {
			UPCProm            *string  `json:"upc_prom" binding:"omitempty,len=12"`
			ProductID          *int     `json:"product_id"`
			SellingPrice       *float64 `json:"selling_price" binding:"omitempty,gte=0"`
			ProductsNumber     *int     `json:"products_number" binding:"omitempty,gte=0"`
			PromotionalProduct *bool    `json:"promotional_product"`
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		// Check if store product exists
		_, err := service.GetStoreProductByUPC(upc)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Store product not found: " + err.Error()})
			return
		}

		if req.SellingPrice != nil && !utils.IsAmountValid(*req.SellingPrice) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: invalid selling price"})
			return
		}

		model := models.StoreProductUpdate{
			UPCProm:            req.UPCProm,
			ProductID:          req.ProductID,
			SellingPrice:       req.SellingPrice,
			ProductsNumber:     req.ProductsNumber,
			PromotionalProduct: req.PromotionalProduct,
		}

		err = service.UpdateStoreProduct(upc, model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update store product: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Store product updated successfully"})
	}
}

type storeProductRemover interface {
	DeleteStoreProduct(upc string) error
}

func NewStoreProductDeleteDELETEHandler(service storeProductRemover) gin.HandlerFunc {
	return func(c *gin.Context) {
		upc := c.Param("upc")
		if len(upc) != 12 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UPC format"})
			return
		}

		err := service.DeleteStoreProduct(upc)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete store product: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Store product deleted successfully"})
	}
}

type storeProductInventoryManager interface {
	UpdateProductQuantity(upc string, quantityChange int) error
	CheckStockAvailability(upc string, requiredQuantity int) (bool, error)
}

func NewStoreProductQuantityUpdatePATCHHandler(service storeProductInventoryManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		upc := c.Param("upc")
		if len(upc) != 12 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UPC format"})
			return
		}

		type request struct {
			QuantityChange int `json:"quantity_change" binding:"required"`
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		err := service.UpdateProductQuantity(upc, req.QuantityChange)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update quantity: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product quantity updated successfully"})
	}
}

func NewStoreProductStockCheckGETHandler(service storeProductInventoryManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		upc := c.Param("upc")
		if len(upc) != 12 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UPC format"})
			return
		}

		requiredQuantityStr := c.Query("quantity")
		if requiredQuantityStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing quantity parameter"})
			return
		}

		requiredQuantity, err := strconv.Atoi(requiredQuantityStr)
		if err != nil || requiredQuantity < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity parameter"})
			return
		}

		available, err := service.CheckStockAvailability(upc, requiredQuantity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check stock: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"upc":               upc,
			"required_quantity": requiredQuantity,
			"stock_available":   available,
		})
	}
}
