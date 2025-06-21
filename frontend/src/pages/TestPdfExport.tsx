import React from "react";
import ExportPdfButton from "../components/ExportPdfButton";

const TestPdfExport: React.FC = () => {
  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-gray-900 mb-2">
          PDF Export Test Page
        </h1>
        <p className="text-gray-600">
          Test various PDF export functionalities with real API endpoints
        </p>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
        {/* Categories Export Test */}
        <div className="bg-white rounded-lg shadow-md p-6">
          <h2 className="text-xl font-semibold text-gray-900 mb-4">
            Categories Export
          </h2>
          <p className="text-sm text-gray-600 mb-4">
            Export all categories with ID and name columns.
          </p>
          <ExportPdfButton
            entityType="Categories"
            apiEndpoint="/api/categories"
            title="Categories Report"
            filename="test-categories-export.pdf"
            columns={[
              { key: "id", label: "Category ID", width: "25%" },
              { key: "name", label: "Category Name", width: "75%" },
            ]}
            className="w-full bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition"
          />
        </div>

        {/* Products Export Test */}
        <div className="bg-white rounded-lg shadow-md p-6">
          <h2 className="text-xl font-semibold text-gray-900 mb-4">
            Products Export
          </h2>
          <p className="text-sm text-gray-600 mb-4">
            Export all products with detailed columns.
          </p>
          <ExportPdfButton
            entityType="Products"
            apiEndpoint="/api/products"
            title="Products Inventory Report"
            filename="test-products-export.pdf"
            columns={[
              { key: "product_id", label: "ID", width: "10%" },
              { key: "name", label: "Product Name", width: "30%" },
              { key: "characteristics", label: "Characteristics", width: "40%" },
              { key: "category_id", label: "Category", width: "20%" },
            ]}
            className="w-full bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600 transition"
          />
        </div>

        {/* Store Products Export Test */}
        <div className="bg-white rounded-lg shadow-md p-6">
          <h2 className="text-xl font-semibold text-gray-900 mb-4">
            Store Products Export
          </h2>
          <p className="text-sm text-gray-600 mb-4">
            Export store products with detailed information including prices and
            stock.
          </p>
          <ExportPdfButton
            entityType="Store Products"
            apiEndpoint="/api/store-products/details"
            title="Store Products Detailed Report"
            filename="test-store-products-export.pdf"
            columns={[
              { key: "upc", label: "UPC", width: "15%" },
              { key: "product_name", label: "Product", width: "25%" },
              { key: "category_name", label: "Category", width: "20%" },
              { key: "selling_price", label: "Price ($)", width: "12%" },
              { key: "products_number", label: "Stock", width: "13%" },
              { key: "promotional_product", label: "Promo", width: "15%" },
            ]}
            className="w-full bg-purple-500 text-white px-4 py-2 rounded hover:bg-purple-600 transition"
          />
        </div>

        {/* Employees Export Test */}
        <div className="bg-white rounded-lg shadow-md p-6">
          <h2 className="text-xl font-semibold text-gray-900 mb-4">
            Employees Export
          </h2>
          <p className="text-sm text-gray-600 mb-4">
            Export employee directory with personal and contact information.
          </p>
          <ExportPdfButton
            entityType="Employees"
            apiEndpoint="/api/employees"
            title="Employee Directory"
            filename="test-employees-export.pdf"
            columns={[
              { key: "employee_id", label: "ID", width: "12%" },
              { key: "empl_surname", label: "Surname", width: "18%" },
              { key: "empl_name", label: "Name", width: "15%" },
              { key: "empl_role", label: "Role", width: "15%" },
              { key: "phone_number", label: "Phone", width: "20%" },
              { key: "city", label: "City", width: "20%" },
            ]}
            className="w-full bg-indigo-500 text-white px-4 py-2 rounded hover:bg-indigo-600 transition"
          />
        </div>

        {/* Customer Cards Export Test */}
        <div className="bg-white rounded-lg shadow-md p-6">
          <h2 className="text-xl font-semibold text-gray-900 mb-4">
            Customer Cards Export
          </h2>
          <p className="text-sm text-gray-600 mb-4">
            Export customer cards with discount information.
          </p>
          <ExportPdfButton
            entityType="Customer Cards"
            apiEndpoint="/api/customer-cards"
            title="Customer Cards Report"
            filename="test-customer-cards-export.pdf"
            columns={[
              { key: "card_number", label: "Card #", width: "18%" },
              { key: "cust_surname", label: "Surname", width: "20%" },
              { key: "cust_name", label: "Name", width: "18%" },
              { key: "phone_number", label: "Phone", width: "20%" },
              { key: "city", label: "City", width: "14%" },
              { key: "percent", label: "Discount", width: "10%" },
            ]}
            className="w-full bg-pink-500 text-white px-4 py-2 rounded hover:bg-pink-600 transition"
          />
        </div>

        {/* Receipts Export Test */}
        <div className="bg-white rounded-lg shadow-md p-6">
          <h2 className="text-xl font-semibold text-gray-900 mb-4">
            Receipts Export
          </h2>
          <p className="text-sm text-gray-600 mb-4">
            Export all receipts with transaction details.
          </p>
          <ExportPdfButton
            entityType="Receipts"
            apiEndpoint="/api/receipts"
            title="Receipts Transaction Report"
            filename="test-receipts-export.pdf"
            columns={[
              { key: "receipt_number", label: "Receipt #", width: "18%" },
              { key: "employee_id", label: "Cashier", width: "15%" },
              { key: "card_number", label: "Card #", width: "15%" },
              { key: "print_date", label: "Date", width: "15%" },
              { key: "sum_total", label: "Total", width: "17%" },
              { key: "vat", label: "VAT", width: "20%" },
            ]}
            className="w-full bg-orange-500 text-white px-4 py-2 rounded hover:bg-orange-600 transition"
          />
        </div>
      </div>

      {/* Custom Styled Export Test */}
      <div className="mt-8 bg-gradient-to-r from-blue-500 to-purple-600 rounded-lg p-6 text-white">
        <h2 className="text-xl font-semibold mb-4">Custom Styled Export</h2>
        <p className="text-blue-100 mb-4">
          Test PDF export with custom styling and branding.
        </p>
        <ExportPdfButton
          entityType="Categories"
          apiEndpoint="/api/categories"
          title="🌟 Custom Styled Categories Report 🌟"
          filename="custom-styled-categories.pdf"
          columns={[
            { key: "id", label: "🆔 Category ID", width: "30%" },
            { key: "name", label: "📂 Category Name", width: "70%" },
          ]}
          customStyles={{
            title: {
              fontSize: 28,
              color: "#3b82f6",
              bold: true,
            },
            header: {
              fontSize: 20,
              color: "#1e40af",
            },
            tableHeader: {
              fontSize: 13,
              bold: true,
              color: "#ffffff",
              fillColor: "#3b82f6",
            },
          }}
          className="bg-white text-blue-600 px-6 py-3 rounded-lg hover:bg-blue-50 transition font-semibold"
        >
          🎨 Export with Custom Styling
        </ExportPdfButton>
      </div>

      {/* Instructions */}
      <div className="mt-8 bg-gray-50 rounded-lg p-6">
        <h2 className="text-lg font-semibold text-gray-900 mb-3">
          Testing Instructions
        </h2>
        <div className="text-sm text-gray-700 space-y-2">
          <p>
            <strong>1.</strong> Make sure your backend API is running on{" "}
            <code className="bg-gray-200 px-1 rounded">localhost:8080</code>
          </p>
          <p>
            <strong>2.</strong> Ensure you're logged in with valid credentials
          </p>
          <p>
            <strong>3.</strong> Click any export button to test PDF generation
          </p>
          <p>
            <strong>4.</strong> The PDF should automatically download with the
            current data from your API
          </p>
          <p>
            <strong>5.</strong> Check the browser console for any errors during
            export
          </p>
        </div>
      </div>

      {/* Debug Information */}
      <div className="mt-6 bg-blue-50 border border-blue-200 rounded-lg p-4">
        <h3 className="text-sm font-semibold text-blue-900 mb-2">
          Debug Information
        </h3>
        <div className="text-xs text-blue-800 space-y-1">
          <p>
            <strong>API Base URL:</strong>{" "}
            {import.meta.env.VITE_FRONTEND_BASE_API_URL || "http://localhost:8080"}
          </p>
          <p>
            <strong>Auth Token:</strong>{" "}
            {localStorage.getItem("token") ? "✅ Present" : "❌ Missing"}
          </p>
          <p>
            <strong>Current Time:</strong> {new Date().toLocaleString()}
          </p>
        </div>
      </div>
    </div>
  );
};

export default TestPdfExport;
