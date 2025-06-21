import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from "./pages/Home";
import CategoriesPage from "./pages/Categories";
import EmployeesPage from "./pages/Employees";
import CustomerCardsPage from "./pages/CustomerCards";
import ProductsPage from "./pages/Products";
import StoreProductsPage from "./pages/StoreProducts";
import CreateCheck from "./pages/CreateCheck";
import Checks from "./pages/Checks";
import CheckDetails from "./pages/CheckDetails";
import Login from "./pages/Login";
import Register from "./pages/Register";
import Account from "./pages/Account";
import Navbar from "./components/Navbar";
import ProtectedRoute from "./components/ProtectedRoute";

function App() {
  return (
    <BrowserRouter>
      <Navbar />
      <div className="pt-16">
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route path="/" element={<ProtectedRoute><Home /></ProtectedRoute>} />
          <Route path="/account" element={<ProtectedRoute><Account /></ProtectedRoute>} />
          <Route path="/categories" element={<ProtectedRoute><CategoriesPage /></ProtectedRoute>} />
          <Route path="/employees" element={<ProtectedRoute><EmployeesPage /></ProtectedRoute>} />
          <Route path="/customer-cards" element={<ProtectedRoute><CustomerCardsPage /></ProtectedRoute>} />
          <Route path="/products" element={<ProtectedRoute><ProductsPage /></ProtectedRoute>} />
          <Route path="/store-products" element={<ProtectedRoute><StoreProductsPage /></ProtectedRoute>} />
          <Route path="/create-check" element={<ProtectedRoute><CreateCheck /></ProtectedRoute>} />
          <Route path="/checks" element={<ProtectedRoute><Checks /></ProtectedRoute>} />
          <Route path="/checks/:receipt_number" element={<ProtectedRoute><CheckDetails /></ProtectedRoute>} />
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
