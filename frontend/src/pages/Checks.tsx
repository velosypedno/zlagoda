import React, { useEffect, useState } from "react";
import type { Check } from "../types/check";
import { Link } from "react-router-dom";
import { fetchChecks } from "../api/checks";

const Checks: React.FC = () => {
  const [checks, setChecks] = useState<Check[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    setLoading(true);
    fetchChecks()
      .then((res) => {
        setChecks(res.data);
        setError(null);
      })
      .catch((err) => {
        setError(err?.response?.data?.error || err.message || "Failed to load checks");
      })
      .finally(() => setLoading(false));
  }, []);

  return (
    <div className="max-w-4xl mx-auto p-6 bg-white rounded shadow">
      <div className="flex justify-between items-center mb-4">
        <h1 className="text-2xl font-bold">Checks (Receipts)</h1>
        <Link to="/create-check" className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600">Create Check</Link>
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
              {checks.map((check) => (
                <tr key={check.receipt_number} className="hover:bg-gray-50">
                  <td className="px-3 py-2 border font-mono">{check.receipt_number}</td>
                  <td className="px-3 py-2 border">{check.employee_id}</td>
                  <td className="px-3 py-2 border">{check.card_number || <span className="text-gray-400">—</span>}</td>
                  <td className="px-3 py-2 border">{check.print_date}</td>
                  <td className="px-3 py-2 border">{check.sum_total.toFixed(2)}</td>
                  <td className="px-3 py-2 border">{check.vat.toFixed(2)}</td>
                  <td className="px-3 py-2 border">
                    <Link to={`/checks/${check.receipt_number}`} className="text-blue-600 hover:underline">View</Link>
                  </td>
                </tr>
              ))}
              {checks.length === 0 && (
                <tr>
                  <td colSpan={6} className="text-center text-gray-500 py-4">No checks found.</td>
                </tr>
              )}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
};

export default Checks; 