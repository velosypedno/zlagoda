import { useState, useEffect } from "react";
import type {
  StoreProductWithDetails,
  StoreProductCreate,
  StoreProductUpdate,
} from "../types/store_product";
import type { Product } from "../types/product";
import type { Category } from "../types/category";
import {
  fetchStoreProductsWithDetails,
  fetchStoreProductsByCategory,
  fetchStoreProductsByName,
  createStoreProduct,
  updateStoreProduct,
  deleteStoreProduct,
} from "../api/store_products";
import { fetchProducts } from "../api/products";
import { fetchCategories } from "../api/categories";
import StoreProductCard from "../components/StoreProductCard";
import ExportPdfButton from "../components/ExportPdfButton";
import { useAuth } from "../contexts/AuthContext";

const StoreProducts = () => {
  const { isManager } = useAuth();
  const [storeProducts, setStoreProducts] = useState<StoreProductWithDetails[]>(
    [],
  );
  const [products, setProducts] = useState<Product[]>([]);
  const [categories, setCategories] = useState<Category[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [showForm, setShowForm] = useState(false);
  const [editingStoreProduct, setEditingStoreProduct] =
    useState<StoreProductWithDetails | null>(null);
  const [formData, setFormData] = useState<StoreProductCreate>({
    upc: "",
    upc_prom: "",
    product_id: 0,
    selling_price: 0,
    products_number: 0,
    promotional_product: false,
  });

  // Filtering and search state
  const [selectedCategory, setSelectedCategory] = useState<number>(0);
  const [searchTerm, setSearchTerm] = useState<string>("");
  const [upcSearch, setUpcSearch] = useState<string>("");
  const [showPromotionalOnly, setShowPromotionalOnly] =
    useState<boolean>(false);
  const [sortBy, setSortBy] = useState<
    "product_name" | "upc" | "selling_price" | "products_number"
  >("product_name");
  const [sortOrder, setSortOrder] = useState<"asc" | "desc">("asc");

  useEffect(() => {
    loadData();
  }, []);

  useEffect(() => {
    loadFilteredStoreProducts();
  }, [selectedCategory, searchTerm, showPromotionalOnly]);

  // Debug form data changes
  useEffect(() => {
    console.log("FormData changed:", formData);
  }, [formData]);

  // Debug products array changes
  useEffect(() => {
    console.log("Products array changed:", products.length, "products");
    if (products.length > 0) {
      console.log("First product structure:", products[0]);
      console.log("Product keys:", Object.keys(products[0]));
      console.log("Product.product_id specifically:", products[0].product_id);
      console.log("All products:", products);
    }
  }, [products]);

  const loadData = async () => {
    try {
      setLoading(true);
      const [storeProductsData, productsData, categoriesData] =
        await Promise.all([
          fetchStoreProductsWithDetails(),
          fetchProducts(),
          fetchCategories(),
        ]);
      setStoreProducts(storeProductsData.data || []);
      setProducts(productsData.data || []);
      setCategories(categoriesData.data || []);
      setError(null);
    } catch (err) {
      setError("Failed to load data");
      console.error("Error loading data:", err);
    } finally {
      setLoading(false);
    }
  };

  const loadFilteredStoreProducts = async () => {
    try {
      setLoading(true);
      let storeProductsData;

      if (searchTerm.trim()) {
        storeProductsData = await fetchStoreProductsByName(searchTerm.trim());
      } else if (selectedCategory > 0) {
        storeProductsData =
          await fetchStoreProductsByCategory(selectedCategory);
      } else {
        storeProductsData = await fetchStoreProductsWithDetails();
      }

      let filteredProducts = storeProductsData.data || [];
      if (showPromotionalOnly) {
        filteredProducts = filteredProducts.filter(
          (p) => p.promotional_product,
        );
      }

      setStoreProducts(filteredProducts);
      setError(null);
    } catch (err) {
      setError("Failed to load filtered store products");
      console.error("Error loading filtered store products:", err);
    } finally {
      setLoading(false);
    }
  };

  const handleSort = (
    field: "product_name" | "upc" | "selling_price" | "products_number",
  ) => {
    if (sortBy === field) {
      setSortOrder(sortOrder === "asc" ? "desc" : "asc");
    } else {
      setSortBy(field);
      setSortOrder("asc");
    }
  };

  const getSortedStoreProducts = () => {
    return [...storeProducts].sort((a, b) => {
      let aValue, bValue;

      switch (sortBy) {
        case "product_name":
          aValue = a.product_name.toLowerCase();
          bValue = b.product_name.toLowerCase();
          break;
        case "upc":
          aValue = a.upc.toLowerCase();
          bValue = b.upc.toLowerCase();
          break;
        case "selling_price":
          aValue = a.selling_price;
          bValue = b.selling_price;
          break;
        case "products_number":
          aValue = a.products_number;
          bValue = b.products_number;
          break;
        default:
          aValue = a.product_name.toLowerCase();
          bValue = b.product_name.toLowerCase();
      }

      if (sortOrder === "asc") {
        return aValue < bValue ? -1 : aValue > bValue ? 1 : 0;
      } else {
        return aValue > bValue ? -1 : aValue < bValue ? 1 : 0;
      }
    });
  };

  const handleCategoryChange = (categoryId: number) => {
    setSelectedCategory(categoryId);
    setSearchTerm(""); // Clear search when changing category
  };

  const handleSearchChange = (term: string) => {
    setSearchTerm(term);
    setSelectedCategory(0); // Clear category filter when searching
  };

  const clearFilters = () => {
    setSelectedCategory(0);
    setSearchTerm("");
    setShowPromotionalOnly(false);
    setSortBy("product_name");
    setSortOrder("asc");
  };

  const handlePromotionalCheckboxChange = (checked: boolean) => {
    const newFormData = { ...formData, promotional_product: checked };
    if (!checked) {
      newFormData.upc_prom = "";
    }
    setFormData(newFormData);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (
      !formData.upc ||
      formData.product_id === 0 ||
      formData.selling_price <= 0 ||
      formData.products_number < 0
    ) {
      setError("Please fill in all required fields correctly");
      return;
    }

    if (formData.upc.length !== 12) {
      setError("UPC must be exactly 12 characters");
      return;
    }

    try {
      if (editingStoreProduct) {
        const updateData: StoreProductUpdate = {
          upc_prom: formData.upc_prom || undefined,
          product_id: formData.product_id,
          selling_price: formData.selling_price,
          products_number: formData.products_number,
          promotional_product: formData.promotional_product,
        };
        await updateStoreProduct(editingStoreProduct.upc, updateData);
      } else {
        await createStoreProduct(formData);
      }

      setShowForm(false);
      setEditingStoreProduct(null);
      setFormData({
        upc: "",
        upc_prom: "",
        product_id: 0,
        selling_price: 0,
        products_number: 0,
        promotional_product: false,
      });
      await loadData();
      setError(null);
    } catch (err) {
      setError("Failed to save store product");
      console.error("Error saving store product:", err);
    }
  };

  const handleEdit = (storeProduct: StoreProductWithDetails) => {
    setEditingStoreProduct(storeProduct);
    setFormData({
      upc: storeProduct.upc,
      upc_prom: storeProduct.upc_prom || "",
      product_id: storeProduct.product_id,
      selling_price: storeProduct.selling_price,
      products_number: storeProduct.products_number,
      promotional_product: storeProduct.promotional_product,
    });
    setShowForm(true);
  };

  const handleDelete = async (upc: string) => {
    try {
      await deleteStoreProduct(upc);
      await loadData();
      setError(null);
    } catch (err) {
      setError("Failed to delete store product");
      console.error("Error deleting store product:", err);
    }
  };

  const handleCancel = () => {
    setShowForm(false);
    setEditingStoreProduct(null);
    setFormData({
      upc: "",
      upc_prom: "",
      product_id: 0,
      selling_price: 0,
      products_number: 0,
      promotional_product: false,
    });
    setError(null);
  };

  if (loading) {
    return (
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="text-center">Loading...</div>
      </div>
    );
  }

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-3xl font-bold text-gray-900">Store Products</h1>
        <div className="flex gap-2">
          <ExportPdfButton
            entityType="Store Products"
            apiEndpoint="/api/store-products/details"
            title="Store Products Report"
            filename="store-products-export.pdf"
            columns={[
              { key: "upc", label: "UPC", width: "15%" },
              { key: "product_name", label: "Product", width: "25%" },
              { key: "category_name", label: "Category", width: "20%" },
              { key: "selling_price", label: "Price", width: "10%" },
              { key: "products_number", label: "Stock", width: "10%" },
              { key: "promotional_product", label: "Promo", width: "10%" },
            ]}
          />
          {isManager && (
            <button
              onClick={() => setShowForm(true)}
              className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition"
            >
              Add Store Product
            </button>
          )}
        </div>
      </div>

      {error && (
        <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
          {error}
        </div>
      )}

      {/* Filtering, Search, and Sorting Controls */}
      <div className="bg-white rounded-lg shadow-md p-6 mb-8">
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-4 mb-4">
          {/* Category Filter */}
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Filter by Category
            </label>
            <select
              value={selectedCategory}
              onChange={(e) => handleCategoryChange(parseInt(e.target.value))}
              className="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
              <option value={0}>All Categories</option>
              {categories.map((category) => (
                <option key={category.id} value={category.id}>
                  {category.name}
                </option>
              ))}
            </select>
          </div>

          {/* Search by Product Name */}
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Search by Product Name
            </label>
            <input
              type="text"
              value={searchTerm}
              onChange={(e) => handleSearchChange(e.target.value)}
              placeholder="Enter product name..."
              className="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          {/* Search by UPC */}
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Search by UPC
            </label>
            <input
              type="text"
              value={upcSearch}
              onChange={(e) => setUpcSearch(e.target.value)}
              placeholder="Enter UPC..."
              className="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          {/* Sort By */}
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Sort By
            </label>
            <select
              value={sortBy}
              onChange={(e) =>
                handleSort(
                  e.target.value as
                    | "product_name"
                    | "upc"
                    | "selling_price"
                    | "products_number",
                )
              }
              className="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
              <option value="product_name">Product Name</option>
              <option value="upc">UPC</option>
              <option value="selling_price">Price</option>
              <option value="products_number">Quantity</option>
            </select>
          </div>

          {/* Sort Order */}
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Sort Order
            </label>
            <button
              onClick={() => setSortOrder(sortOrder === "asc" ? "desc" : "asc")}
              className="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500 bg-white hover:bg-gray-50"
            >
              {sortOrder === "asc" ? "↑ Ascending" : "↓ Descending"}
            </button>
          </div>

          {/* Promotional Filter */}
          <div className="flex items-end justify-center h-full">
            <label className="flex items-center cursor-pointer">
              <input
                type="checkbox"
                checked={showPromotionalOnly}
                onChange={(e) => setShowPromotionalOnly(e.target.checked)}
                className="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <span className="ml-2 text-sm font-medium text-gray-700">
                Promotional Only
              </span>
            </label>
          </div>
        </div>

        {/* Clear Filters Button */}
        <div className="flex justify-between items-center">
          <button
            onClick={clearFilters}
            className="text-gray-600 hover:text-gray-800 text-sm underline"
          >
            Clear all filters
          </button>
          <div className="text-sm text-gray-500">
            Showing {getSortedStoreProducts().length} store products
          </div>
        </div>
      </div>

      {showForm && isManager && (
        <div className="bg-white rounded-lg shadow-md p-6 mb-8">
          <h2 className="text-xl font-semibold mb-4">
            {editingStoreProduct
              ? "Edit Store Product"
              : "Add New Store Product"}
          </h2>
          <form onSubmit={handleSubmit} className="space-y-4">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  UPC * (12 characters)
                </label>
                <input
                  type="text"
                  value={formData.upc}
                  onChange={(e) =>
                    setFormData({ ...formData, upc: e.target.value })
                  }
                  className="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                  maxLength={12}
                  required
                  disabled={!!editingStoreProduct}
                />
              </div>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  Product *
                </label>
                <select
                  value={formData.product_id}
                  onChange={(e) => {
                    console.log("Product selection changed:", e.target.value);
                    console.log("Products array at selection:", products);
                    const selectedProduct = products.find(
                      (p) => p.product_id === parseInt(e.target.value),
                    );
                    console.log("Selected product object:", selectedProduct);
                    setFormData({
                      ...formData,
                      product_id: parseInt(e.target.value),
                    });
                  }}
                  className="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                  required
                >
                  <option value={0}>Select a product</option>
                  {products.map((product) => (
                    <option
                      key={`product-${product.product_id}`}
                      value={product.product_id}
                    >
                      {product.name} (ID: {product.product_id})
                    </option>
                  ))}
                </select>
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  Selling Price *
                </label>
                <input
                  type="number"
                  step="0.01"
                  min="0"
                  value={formData.selling_price}
                  onChange={(e) =>
                    setFormData({
                      ...formData,
                      selling_price: parseFloat(e.target.value) || 0,
                    })
                  }
                  className="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                  required
                />
              </div>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  Quantity in Stock *
                </label>
                <input
                  type="number"
                  min="0"
                  value={formData.products_number}
                  onChange={(e) =>
                    setFormData({
                      ...formData,
                      products_number: parseInt(e.target.value) || 0,
                    })
                  }
                  className="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                  required
                />
              </div>

              <div className="flex items-center">
                <label className="flex items-center">
                  <input
                    type="checkbox"
                    checked={formData.promotional_product}
                    onChange={(e) =>
                      handlePromotionalCheckboxChange(e.target.checked)
                    }
                    className="mr-2 h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                  />
                  <span className="text-sm font-medium text-gray-700">
                    Promotional Product
                  </span>
                </label>
              </div>
            </div>

            <div className="flex space-x-4">
              <button
                type="submit"
                className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition"
              >
                {editingStoreProduct ? "Update" : "Create"}
              </button>
              <button
                type="button"
                onClick={handleCancel}
                className="bg-gray-500 text-white px-4 py-2 rounded hover:bg-gray-600 transition"
              >
                Cancel
              </button>
            </div>
          </form>
        </div>
      )}

      {getSortedStoreProducts().length === 0 ? (
        <div className="text-center text-gray-500 py-8">
          No store products found.
        </div>
      ) : (
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
          {getSortedStoreProducts()
            .filter((product) => {
              const upcQ = upcSearch.trim().toLowerCase();
              return !upcQ || product.upc.toLowerCase().includes(upcQ);
            })
            .map((storeProduct) => (
              <StoreProductCard
                key={storeProduct.upc}
                storeProduct={storeProduct}
                onEdit={handleEdit}
                onDelete={handleDelete}
                onDeliveryUpdate={loadData}
              />
            ))}
        </div>
      )}
    </div>
  );
};

export default StoreProducts;
