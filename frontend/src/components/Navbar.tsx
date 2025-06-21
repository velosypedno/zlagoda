import { Link } from "react-router-dom";
import { useAuth } from "../contexts/AuthContext";

const Navbar = () => {
  const { isAuthenticated, isManager, isCashier, logout } = useAuth();

  const handleLogout = () => {
    logout();
    window.location.href = "/login";
  };

  return (
    <nav className="bg-white shadow-md fixed top-0 left-0 w-full z-10">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between h-16">
          <div className="flex-shrink-0">
            <Link to="/" className="text-2xl font-bold text-blue-600">
              Zlagoda
            </Link>
          </div>

          <div className="hidden md:flex space-x-4">
            {isAuthenticated ? (
              <>
                <Link
                  to="/"
                  className="text-gray-700 hover:text-blue-500 transition"
                >
                  Home
                </Link>
                <Link
                  to="/categories"
                  className="text-gray-700 hover:text-blue-500 transition"
                >
                  Categories
                </Link>
                <Link
                  to="/products"
                  className="text-gray-700 hover:text-blue-500 transition"
                >
                  Products
                </Link>
                <Link
                  to="/store-products"
                  className="text-gray-700 hover:text-blue-500 transition"
                >
                  Store Products
                </Link>
                <Link
                  to="/customer-cards"
                  className="text-gray-700 hover:text-blue-500 transition"
                >
                  Customer Cards
                </Link>
                <Link
                  to="/receipts"
                  className="text-gray-700 hover:text-blue-500 transition"
                >
                  Receipts
                </Link>

                {/* Manager-only navigation */}
                {isManager && (
                  <Link
                    to="/employees"
                    className="text-gray-700 hover:text-blue-500 transition"
                  >
                    Employees
                  </Link>
                )}

                {/* Cashier-only navigation */}
                {isCashier && (
                  <Link
                    to="/create-receipt"
                    className="text-gray-700 hover:text-blue-500 transition"
                  >
                    Create Receipt
                  </Link>
                )}

                <Link
                  to="/account"
                  className="text-gray-700 hover:text-blue-500 transition"
                >
                  Account
                </Link>
                <button
                  onClick={handleLogout}
                  className="text-gray-700 hover:text-red-500 transition"
                >
                  Logout
                </button>
              </>
            ) : (
              <>
                <Link
                  to="/login"
                  className="text-gray-700 hover:text-blue-500 transition"
                >
                  Login
                </Link>
                <Link
                  to="/register"
                  className="text-gray-700 hover:text-blue-500 transition"
                >
                  Register
                </Link>
              </>
            )}
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
