import { useEffect, useState } from "react";

interface Category {
  id: number;
  name: string;
}

const Categories = () => {
  const [categories, setCategories] = useState<Category[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const baseApiUrl = import.meta.env.VITE_FRONTEND_BASE_API_URL || "";

  useEffect(() => {
    fetch(`http://localhost:8080/api/categories`)
      .then((res) => {
        console.log("Fetch response:", res);
        if (!res.ok) throw new Error(`Failed to fetch categories, status: ${res.status}`);
        return res.json();
      })
      .then((data) => {
        console.log("Parsed data:", data);
        setCategories(data);
        setLoading(false);
      })
      .catch((err) => {
        console.error("Fetch error:", err);
        setError(err.message);
        setLoading(false);
      });
  }, [baseApiUrl]);

  if (loading) return <div className="p-4">Loading...</div>;
  if (error) return <div className="p-4 text-red-500">Error: {error}</div>;

  return (
    <div className="p-6">
      <h1 className="text-3xl font-bold mb-6">Categories</h1>
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
        {categories.map((category) => (
          <div
            key={category.id}
            className="bg-white shadow-md rounded-lg p-6 flex flex-col justify-center items-center hover:shadow-lg transition-shadow duration-300"
          >
            <h2 className="text-xl font-semibold mb-2">{category.name}</h2>
            <p className="text-gray-500 text-sm">ID: {category.id}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Categories;
