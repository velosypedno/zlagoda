import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from "./pages/Home";
import CategoriesPage from "./pages/Categories";
import EmployeesPage from "./pages/Employees";
import CustomerCardsPage from "./pages/CustomerCards";
import ProductsPage from "./pages/Products";
import StoreProductsPage from "./pages/StoreProducts";
import CreateReceipt from "./pages/CreateReceipt";
import Receipts from "./pages/Receipts";
import ReceiptDetails from "./pages/ReceiptDetails";

import Login from "./pages/Login";
import Register from "./pages/Register";
import Account from "./pages/Account";
import Unauthorized from "./pages/Unauthorized";
import Navbar from "./components/Navbar";
import ProtectedRoute from "./components/ProtectedRoute";
import { AuthProvider } from "./contexts/AuthContext";

function App() {
  return (
    <AuthProvider>
      <BrowserRouter>
        <Navbar />
        <div className="pt-16">
          <Routes>
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />
            <Route path="/unauthorized" element={<Unauthorized />} />

            {/* Public routes (require authentication) */}
            <Route
              path="/"
              element={
                <ProtectedRoute>
                  <Home />
                </ProtectedRoute>
              }
            />
            <Route
              path="/account"
              element={
                <ProtectedRoute>
                  <Account />
                </ProtectedRoute>
              }
            />
            <Route
              path="/categories"
              element={
                <ProtectedRoute>
                  <CategoriesPage />
                </ProtectedRoute>
              }
            />
            <Route
              path="/customer-cards"
              element={
                <ProtectedRoute>
                  <CustomerCardsPage />
                </ProtectedRoute>
              }
            />
            <Route
              path="/products"
              element={
                <ProtectedRoute>
                  <ProductsPage />
                </ProtectedRoute>
              }
            />
            <Route
              path="/store-products"
              element={
                <ProtectedRoute>
                  <StoreProductsPage />
                </ProtectedRoute>
              }
            />
            <Route
              path="/receipts"
              element={
                <ProtectedRoute>
                  <Receipts />
                </ProtectedRoute>
              }
            />
            <Route
              path="/receipts/:receipt_number"
              element={
                <ProtectedRoute>
                  <ReceiptDetails />
                </ProtectedRoute>
              }
            />

            {/* Manager-only routes */}
            <Route
              path="/employees"
              element={
                <ProtectedRoute requireManager>
                  <EmployeesPage />
                </ProtectedRoute>
              }
            />

            {/* Cashier-only routes */}
            <Route
              path="/create-receipt"
              element={
                <ProtectedRoute requireCashier>
                  <CreateReceipt />
                </ProtectedRoute>
              }
            />
          </Routes>
        </div>
      </BrowserRouter>
    </AuthProvider>
  );
}

export default App;
