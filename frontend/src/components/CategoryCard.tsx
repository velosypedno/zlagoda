import { useState } from "react";
import { type Category } from "../types/category";

interface Props {
  category: Category;
  onDelete: () => void;
  onUpdate: (name: string) => void;
}

const CategoryCard = ({ category, onDelete, onUpdate }: Props) => {
  const [isEditing, setIsEditing] = useState(false);
  const [name, setName] = useState(category.name);

  const handleSave = () => {
    onUpdate(name);
    setIsEditing(false);
  };

  const handleCancel = () => {
    setName(category.name);
    setIsEditing(false);
  };

  return (
    <div className="border shadow p-4 rounded bg-white">
      {isEditing ? (
        <input
          className="border p-1 w-full mb-2"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
      ) : (
        <>
          <h2 className="text-xl font-semibold mb-1">{category.name}</h2>
          <p className="text-sm text-gray-400">ID: {category.id}</p>
        </>
      )}
  
      <div className="flex gap-2 mt-3 items-center min-h-[36px]">
        {isEditing ? (
          <>
            <button
              className="bg-green-500 text-white px-3 py-1 rounded"
              onClick={handleSave}
            >
              Save
            </button>
            <button
              className="bg-gray-300 text-gray-800 px-3 py-1 rounded"
              onClick={handleCancel}
            >
              Cancel
            </button>
            <button
              className="bg-red-500 text-white px-3 py-1 rounded"
              onClick={onDelete}
            >
              Delete
            </button>
          </>
        ) : (
          <>
            <button
              className="bg-blue-500 text-white px-3 py-1 rounded"
              onClick={() => setIsEditing(true)}
            >
              Edit
            </button>
            <button
              className="bg-red-500 text-white px-3 py-1 rounded"
              onClick={onDelete}
            >
              Delete
            </button>
          </>
        )}
      </div>
    </div>
  );
};

export default CategoryCard;
