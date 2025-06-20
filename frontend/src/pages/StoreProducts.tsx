import { useState, useEffect } from "react";
import type { StoreProductWithDetails, StoreProductCreate, StoreProductUpdate } from "../types/store_product";
import type { Product } from "../types/product";
import { 
  fetchStoreProductsWithDetails, 
  createStoreProduct, 
  updateStoreProduct, 
  deleteStoreProduct,
  updateProductQuantity 
} from "../api/store_products";
import { fetchProducts } from "../api/products";
import StoreProductCard from "../components/StoreProductCard";

const StoreProducts = () => {
  const [storeProducts, setStoreProducts] = useState<StoreProductWithDetails[]>([]);
  const [products, setProducts] = useState<Product[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [showForm, setShowForm] = useState(false);
  const [editingStoreProduct, setEditingStoreProduct] = useState<StoreProductWithDetails | null>(null);
  const [formData, setFormData] = useState<StoreProductCreate>({
    upc: "",
    upc_prom: "",
    product_id: 0,
    selling_price: 0,
    products_number: 0,
    promotional_product: false,
  });

  useEffect(() => {
    loadData();
  }, []);

  const loadData = async () => {
    try {
      setLoading(true);
      const [storeProductsData, productsData] = await Promise.all([
        fetchStoreProductsWithDetails(),
        fetchProducts(),
      ]);
      setStoreProducts(storeProductsData.data || []);
      setProducts(productsData.data || []);
      setError(null);
    } catch (err) {
      setError("Failed to load data");
      console.error("Error loading data:", err);
    } finally {
      setLoading(false);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!formData.upc || formData.product_id === 0 || formData.selling_price <= 0 || formData.products_number < 0) {
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

  const handleUpdateQuantity = async (upc: string, quantityChange: number) => {
    try {
      await updateProductQuantity(upc, quantityChange);
      await loadData();
      setError(null);
    } catch (err) {
      setError("Failed to update quantity");
      console.error("Error updating quantity:", err);
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
        <button
          onClick={() => setShowForm(true)}
          className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition"
        >
          Add Store Product
        </button>
      </div>

      {error && (
        <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
          {error}
        </div>
      )}

      {showForm && (
        <div className="bg-white rounded-lg shadow-md p-6 mb-8">
          <h2 className="text-xl font-semibold mb-4">
            {editingStoreProduct ? "Edit Store Product" : "Add New Store Product"}
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
                  onChange={(e) => setFormData({ ...formData, upc: e.target.value })}
                  className="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                  maxLength={12}
                  required
                  disabled={!!editingStoreProduct}
                />
              </div>
              
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  Promotional UPC (12 characters)
                </label>
                <input
                  type="text"
                  value={formData.upc_prom}
                  onChange={(e) => setFormData({ ...formData, upc_prom: e.target.value })}
                  className="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                  maxLength={12}
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
                  onChange={(e) => setFormData({ ...formData, product_id: parseInt(e.target.value) })}
                  className="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                  required
                >
                  <option value={0}>Select a product</option>
                  {products.map((product) => (
                    <option key={product.id} value={product.id}>
                      {product.name}
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
                  onChange={(e) => setFormData({ ...formData, selling_price: parseFloat(e.target.value) || 0 })}
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
                  onChange={(e) => setFormData({ ...formData, products_number: parseInt(e.target.value) || 0 })}
                  className="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                  required
                />
              </div>
              
              <div className="flex items-center">
                <label className="flex items-center">
                  <input
                    type="checkbox"
                    checked={formData.promotional_product}
                    onChange={(e) => setFormData({ ...formData, promotional_product: e.target.checked })}
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

      {storeProducts.length === 0 ? (
        <div className="text-center text-gray-500 py-8">
          No store products found. Create your first store product!
        </div>
      ) : (
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
          {storeProducts.map((storeProduct) => (
            <StoreProductCard
              key={storeProduct.upc}
              storeProduct={storeProduct}
              onEdit={handleEdit}
              onDelete={handleDelete}
              onUpdateQuantity={handleUpdateQuantity}
            />
          ))}
        </div>
      )}
    </div>
  );
};

export default StoreProducts; 