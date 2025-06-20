import { Link } from "react-router-dom";

const Home = () => (
  <div className="p-8">
    <h1 className="text-3xl font-bold mb-4">Welcome to Zlagoda Store Management</h1>
    <div className="space-y-2">
      <Link to="/checks" className="text-blue-600 hover:underline">View Checks (Receipts)</Link>
      <Link to="/products" className="block text-blue-600 hover:underline">Products</Link>
      <Link to="/store-products" className="block text-blue-600 hover:underline">Store Products</Link>
      <Link to="/employees" className="block text-blue-600 hover:underline">Employees</Link>
      <Link to="/customer-cards" className="block text-blue-600 hover:underline">Customer Cards</Link>
      <Link to="/categories" className="block text-blue-600 hover:underline">Categories</Link>
      <Link to="/create-check" className="block text-blue-600 hover:underline font-semibold">Create Check</Link>
    </div>
  </div>
);

export default Home;
