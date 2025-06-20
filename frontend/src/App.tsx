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
          <Route path="/create-check" element={<CreateCheck />} />
          <Route path="/checks" element={<Checks />} />
          <Route path="/checks/:receipt_number" element={<CheckDetails />} />
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
