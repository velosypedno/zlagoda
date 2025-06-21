import React, { useEffect, useState } from "react";
import type { ReceiptRetrieve } from "../types/receipt";
import type { Employee } from "../types/employee";
import { Link } from "react-router-dom";
import { fetchReceipts } from "../api/receipts";
import { fetchEmployees } from "../api/employees";
import ExportPdfButton from "../components/ExportPdfButton";

const Receipts: React.FC = () => {
  const [receipts, setReceipts] = useState<ReceiptRetrieve[]>([]);
  const [employees, setEmployees] = useState<Employee[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  // Filter states
  const [filterToday, setFilterToday] = useState(false);
  const [filterCashier, setFilterCashier] = useState<string>("");
  const [filterDateFrom, setFilterDateFrom] = useState<string>("");
  const [filterDateTo, setFilterDateTo] = useState<string>("");

  useEffect(() => {
    setLoading(true);
    Promise.all([
      fetchReceipts().then((res) => res.data),
      fetchEmployees().then((res) => res.data),
    ])
      .then(([receiptsData, employeesData]) => {
        setReceipts(receiptsData);
        setEmployees(employeesData);
        setError(null);
      })
      .catch((err) => {
        setError(
          err?.response?.data?.error || err.message || "Failed to load data",
        );
      })
      .finally(() => setLoading(false));
  }, []);

  // Get today's date in YYYY-MM-DD format
  const getTodayDate = () => {
    const today = new Date();
    return today.toISOString().split("T")[0];
  };

  // Filter receipts based on current filter states
  const filteredReceipts = receipts.filter((receipt) => {
    // Filter by today
    if (filterToday && receipt.print_date !== getTodayDate()) {
      return false;
    }

    // Filter by cashier
    if (filterCashier && receipt.employee_id !== filterCashier) {
      return false;
    }

    // Filter by date range
    if (filterDateFrom && receipt.print_date < filterDateFrom) {
      return false;
    }
    if (filterDateTo && receipt.print_date > filterDateTo) {
      return false;
    }

    return true;
  });

  // Get cashier name by employee ID
  const getCashierName = (employeeId: string) => {
    const employee = employees.find((emp) => emp.employee_id === employeeId);
    if (!employee) return employeeId;
    return `${employee.empl_surname} ${employee.empl_name} ${employee.empl_patronymic || ""}`.trim();
  };

  // Clear all filters
  const clearFilters = () => {
    setFilterToday(false);
    setFilterCashier("");
    setFilterDateFrom("");
    setFilterDateTo("");
  };

  return (
    <div className="max-w-6xl mx-auto p-6 bg-white rounded shadow">
      <div className="flex justify-between items-center mb-4">
        <h1 className="text-2xl font-bold">Receipts</h1>
        <div className="flex gap-2">
          <ExportPdfButton
            entityType="Receipts"
            apiEndpoint="/api/receipts"
            title="Receipts Report"
            filename="receipts-export.pdf"
            columns={[
              { key: "receipt_number", label: "Receipt #", width: "15%" },
              { key: "employee_id", label: "Cashier ID", width: "15%" },
              { key: "card_number", label: "Card #", width: "15%" },
              { key: "print_date", label: "Date", width: "15%" },
              { key: "sum_total", label: "Total", width: "12%" },
              { key: "vat", label: "VAT", width: "13%" },
            ]}
          />
          <Link
            to="/create-receipt"
            className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
          >
            Create Receipt
          </Link>
        </div>
      </div>

      {/* Filters Section */}
      <div className="bg-gray-50 p-4 rounded-lg mb-6">
        <h2 className="text-lg font-semibold mb-3">Filters</h2>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
          {/* Today's receipts filter */}
          <div className="flex items-center">
            <input
              type="checkbox"
              id="filterToday"
              checked={filterToday}
              onChange={(e) => setFilterToday(e.target.checked)}
              className="mr-2"
            />
            <label htmlFor="filterToday" className="text-sm font-medium">
              Today's receipts
            </label>
          </div>

          {/* Cashier filter */}
          <div>
            <label className="block text-sm font-medium mb-1">Cashier</label>
            <select
              value={filterCashier}
              onChange={(e) => setFilterCashier(e.target.value)}
              className="w-full border rounded px-3 py-2 text-sm"
            >
              <option value="">All cashiers</option>
              {employees
                .filter((emp) => emp.empl_role === "Cashier")
                .map((emp) => (
                  <option key={emp.employee_id} value={emp.employee_id}>
                    {emp.empl_surname} {emp.empl_name}{" "}
                    {emp.empl_patronymic || ""}
                  </option>
                ))}
            </select>
          </div>

          {/* Date range filters */}
          <div>
            <label className="block text-sm font-medium mb-1">From Date</label>
            <input
              type="date"
              value={filterDateFrom}
              onChange={(e) => setFilterDateFrom(e.target.value)}
              className="w-full border rounded px-3 py-2 text-sm"
            />
          </div>

          <div>
            <label className="block text-sm font-medium mb-1">To Date</label>
            <input
              type="date"
              value={filterDateTo}
              onChange={(e) => setFilterDateTo(e.target.value)}
              className="w-full border rounded px-3 py-2 text-sm"
            />
          </div>
        </div>

        {/* Clear filters button */}
        <div className="mt-3">
          <button
            onClick={clearFilters}
            className="bg-gray-500 text-white px-3 py-1 rounded text-sm hover:bg-gray-600"
          >
            Clear Filters
          </button>
          <span className="ml-3 text-sm text-gray-600">
            Showing {filteredReceipts.length} of {receipts.length} receipts
          </span>
        </div>
      </div>

      {loading && <div>Loading...</div>}
      {error && (
        <div className="bg-red-100 text-red-700 p-2 mb-2 rounded">{error}</div>
      )}
      {!loading && !error && (
        <div className="overflow-x-auto">
          <table className="min-w-full border">
            <thead>
              <tr className="bg-gray-100">
                <th className="px-3 py-2 border">Receipt #</th>
                <th className="px-3 py-2 border">Cashier</th>
                <th className="px-3 py-2 border">Card #</th>
                <th className="px-3 py-2 border">Print Date</th>
                <th className="px-3 py-2 border">Total</th>
                <th className="px-3 py-2 border">VAT</th>
                <th className="px-3 py-2 border">Actions</th>
              </tr>
            </thead>
            <tbody>
              {filteredReceipts.map((receipt) => (
                <tr key={receipt.receipt_number} className="hover:bg-gray-50">
                  <td className="px-3 py-2 border font-mono">
                    {receipt.receipt_number}
                  </td>
                  <td className="px-3 py-2 border">
                    <div className="font-medium">
                      {getCashierName(receipt.employee_id)}
                    </div>
                    <div className="text-xs text-gray-500">
                      ID: {receipt.employee_id}
                    </div>
                  </td>
                  <td className="px-3 py-2 border">
                    {receipt.card_number || (
                      <span className="text-gray-400">—</span>
                    )}
                  </td>
                  <td className="px-3 py-2 border">{receipt.print_date}</td>
                  <td className="px-3 py-2 border">
                    {receipt.sum_total.toFixed(2)}
                  </td>
                  <td className="px-3 py-2 border">{receipt.vat.toFixed(2)}</td>
                  <td className="px-3 py-2 border">
                    <Link
                      to={`/receipts/${receipt.receipt_number}`}
                      className="text-blue-600 hover:underline"
                    >
                      View
                    </Link>
                  </td>
                </tr>
              ))}
              {filteredReceipts.length === 0 && (
                <tr>
                  <td colSpan={7} className="text-center text-gray-500 py-4">
                    {receipts.length === 0
                      ? "No receipts found."
                      : "No receipts match the current filters."}
                  </td>
                </tr>
              )}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
};

export default Receipts;
