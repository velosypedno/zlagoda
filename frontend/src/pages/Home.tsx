import { Link } from "react-router-dom";

const Home = () => (
  <div className="p-8 min-h-screen bg-white flex flex-col items-center justify-center">
    <h1
      className="text-4xl font-bold mb-4 text-blue-700"
      style={{ fontFamily: "cursive, sans-serif" }}
    >
      Welcome to <span className="italic">Zlagoda</span> Store Management
    </h1>
    <p className="mb-8 text-base text-gray-700 max-w-xl text-center">
      Manage products, employees, customer cards, and receipts in your store
    </p>
    <div className="space-y-2 w-full max-w-xs">
      <Link
        to="/receipts"
        className="block border border-blue-500 text-blue-700 hover:bg-blue-50 font-medium py-2 px-4 rounded text-center transition"
      >
        View Receipts
      </Link>
      <Link
        to="/products"
        className="block border border-blue-500 text-blue-700 hover:bg-blue-50 font-medium py-2 px-4 rounded text-center transition"
      >
        Products
      </Link>
      <Link
        to="/store-products"
        className="block border border-blue-500 text-blue-700 hover:bg-blue-50 font-medium py-2 px-4 rounded text-center transition"
      >
        Store Products
      </Link>
      <Link
        to="/employees"
        className="block border border-blue-500 text-blue-700 hover:bg-blue-50 font-medium py-2 px-4 rounded text-center transition"
      >
        Employees
      </Link>
      <Link
        to="/customer-cards"
        className="block border border-blue-500 text-blue-700 hover:bg-blue-50 font-medium py-2 px-4 rounded text-center transition"
      >
        Customer Cards
      </Link>
      <Link
        to="/categories"
        className="block border border-blue-500 text-blue-700 hover:bg-blue-50 font-medium py-2 px-4 rounded text-center transition"
      >
        Categories
      </Link>
      <Link
        to="/individuals/vlad"
        className="block border border-purple-500 text-purple-700 hover:bg-purple-50 font-medium py-2 px-4 rounded text-center transition"
      >
        Vlad's Queries
      </Link>
      <Link
        to="/individuals/arthur"
        className="block border border-purple-500 text-purple-700 hover:bg-purple-50 font-medium py-2 px-4 rounded text-center transition"
      >
        Arthur's Queries
      </Link>
      <Link
        to="/individuals/oleksii"
        className="block border border-purple-500 text-purple-700 hover:bg-purple-50 font-medium py-2 px-4 rounded text-center transition"
      >
        Oleksii's Queries
      </Link>
      <Link
        to="/create-receipt"
        className="block bg-blue-500 hover:bg-blue-600 text-white font-medium py-2 px-4 rounded text-center transition"
      >
        Create Receipt
      </Link>
    </div>
  </div>
);

export default Home;
