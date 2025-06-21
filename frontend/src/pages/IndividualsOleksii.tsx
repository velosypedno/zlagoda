import { useState } from 'react';
import { queryOleksii1, queryOleksii2, type Oleksii1Response, type Oleksii2Response } from '../api/individuals';

const IndividualsOleksii = () => {
  // Oleksii1 state
  const [oleksii1Results, setOleksii1Results] = useState<Oleksii1Response[]>([]);
  const [oleksii1Loading, setOleksii1Loading] = useState(false);
  const [oleksii1Error, setOleksii1Error] = useState<string | null>(null);
  const [oleksii1DiscountThreshold, setOleksii1DiscountThreshold] = useState<number>(15);

  // Oleksii2 state
  const [oleksii2Results, setOleksii2Results] = useState<Oleksii2Response[]>([]);
  const [oleksii2Loading, setOleksii2Loading] = useState(false);
  const [oleksii2Error, setOleksii2Error] = useState<string | null>(null);

  // Oleksii1 query execution
  const executeOleksii1 = async () => {
    if (oleksii1DiscountThreshold < 0 || oleksii1DiscountThreshold > 100) {
      setOleksii1Error('Discount threshold must be between 0 and 100');
      return;
    }

    setOleksii1Loading(true);
    setOleksii1Error(null);

    try {
      const response = await queryOleksii1(oleksii1DiscountThreshold);
      setOleksii1Results(response.results || []);
    } catch (err: any) {
      setOleksii1Error(err?.response?.data?.error || err.message || 'Failed to execute Oleksii1 query');
      setOleksii1Results([]);
    } finally {
      setOleksii1Loading(false);
    }
  };

  // Oleksii2 query execution
  const executeOleksii2 = async () => {
    setOleksii2Loading(true);
    setOleksii2Error(null);

    try {
      const response = await queryOleksii2();
      setOleksii2Results(response.results || []);
    } catch (err: any) {
      setOleksii2Error(err?.response?.data?.error || err.message || 'Failed to execute Oleksii2 query');
      setOleksii2Results([]);
    } finally {
      setOleksii2Loading(false);
    }
  };

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-gray-900 mb-2">Oleksii's Individual Queries</h1>
        <p className="text-gray-600">Customer behavior and cashier performance analysis</p>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
        {/* Oleksii1 Query Section */}
        <div className="bg-white rounded-lg shadow-md p-6">
          <div className="border-b border-gray-200 pb-4 mb-6">
            <h2 className="text-xl font-semibold text-gray-900 mb-2">
              Query 1: Cashiers with High Discount Customers
            </h2>
            <p className="text-sm text-gray-600">
              Find cashiers who served customers with discount above a certain threshold
            </p>
          </div>

          {/* Oleksii1 Parameters */}
          <div className="space-y-4 mb-6">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Discount Threshold (%) *
              </label>
              <input
                type="number"
                min="0"
                max="100"
                value={oleksii1DiscountThreshold}
                onChange={(e) => setOleksii1DiscountThreshold(parseInt(e.target.value) || 0)}
                className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                placeholder="Enter discount threshold (0-100)"
              />
              <p className="text-xs text-gray-500 mt-1">
                Find cashiers who served customers with discount above this percentage
              </p>
            </div>

            <button
              onClick={executeOleksii1}
              disabled={oleksii1Loading}
              className="w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
            >
              {oleksii1Loading ? 'Executing...' : 'Execute Query'}
            </button>
          </div>

          {/* Oleksii1 Error */}
          {oleksii1Error && (
            <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
              {oleksii1Error}
            </div>
          )}

          {/* Oleksii1 Results */}
          <div className="bg-gray-50 rounded-lg p-4">
            <h3 className="font-medium text-gray-900 mb-3">Results ({oleksii1Results.length})</h3>
            {oleksii1Results.length > 0 ? (
              <div className="overflow-x-auto">
                <table className="min-w-full bg-white border border-gray-200 rounded">
                  <thead className="bg-gray-100">
                    <tr>
                      <th className="px-2 py-2 text-left text-xs font-medium text-gray-500 uppercase border-b">Employee ID</th>
                      <th className="px-2 py-2 text-left text-xs font-medium text-gray-500 uppercase border-b">Name</th>
                      <th className="px-2 py-2 text-right text-xs font-medium text-gray-500 uppercase border-b">High Disc. Customers</th>
                      <th className="px-2 py-2 text-right text-xs font-medium text-gray-500 uppercase border-b">Total Receipts</th>
                      <th className="px-2 py-2 text-right text-xs font-medium text-gray-500 uppercase border-b">Total Revenue</th>
                      <th className="px-2 py-2 text-right text-xs font-medium text-gray-500 uppercase border-b">Avg Receipt</th>
                      <th className="px-2 py-2 text-right text-xs font-medium text-gray-500 uppercase border-b">Avg Discount</th>
                    </tr>
                  </thead>
                  <tbody>
                    {oleksii1Results.map((result, index) => (
                      <tr key={index} className="border-b hover:bg-gray-50">
                        <td className="px-2 py-2 text-xs text-gray-900 font-mono">{result.employee_id}</td>
                        <td className="px-2 py-2 text-xs text-gray-900">
                          {result.employee_surname} {result.employee_name}
                        </td>
                        <td className="px-2 py-2 text-xs text-gray-900 text-right">{result.high_discount_customers}</td>
                        <td className="px-2 py-2 text-xs text-gray-900 text-right">{result.total_receipts_high_discount}</td>
                        <td className="px-2 py-2 text-xs text-gray-900 text-right">${result.total_revenue_high_discount.toFixed(2)}</td>
                        <td className="px-2 py-2 text-xs text-gray-900 text-right">${result.avg_receipt_amount.toFixed(2)}</td>
                        <td className="px-2 py-2 text-xs text-gray-900 text-right">{result.avg_customer_discount.toFixed(1)}%</td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            ) : (
              <p className="text-gray-500 text-sm">No results found. Click "Execute Query" to run the analysis.</p>
            )}
          </div>
        </div>

        {/* Oleksii2 Query Section */}
        <div className="bg-white rounded-lg shadow-md p-6">
          <div className="border-b border-gray-200 pb-4 mb-6">
            <h2 className="text-xl font-semibold text-gray-900 mb-2">
              Query 2: Customers from All Categories
            </h2>
            <p className="text-sm text-gray-600">
              Find customers who bought products from all categories in the last month
            </p>
          </div>

          {/* Oleksii2 Parameters */}
          <div className="mb-6">
            <p className="text-sm text-gray-600 mb-4">
              This query analyzes purchases from the last month automatically. Click the button below to execute.
            </p>

            <button
              onClick={executeOleksii2}
              disabled={oleksii2Loading}
              className="w-full bg-green-600 text-white py-2 px-4 rounded-md hover:bg-green-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
            >
              {oleksii2Loading ? 'Executing...' : 'Execute Query'}
            </button>
          </div>

          {/* Oleksii2 Error */}
          {oleksii2Error && (
            <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
              {oleksii2Error}
            </div>
          )}

          {/* Oleksii2 Results */}
          <div className="bg-gray-50 rounded-lg p-4">
            <h3 className="font-medium text-gray-900 mb-3">Results ({oleksii2Results.length})</h3>
            {oleksii2Results.length > 0 ? (
              <div className="overflow-x-auto">
                <table className="min-w-full bg-white border border-gray-200 rounded">
                  <thead className="bg-gray-100">
                    <tr>
                      <th className="px-3 py-2 text-left text-xs font-medium text-gray-500 uppercase border-b">Card Number</th>
                      <th className="px-3 py-2 text-left text-xs font-medium text-gray-500 uppercase border-b">Surname</th>
                      <th className="px-3 py-2 text-left text-xs font-medium text-gray-500 uppercase border-b">Name</th>
                      <th className="px-3 py-2 text-left text-xs font-medium text-gray-500 uppercase border-b">Phone Number</th>
                    </tr>
                  </thead>
                  <tbody>
                    {oleksii2Results.map((result, index) => (
                      <tr key={index} className="border-b hover:bg-gray-50">
                        <td className="px-3 py-2 text-sm text-gray-900 font-mono">{result.card_number}</td>
                        <td className="px-3 py-2 text-sm text-gray-900">{result.cust_surname}</td>
                        <td className="px-3 py-2 text-sm text-gray-900">{result.cust_name}</td>
                        <td className="px-3 py-2 text-sm text-gray-900">{result.phone_number}</td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            ) : (
              <p className="text-gray-500 text-sm">No results found. Click "Execute Query" to run the analysis.</p>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default IndividualsOleksii;
