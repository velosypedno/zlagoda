package handlers

import (
	"log"
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
		log.Printf("[StoreProductCreatePOST] Starting store product creation request")
		
		type request struct {
			UPC                string  `json:"upc" binding:"required,len=12"`
			UPCProm            *string `json:"upc_prom" binding:"omitempty"`
			ProductID          int     `json:"product_id" binding:"required"`
			SellingPrice       float64 `json:"selling_price" binding:"required,gte=0"`
			ProductsNumber     int     `json:"products_number" binding:"required,gte=0"`
			PromotionalProduct bool    `json:"promotional_product"`
		}
		var req request
		
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("[StoreProductCreatePOST] BindJSON error: %v", err)
			log.Printf("[StoreProductCreatePOST] Request validation failed: %+v", req)
			log.Printf("[StoreProductCreatePOST] Content-Type: %s", c.GetHeader("Content-Type"))
			log.Printf("[StoreProductCreatePOST] Content-Length: %s", c.GetHeader("Content-Length"))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		// Handle empty string for UPCProm - convert to null
		if req.UPCProm != nil && *req.UPCProm == "" {
			log.Printf("[StoreProductCreatePOST] Converting empty UPCProm to null")
			req.UPCProm = nil
		}

		// Log the parsed request data
		log.Printf("[StoreProductCreatePOST] Parsed request data: UPC=%s, ProductID=%d, SellingPrice=%f, ProductsNumber=%d, PromotionalProduct=%t", 
			req.UPC, req.ProductID, req.SellingPrice, req.ProductsNumber, req.PromotionalProduct)
		if req.UPCProm != nil {
			log.Printf("[StoreProductCreatePOST] Promotional UPC: %s", *req.UPCProm)
		}

		// Validate UPC format
		if len(req.UPC) != 12 {
			log.Printf("[StoreProductCreatePOST] Invalid UPC length: %d (expected 12)", len(req.UPC))
			c.JSON(http.StatusBadRequest, gin.H{"error": "UPC must be exactly 12 characters"})
			return
		}

		// Validate promotional UPC if provided
		if req.UPCProm != nil && len(*req.UPCProm) != 12 {
			log.Printf("[StoreProductCreatePOST] Invalid promotional UPC length: %d (expected 12)", len(*req.UPCProm))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Promotional UPC must be exactly 12 characters"})
			return
		}

		// Validate product ID
		if req.ProductID <= 0 {
			log.Printf("[StoreProductCreatePOST] Invalid product ID: %d (must be positive)", req.ProductID)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID must be a positive integer"})
			return
		}

		// Validate selling price
		if req.SellingPrice < 0 {
			log.Printf("[StoreProductCreatePOST] Invalid selling price: %f (must be non-negative)", req.SellingPrice)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Selling price must be non-negative"})
			return
		}

		if !utils.IsAmountValid(req.SellingPrice) {
			log.Printf("[StoreProductCreatePOST] Invalid selling price format: %v", req.SellingPrice)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: invalid selling price"})
			return
		}

		// Validate products number
		if req.ProductsNumber < 0 {
			log.Printf("[StoreProductCreatePOST] Invalid products number: %d (must be non-negative)", req.ProductsNumber)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Products number must be non-negative"})
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

		log.Printf("[StoreProductCreatePOST] Calling service.CreateStoreProduct with model: %+v", model)
		
		upc, err := service.CreateStoreProduct(model)
		if err != nil {
			log.Printf("[StoreProductCreatePOST] Service error: %v", err)
			log.Printf("[StoreProductCreatePOST] Service error details - UPC: %s, ProductID: %d, Error: %s", 
				req.UPC, req.ProductID, err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create store product: " + err.Error()})
			return
		}

		log.Printf("[StoreProductCreatePOST] Successfully created store product with UPC: %s", upc)
		c.JSON(http.StatusCreated, gin.H{"upc": upc})
	}
}

type storeProductReader interface {
	GetStoreProductByUPC(upc string) (models.StoreProductRetrieve, error)
	GetStoreProducts() ([]models.StoreProductRetrieve, error)
	GetStoreProductsWithDetails() ([]models.StoreProductWithDetails, error)
	GetStoreProductsByProductID(productID int) ([]models.StoreProductRetrieve, error)
	GetPromotionalProducts() ([]models.StoreProductRetrieve, error)
	GetStoreProductsByCategory(categoryID int) ([]models.StoreProductWithDetails, error)
	GetStoreProductsByName(name string) ([]models.StoreProductWithDetails, error)
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
			log.Printf("[StoreProductRetrieveGET] Invalid UPC format: %s", upc)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UPC format"})
			return
		}

		storeProduct, err := service.GetStoreProductByUPC(upc)
		if err != nil {
			log.Printf("[StoreProductRetrieveGET] Service error for UPC %s: %v", upc, err)
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
			log.Printf("[StoreProductsListGET] Service error: %v", err)
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
			log.Printf("[StoreProductsWithDetailsListGET] Service error: %v", err)
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
			log.Printf("[StoreProductsByProductIDGET] Invalid product ID: %s", productIDStr)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}

		storeProducts, err := service.GetStoreProductsByProductID(productID)
		if err != nil {
			log.Printf("[StoreProductsByProductIDGET] Service error for product ID %d: %v", productID, err)
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
			log.Printf("[PromotionalProductsGET] Service error: %v", err)
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

func NewStoreProductsByCategoryGETHandler(service storeProductReader) gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryIDStr := c.Param("category_id")
		categoryID, err := strconv.Atoi(categoryIDStr)
		if err != nil {
			log.Printf("[StoreProductsByCategoryGET] Invalid category ID: %s", categoryIDStr)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
			return
		}

		storeProducts, err := service.GetStoreProductsByCategory(categoryID)
		if err != nil {
			log.Printf("[StoreProductsByCategoryGET] Service error for category ID %d: %v", categoryID, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve store products: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, storeProducts)
	}
}

func NewStoreProductsByNameGETHandler(service storeProductReader) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			log.Printf("[StoreProductsByNameGET] Missing name parameter")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing name parameter"})
			return
		}

		storeProducts, err := service.GetStoreProductsByName(name)
		if err != nil {
			log.Printf("[StoreProductsByNameGET] Service error for name '%s': %v", name, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve store products: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, storeProducts)
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
			log.Printf("[StoreProductUpdatePATCH] Invalid UPC format: %s", upc)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UPC format"})
			return
		}

		type request struct {
			UPCProm            *string  `json:"upc_prom" binding:"omitempty"`
			ProductID          *int     `json:"product_id"`
			SellingPrice       *float64 `json:"selling_price" binding:"omitempty,gte=0"`
			ProductsNumber     *int     `json:"products_number" binding:"omitempty,gte=0"`
			PromotionalProduct *bool    `json:"promotional_product"`
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("[StoreProductUpdatePATCH] BindJSON error for UPC %s: %v", upc, err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		// Handle empty string for UPCProm - convert to null
		if req.UPCProm != nil && *req.UPCProm == "" {
			log.Printf("[StoreProductUpdatePATCH] Converting empty UPCProm to null for UPC %s", upc)
			req.UPCProm = nil
		}

		// Check if store product exists
		_, err := service.GetStoreProductByUPC(upc)
		if err != nil {
			log.Printf("[StoreProductUpdatePATCH] Store product not found for UPC %s: %v", upc, err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Store product not found: " + err.Error()})
			return
		}

		// Validate promotional UPC if provided
		if req.UPCProm != nil && len(*req.UPCProm) != 12 {
			log.Printf("[StoreProductUpdatePATCH] Invalid promotional UPC length for UPC %s: %d (expected 12)", upc, len(*req.UPCProm))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Promotional UPC must be exactly 12 characters"})
			return
		}

		if req.SellingPrice != nil && !utils.IsAmountValid(*req.SellingPrice) {
			log.Printf("[StoreProductUpdatePATCH] Invalid selling price for UPC %s: %v", upc, *req.SellingPrice)
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
			log.Printf("[StoreProductUpdatePATCH] Service error for UPC %s: %v", upc, err)
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
			log.Printf("[StoreProductDeleteDELETE] Invalid UPC format: %s", upc)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UPC format"})
			return
		}

		err := service.DeleteStoreProduct(upc)
		if err != nil {
			log.Printf("[StoreProductDeleteDELETE] Service error for UPC %s: %v", upc, err)
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
			log.Printf("[StoreProductQuantityUpdatePATCH] Invalid UPC format: %s", upc)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UPC format"})
			return
		}

		type request struct {
			QuantityChange int `json:"quantity_change" binding:"required"`
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("[StoreProductQuantityUpdatePATCH] BindJSON error for UPC %s: %v", upc, err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		err := service.UpdateProductQuantity(upc, req.QuantityChange)
		if err != nil {
			log.Printf("[StoreProductQuantityUpdatePATCH] Service error for UPC %s, quantity change %d: %v", upc, req.QuantityChange, err)
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
			log.Printf("[StoreProductStockCheckGET] Invalid UPC format: %s", upc)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UPC format"})
			return
		}

		requiredQuantityStr := c.Query("quantity")
		if requiredQuantityStr == "" {
			log.Printf("[StoreProductStockCheckGET] Missing quantity parameter for UPC %s", upc)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing quantity parameter"})
			return
		}

		requiredQuantity, err := strconv.Atoi(requiredQuantityStr)
		if err != nil || requiredQuantity < 1 {
			log.Printf("[StoreProductStockCheckGET] Invalid quantity parameter for UPC %s: %s", upc, requiredQuantityStr)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity parameter"})
			return
		}

		available, err := service.CheckStockAvailability(upc, requiredQuantity)
		if err != nil {
			log.Printf("[StoreProductStockCheckGET] Service error for UPC %s, quantity %d: %v", upc, requiredQuantity, err)
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
