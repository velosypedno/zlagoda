import { Link } from "react-router-dom";

const Navbar = () => {
  return (
    <nav className="bg-white shadow-md fixed top-0 left-0 w-full z-10">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between h-16">
          <div className="flex-shrink-0">
            <Link to="/" className="text-2xl font-bold text-blue-600">Zlagoda</Link>
          </div>

          <div className="hidden md:flex space-x-4">
            <Link to="/" className="text-gray-700 hover:text-blue-500 transition">Home</Link>
            <Link to="/categories" className="text-gray-700 hover:text-blue-500 transition">Categories</Link>
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
