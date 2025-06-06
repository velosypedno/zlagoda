import { useEffect, useState } from "react";
import { fetchCategories, deleteCategory, updateCategory, createCategory } from "../api/categories";
import { type Category } from "../types/category";
import CategoryCard from "../components/CategoryCard";

const CategoriesPage = () => {
  const [categories, setCategories] = useState<Category[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [newName, setNewName] = useState("");

  const loadCategories = async () => {
    try {
      const res = await fetchCategories();
      setCategories(res.data);
    } catch (err) {
      setError("Failed to load categories");
    }
  };

  useEffect(() => {
    loadCategories();
  }, []);

  const handleDelete = async (id: number) => {
    try {
      await deleteCategory(id);
      setCategories((prev) => prev.filter((cat) => cat.id !== id));
    } catch {
      setError("Failed to delete category");
    }
  };

  const handleUpdate = async (id: number, name: string) => {
    try {
      await updateCategory(id, name);
      setCategories((prev) =>
        prev.map((cat) => (cat.id === id ? { ...cat, name } : cat))
      );
    } catch {
      setError("Failed to update category");
    }
  };

  const handleCreate = async () => {
    if (!newName.trim()) return;

    try {
      const res = await createCategory(newName.trim());
      setCategories((prev) => [...prev, res.data]);
      setNewName("");
      setError(null);
      await loadCategories()
    } catch {
      setError("Failed to create category");
    }
  };

  return (
    <div className="p-6 max-w-5xl mx-auto">
      <h1 className="text-3xl font-bold mb-6">Categories</h1>

      {error && <div className="mb-4 text-red-500">{error}</div>}

      <div className="flex mb-6 gap-2">
        <input
          type="text"
          placeholder="New category name"
          value={newName}
          onChange={(e) => setNewName(e.target.value)}
          className="flex-grow border rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
        />
        <button
          onClick={handleCreate}
          className="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700 disabled:opacity-50"
          disabled={!newName.trim()}
        >
          Add
        </button>
      </div>


      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
        {categories.map((category) => (
          <CategoryCard
            key={category.id}
            category={category}
            onDelete={() => handleDelete(category.id)}
            onUpdate={(name) => handleUpdate(category.id, name)}
          />
        ))}
      </div>
    </div>
  );
};

export default CategoriesPage;
