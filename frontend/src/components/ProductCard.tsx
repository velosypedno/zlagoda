import { useState } from "react";
import type { Product } from "../types/product";
import type { Category } from "../types/category";

interface ProductCardProps {
  product: Product;
  categories: Category[];
  onEdit: (product: Product) => void;
  onDelete: (id: number) => void;
}

const ProductCard = ({ product, categories, onEdit, onDelete }: ProductCardProps) => {
  const [isDeleting, setIsDeleting] = useState(false);

  const category = categories.find(cat => cat.id === product.category_id);

  const handleDelete = async () => {
    if (window.confirm("Are you sure you want to delete this product?")) {
      setIsDeleting(true);
      try {
        await onDelete(product.id);
      } finally {
        setIsDeleting(false);
      }
    }
  };

  return (
    <div className="bg-white rounded-lg shadow-md p-6 hover:shadow-lg transition-shadow">
      <div className="flex justify-between items-start mb-4">
        <h3 className="text-xl font-semibold text-gray-800 truncate">
          {product.name}
        </h3>
        <div className="flex space-x-2">
          <button
            onClick={() => onEdit(product)}
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
        </div>
      </div>
      
      <div className="space-y-2">
        <div>
          <span className="text-sm font-medium text-gray-600">Characteristics:</span>
          <p className="text-gray-800 mt-1">{product.characteristics}</p>
        </div>
        
        <div>
          <span className="text-sm font-medium text-gray-600">Category:</span>
          <p className="text-gray-800 mt-1">
            {category ? category.name : `ID: ${product.category_id}`}
          </p>
        </div>
        
        <div>
          <span className="text-sm font-medium text-gray-600">Product ID:</span>
          <p className="text-gray-800 mt-1">{product.id}</p>
        </div>
      </div>
    </div>
  );
};

export default ProductCard; 