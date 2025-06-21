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
import Navbar from "./components/Navbar";

function App() {
  return (
    <BrowserRouter>
      <Navbar />
      <div className="pt-16">
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/categories" element={<CategoriesPage />} />
          <Route path="/employees" element={<EmployeesPage />} />
          <Route path="/customer-cards" element={<CustomerCardsPage />} />
          <Route path="/products" element={<ProductsPage />} />
          <Route path="/store-products" element={<StoreProductsPage />} />
          <Route path="/create-receipt" element={<CreateReceipt />} />
          <Route path="/receipts" element={<Receipts />} />
          <Route
            path="/receipts/:receipt_number"
            element={<ReceiptDetails />}
          />
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
