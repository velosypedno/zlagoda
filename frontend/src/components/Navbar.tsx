import { Link, useLocation } from "react-router-dom";
import { useAuth } from "../contexts/AuthContext";
import { useState, useEffect, useRef } from "react";

const Navbar = () => {
  const { user, isAuthenticated, isManager, isCashier, logout } = useAuth();
  const location = useLocation();
  const [isQueriesDropdownOpen, setIsQueriesDropdownOpen] = useState(false);
  const [isProfileDropdownOpen, setIsProfileDropdownOpen] = useState(false);
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

  const queriesDropdownRef = useRef<HTMLDivElement>(null);
  const profileDropdownRef = useRef<HTMLDivElement>(null);

  const handleLogout = () => {
    logout();
    window.location.href = "/login";
  };

  const getUserInitials = () => {
    if (!user) return "U";
    const firstName = user.empl_name.charAt(0).toUpperCase();
    const lastName = user.empl_surname.charAt(0).toUpperCase();
    return `${firstName}${lastName}`;
  };

  const isActivePage = (path: string) => {
    return location.pathname === path;
  };

  // Close dropdowns when clicking outside
  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        queriesDropdownRef.current &&
        !queriesDropdownRef.current.contains(event.target as Node)
      ) {
        setIsQueriesDropdownOpen(false);
      }
      if (
        profileDropdownRef.current &&
        !profileDropdownRef.current.contains(event.target as Node)
      ) {
        setIsProfileDropdownOpen(false);
      }
    };

    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, []);

  // Close mobile menu when route changes
  useEffect(() => {
    setIsMobileMenuOpen(false);
  }, [location]);

  const NavLink = ({
    to,
    children,
    icon,
    className = "",
  }: {
    to: string;
    children: React.ReactNode;
    icon?: string;
    className?: string;
  }) => (
    <Link
      to={to}
      className={`flex items-center space-x-2 px-3 py-2 rounded-md text-sm font-medium transition-colors duration-200 ${
        isActivePage(to)
          ? "bg-blue-100 text-blue-700 border-b-2 border-blue-500"
          : "text-gray-700 hover:text-blue-600 hover:bg-gray-50"
      } ${className}`}
    >
      {icon && <span className="text-lg">{icon}</span>}
      <span>{children}</span>
    </Link>
  );

  const DropdownButton = ({
    isOpen,
    onClick,
    children,
    className = "",
    ariaLabel,
  }: {
    isOpen: boolean;
    onClick: () => void;
    children: React.ReactNode;
    className?: string;
    ariaLabel: string;
  }) => (
    <button
      onClick={onClick}
      className={`flex items-center space-x-1 px-3 py-2 rounded-md text-sm font-medium transition-colors duration-200 text-gray-700 hover:text-blue-600 hover:bg-gray-50 ${className}`}
      aria-expanded={isOpen}
      aria-haspopup="true"
      aria-label={ariaLabel}
    >
      {children}
      <svg
        className={`w-4 h-4 transition-transform duration-200 ${isOpen ? "rotate-180" : ""}`}
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
      >
        <path
          strokeLinecap="round"
          strokeLinejoin="round"
          strokeWidth={2}
          d="M19 9l-7 7-7-7"
        />
      </svg>
    </button>
  );

  if (!isAuthenticated) {
    return (
      <nav
        className="bg-white shadow-md fixed top-0 left-0 w-full z-50"
        role="navigation"
      >
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex items-center justify-between h-16">
            <div className="flex-shrink-0">
              <Link
                to="/"
                className="text-2xl font-bold text-blue-600 hover:text-blue-700 transition-colors"
              >
                🏪 Zlagoda
              </Link>
            </div>
            <div className="flex space-x-4">
              <Link
                to="/login"
                className="text-gray-700 hover:text-blue-500 px-3 py-2 rounded-md text-sm font-medium transition-colors"
              >
                Login
              </Link>
              <Link
                to="/register"
                className="bg-blue-600 text-white hover:bg-blue-700 px-4 py-2 rounded-md text-sm font-medium transition-colors"
              >
                Register
              </Link>
            </div>
          </div>
        </div>
      </nav>
    );
  }

  return (
    <nav
      className="bg-white shadow-md fixed top-0 left-0 w-full z-50"
      role="navigation"
    >
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between h-16">
          {/* Logo */}
          <div className="flex-shrink-0">
            <Link
              to="/"
              className="text-2xl font-bold text-blue-600 hover:text-blue-700 transition-colors"
            >
              🏪 Zlagoda
            </Link>
          </div>

          {/* Desktop Navigation */}
          <div className="hidden lg:flex items-center space-x-1">
            {/* Primary Navigation */}
            <div className="flex items-center space-x-1 border-r border-gray-200 pr-4 mr-4">
              <NavLink to="/" icon="🏠">
                Home
              </NavLink>
              <NavLink to="/categories" icon="📂">
                Categories
              </NavLink>
              <NavLink to="/products" icon="📦">
                Products
              </NavLink>
              <NavLink to="/store-products" icon="🏬">
                Store Products
              </NavLink>
              <NavLink to="/customer-cards" icon="💳">
                Customer Cards
              </NavLink>
              <NavLink to="/receipts" icon="🧾">
                Receipts
              </NavLink>

              {isManager && (
                <NavLink to="/employees" icon="👥">
                  Employees
                </NavLink>
              )}
              {isCashier && (
                <NavLink to="/create-receipt" icon="➕">
                  Create Receipt
                </NavLink>
              )}
            </div>

            {/* Secondary Navigation */}
            <div className="flex items-center space-x-1 border-l border-gray-200 pl-4 ml-4">
              {/* Saved Queries Dropdown */}
              <div className="relative" ref={queriesDropdownRef}>
                <DropdownButton
                  isOpen={isQueriesDropdownOpen}
                  onClick={() =>
                    setIsQueriesDropdownOpen(!isQueriesDropdownOpen)
                  }
                  ariaLabel="Saved queries menu"
                >
                  <span className="text-lg">📊</span>
                  <span>Saved Queries</span>
                </DropdownButton>

                {isQueriesDropdownOpen && (
                  <div
                    className="absolute right-0 mt-2 w-48 bg-white rounded-md shadow-lg py-1 ring-1 ring-black ring-opacity-5"
                    role="menu"
                  >
                    <Link
                      to="/individuals/vlad"
                      className="flex items-center space-x-3 px-4 py-2 text-sm text-gray-700 hover:bg-gray-50"
                      role="menuitem"
                    >
                      <div className="w-6 h-6 bg-purple-500 rounded-full flex items-center justify-center text-white text-xs font-medium">
                        V
                      </div>
                      <span>Vlad's Queries</span>
                    </Link>
                    <Link
                      to="/individuals/arthur"
                      className="flex items-center space-x-3 px-4 py-2 text-sm text-gray-700 hover:bg-gray-50"
                      role="menuitem"
                    >
                      <div className="w-6 h-6 bg-green-500 rounded-full flex items-center justify-center text-white text-xs font-medium">
                        A
                      </div>
                      <span>Arthur's Queries</span>
                    </Link>
                    <Link
                      to="/individuals/oleksii"
                      className="flex items-center space-x-3 px-4 py-2 text-sm text-gray-700 hover:bg-gray-50"
                      role="menuitem"
                    >
                      <div className="w-6 h-6 bg-blue-500 rounded-full flex items-center justify-center text-white text-xs font-medium">
                        O
                      </div>
                      <span>Oleksii's Queries</span>
                    </Link>
                  </div>
                )}
              </div>

              {/* Profile Dropdown */}
              <div className="relative" ref={profileDropdownRef}>
                <button
                  onClick={() =>
                    setIsProfileDropdownOpen(!isProfileDropdownOpen)
                  }
                  className="flex items-center space-x-2 px-3 py-2 rounded-md text-sm font-medium text-gray-700 hover:text-blue-600 hover:bg-gray-50 transition-colors"
                  aria-expanded={isProfileDropdownOpen}
                  aria-haspopup="true"
                  aria-label="User profile menu"
                >
                  <div className="w-8 h-8 bg-blue-600 rounded-full flex items-center justify-center text-white text-sm font-medium">
                    {getUserInitials()}
                  </div>
                  <span className="hidden xl:block">{user?.empl_name}</span>
                  <svg
                    className={`w-4 h-4 transition-transform duration-200 ${isProfileDropdownOpen ? "rotate-180" : ""}`}
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M19 9l-7 7-7-7"
                    />
                  </svg>
                </button>

                {isProfileDropdownOpen && (
                  <div
                    className="absolute right-0 mt-2 w-56 bg-white rounded-md shadow-lg py-1 ring-1 ring-black ring-opacity-5"
                    role="menu"
                  >
                    <div className="px-4 py-2 border-b border-gray-100">
                      <div className="text-sm font-medium text-gray-900">
                        {user?.empl_name} {user?.empl_surname}
                      </div>
                      <div className="text-sm text-gray-500">
                        {user?.empl_role}
                      </div>
                    </div>
                    <Link
                      to="/account"
                      className="flex items-center space-x-2 px-4 py-2 text-sm text-gray-700 hover:bg-gray-50"
                      role="menuitem"
                    >
                      <span>👤</span>
                      <span>Profile Settings</span>
                    </Link>
                    <button
                      onClick={handleLogout}
                      className="flex items-center space-x-2 w-full px-4 py-2 text-sm text-red-600 hover:bg-red-50 text-left"
                      role="menuitem"
                    >
                      <span>🚪</span>
                      <span>Logout</span>
                    </button>
                  </div>
                )}
              </div>
            </div>
          </div>

          {/* Mobile menu button */}
          <div className="lg:hidden">
            <button
              onClick={() => setIsMobileMenuOpen(!isMobileMenuOpen)}
              className="inline-flex items-center justify-center p-2 rounded-md text-gray-700 hover:text-blue-600 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-blue-500"
              aria-expanded={isMobileMenuOpen}
              aria-label="Toggle mobile menu"
            >
              <svg
                className="block h-6 w-6"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                {isMobileMenuOpen ? (
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M6 18L18 6M6 6l12 12"
                  />
                ) : (
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M4 6h16M4 12h16M4 18h16"
                  />
                )}
              </svg>
            </button>
          </div>
        </div>
      </div>

      {/* Mobile menu */}
      {isMobileMenuOpen && (
        <div className="lg:hidden bg-white shadow-lg" role="menu">
          <div className="px-2 pt-2 pb-3 space-y-1 border-t border-gray-200">
            {/* Primary Navigation */}
            <div className="space-y-1">
              <NavLink to="/" icon="🏠" className="block">
                Home
              </NavLink>
              <NavLink to="/categories" icon="📂" className="block">
                Categories
              </NavLink>
              <NavLink to="/products" icon="📦" className="block">
                Products
              </NavLink>
              <NavLink to="/store-products" icon="🏬" className="block">
                Store Products
              </NavLink>
              <NavLink to="/customer-cards" icon="💳" className="block">
                Customer Cards
              </NavLink>
              <NavLink to="/receipts" icon="🧾" className="block">
                Receipts
              </NavLink>

              {isManager && (
                <NavLink to="/employees" icon="👥" className="block">
                  Employees
                </NavLink>
              )}
              {isCashier && (
                <NavLink to="/create-receipt" icon="➕" className="block">
                  Create Receipt
                </NavLink>
              )}
            </div>

            {/* Queries */}
            <div className="pt-2 border-t border-gray-200">
              <div className="px-3 py-2 text-xs font-semibold text-gray-500 uppercase tracking-wider">
                Saved Queries
              </div>
              <Link
                to="/individuals/vlad"
                className="flex items-center space-x-3 px-3 py-2 text-sm text-gray-700 hover:bg-gray-50"
              >
                <div className="w-6 h-6 bg-purple-500 rounded-full flex items-center justify-center text-white text-xs font-medium">
                  V
                </div>
                <span>Vlad's Queries</span>
              </Link>
              <Link
                to="/individuals/arthur"
                className="flex items-center space-x-3 px-3 py-2 text-sm text-gray-700 hover:bg-gray-50"
              >
                <div className="w-6 h-6 bg-green-500 rounded-full flex items-center justify-center text-white text-xs font-medium">
                  A
                </div>
                <span>Arthur's Queries</span>
              </Link>
              <Link
                to="/individuals/oleksii"
                className="flex items-center space-x-3 px-3 py-2 text-sm text-gray-700 hover:bg-gray-50"
              >
                <div className="w-6 h-6 bg-blue-500 rounded-full flex items-center justify-center text-white text-xs font-medium">
                  O
                </div>
                <span>Oleksii's Queries</span>
              </Link>
            </div>

            {/* Profile & Logout */}
            <div className="pt-2 border-t border-gray-200">
              <div className="flex items-center space-x-3 px-3 py-2">
                <div className="w-8 h-8 bg-blue-600 rounded-full flex items-center justify-center text-white text-sm font-medium">
                  {getUserInitials()}
                </div>
                <div>
                  <div className="text-sm font-medium text-gray-900">
                    {user?.empl_name} {user?.empl_surname}
                  </div>
                  <div className="text-xs text-gray-500">{user?.empl_role}</div>
                </div>
              </div>
              <Link
                to="/account"
                className="flex items-center space-x-3 px-3 py-2 text-sm text-gray-700 hover:bg-gray-50"
              >
                <span>👤</span>
                <span>Profile Settings</span>
              </Link>
              <button
                onClick={handleLogout}
                className="flex items-center space-x-3 w-full px-3 py-2 text-sm text-red-600 hover:bg-red-50 text-left"
              >
                <span>🚪</span>
                <span>Logout</span>
              </button>
            </div>
          </div>
        </div>
      )}
    </nav>
  );
};

export default Navbar;
