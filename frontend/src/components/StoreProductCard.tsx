import { useState } from "react";
import type { StoreProductWithDetails } from "../types/store_product";
import { updateProductDelivery } from "../api/store_products";
import { useAuth } from "../contexts/AuthContext";

interface StoreProductCardProps {
  storeProduct: StoreProductWithDetails;
  onEdit: (storeProduct: StoreProductWithDetails) => void;
  onDelete: (upc: string) => void;
  onDeliveryUpdate: () => void;
}

const StoreProductCard = ({ 
  storeProduct, 
  onEdit, 
  onDelete, 
  onDeliveryUpdate
}: StoreProductCardProps) => {
  const { isManager } = useAuth();
  const [isDeleting, setIsDeleting] = useState(false);
  const [isUpdatingDelivery, setIsUpdatingDelivery] = useState(false);
  const [quantityChange, setQuantityChange] = useState(0);
  const [newPrice, setNewPrice] = useState<string>("");
  const [error, setError] = useState<string | null>(null);
  
  const discountedPrice = storeProduct.selling_price * 0.8;

  const handleDelete = async () => {
    if (window.confirm("Are you sure you want to delete this store product?")) {
      setIsDeleting(true);
      try {
        await onDelete(storeProduct.upc);
      } finally {
        setIsDeleting(false);
      }
    }
  };

  const handleDeliveryUpdate = async () => {
    if (quantityChange === 0 && newPrice === "") return;
    setIsUpdatingDelivery(true);
    setError(null);
    try {
      const priceValue = newPrice !== "" ? parseFloat(newPrice) : undefined;
      if (newPrice !== "" && (isNaN(priceValue!) || priceValue! < 0)) {
        setError("Price must be a non-negative number");
        setIsUpdatingDelivery(false);
        return;
      }
      await updateProductDelivery(storeProduct.upc, quantityChange, priceValue);
      setQuantityChange(0);
      setNewPrice("");
      setError(null);
      onDeliveryUpdate();
    } catch (err: any) {
      setError(err?.response?.data?.error || "Failed to update delivery");
    } finally {
      setIsUpdatingDelivery(false);
    }
  };

  return (
    <div className="relative bg-white rounded-lg shadow-md p-6 hover:shadow-lg transition-shadow">
      {storeProduct.promotional_product && (
        <div className="absolute top-0 right-0 -mt-2 -mr-2">
          <span className="inline-flex items-center px-3 py-1 rounded-full text-sm font-semibold bg-orange-500 text-white">
            Sale
          </span>
        </div>
      )}
      <div className="flex justify-between items-start mb-4">
        <div className="flex-1">
          <h3 className="text-xl font-semibold text-gray-800 truncate">
            {storeProduct.product_name}
          </h3>
          <p className="text-sm text-gray-600 mt-1">UPC: {storeProduct.upc}</p>
        </div>
        <div className="flex space-x-2">
          {isManager && (
            <>
              <button
                onClick={() => onEdit(storeProduct)}
                className="px-3 py-1 text-sm bg-blue-500 text-white rounded hover:bg-blue-600 transition"
              >
                Edit
              </button>
              <button
                onClick={handleDelete}
                disabled={isDeleting}
                className="px-3 py-1 text-sm bg-red-500 text-white rounded hover:bg-red-600 transition disabled:opacity-50"
              >
                {isDeleting ? "Deleting..." : "Delete"}
              </button>
            </>
          )}
        </div>
      </div>
      
      <div className="grid grid-cols-2 gap-4 mb-4">
        <div>
          <span className="text-sm font-medium text-gray-600">Category:</span>
          <p className="text-gray-800 mt-1">{storeProduct.category_name}</p>
        </div>
        
        <div>
          <span className="text-sm font-medium text-gray-600">Price:</span>
          {storeProduct.promotional_product ? (
            <div className="mt-1">
              <span className="text-gray-500 line-through mr-2">
                ${storeProduct.selling_price.toFixed(2)}
              </span>
              <span className="text-red-600 font-bold text-lg">
                ${discountedPrice.toFixed(2)}
              </span>
            </div>
          ) : (
            <p className="text-gray-800 mt-1">${storeProduct.selling_price.toFixed(2)}</p>
          )}
        </div>
        
        <div>
          <span className="text-sm font-medium text-gray-600">Stock:</span>
          <p className="text-gray-800 mt-1">{storeProduct.products_number}</p>
        </div>
        
        <div>
          <span className="text-sm font-medium text-gray-600">Type:</span>
          <p className="text-gray-800 mt-1">
            {storeProduct.promotional_product ? (
              <span className="text-orange-600 font-medium">Promotional</span>
            ) : (
              <span className="text-green-600 font-medium">Regular</span>
            )}
          </p>
        </div>
      </div>

      {isManager && (
        <div className="border-t pt-4">
          <div className="flex flex-col gap-2">
            <div className="flex items-center space-x-2">
              <input
                type="number"
                value={quantityChange}
                onChange={(e) => setQuantityChange(parseInt(e.target.value) || 0)}
                placeholder="Quantity change"
                className="flex-1 px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
              <input
                type="number"
                value={newPrice}
                onChange={(e) => setNewPrice(e.target.value)}
                placeholder="New price (optional)"
                min="0"
                step="0.01"
                className="flex-1 px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
              <button
                onClick={handleDeliveryUpdate}
                disabled={isUpdatingDelivery || (quantityChange === 0 && newPrice === "")}
                className="px-4 py-2 text-sm bg-green-500 text-white rounded hover:bg-green-600 transition disabled:opacity-50"
              >
                {isUpdatingDelivery ? "Updating..." : "Update Stock & Price"}
              </button>
            </div>
            {error && <div className="text-red-600 text-sm mt-1">{error}</div>}
          </div>
        </div>
      )}
    </div>
  );
};

export default StoreProductCard; 