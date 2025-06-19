import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from "./pages/Home";
import CategoriesPage from "./pages/Categories";
import EmployeesPage from "./pages/Employees";
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
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
