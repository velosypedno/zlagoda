import React, { useEffect, useState } from "react";
import type { ReceiptRetrieve } from "../types/receipt";
import { Link } from "react-router-dom";
import { fetchReceipts } from "../api/receipts";

const Receipts: React.FC = () => {
  const [receipts, setReceipts] = useState<ReceiptRetrieve[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    setLoading(true);
    fetchReceipts()
      .then((res) => {
        setReceipts(res.data);
        setError(null);
      })
      .catch((err) => {
        setError(err?.response?.data?.error || err.message || "Failed to load receipts");
      })
      .finally(() => setLoading(false));
  }, []);

  return (
    <div className="max-w-4xl mx-auto p-6 bg-white rounded shadow">
      <div className="flex justify-between items-center mb-4">
        <h1 className="text-2xl font-bold">Receipts</h1>
        <Link to="/create-receipt" className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600">Create Receipt</Link>
      </div>
      {loading && <div>Loading...</div>}
      {error && <div className="bg-red-100 text-red-700 p-2 mb-2 rounded">{error}</div>}
      {!loading && !error && (
        <div className="overflow-x-auto">
          <table className="min-w-full border">
            <thead>
              <tr className="bg-gray-100">
                <th className="px-3 py-2 border">Receipt #</th>
                <th className="px-3 py-2 border">Employee ID</th>
                <th className="px-3 py-2 border">Card #</th>
                <th className="px-3 py-2 border">Print Date</th>
                <th className="px-3 py-2 border">Total</th>
                <th className="px-3 py-2 border">VAT</th>
                <th className="px-3 py-2 border">Actions</th>
              </tr>
            </thead>
            <tbody>
              {receipts.map((receipt) => (
                <tr key={receipt.receipt_number} className="hover:bg-gray-50">
                  <td className="px-3 py-2 border font-mono">{receipt.receipt_number}</td>
                  <td className="px-3 py-2 border">{receipt.employee_id}</td>
                  <td className="px-3 py-2 border">{receipt.card_number || <span className="text-gray-400">—</span>}</td>
                  <td className="px-3 py-2 border">{receipt.print_date}</td>
                  <td className="px-3 py-2 border">{receipt.sum_total.toFixed(2)}</td>
                  <td className="px-3 py-2 border">{receipt.vat.toFixed(2)}</td>
                  <td className="px-3 py-2 border">
                    <Link to={`/receipts/${receipt.receipt_number}`} className="text-blue-600 hover:underline">View</Link>
                  </td>
                </tr>
              ))}
              {receipts.length === 0 && (
                <tr>
                  <td colSpan={7} className="text-center text-gray-500 py-4">No receipts found.</td>
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
