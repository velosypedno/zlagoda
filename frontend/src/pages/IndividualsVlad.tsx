import { useState } from 'react';
import { queryVlad1, queryVlad2, type Vlad1Response, type Vlad2Response } from '../api/individuals';
import { fetchCategories } from '../api/categories';
import type { Category } from '../types/category';
import { useEffect } from 'react';

const IndividualsVlad = () => {
  // Vlad1 state
  const [vlad1Results, setVlad1Results] = useState<Vlad1Response[]>([]);
  const [vlad1Loading, setVlad1Loading] = useState(false);
  const [vlad1Error, setVlad1Error] = useState<string | null>(null);
  const [vlad1CategoryId, setVlad1CategoryId] = useState<number>(1);
  const [vlad1Months, setVlad1Months] = useState<number>(3);
  const [categories, setCategories] = useState<Category[]>([]);

  // Vlad2 state
  const [vlad2Results, setVlad2Results] = useState<Vlad2Response[]>([]);
  const [vlad2Loading, setVlad2Loading] = useState(false);
  const [vlad2Error, setVlad2Error] = useState<string | null>(null);

  // Load categories for Vlad1 dropdown
  useEffect(() => {
    const loadCategories = async () => {
      try {
        const response = await fetchCategories();
        setCategories(response.data || []);
      } catch (err) {
        console.error('Failed to load categories:', err);
      }
    };
    loadCategories();
  }, []);

  // Vlad1 query execution
  const executeVlad1 = async () => {
    if (vlad1CategoryId <= 0) {
      setVlad1Error('Please select a valid category');
      return;
    }

    setVlad1Loading(true);
    setVlad1Error(null);

    try {
      const response = await queryVlad1(vlad1CategoryId, vlad1Months);
      setVlad1Results(response.results || []);
    } catch (err: any) {
      setVlad1Error(err?.response?.data?.error || err.message || 'Failed to execute Vlad1 query');
      setVlad1Results([]);
    } finally {
      setVlad1Loading(false);
    }
  };

  // Vlad2 query execution
  const executeVlad2 = async () => {
    setVlad2Loading(true);
    setVlad2Error(null);

    try {
      const response = await queryVlad2();
      setVlad2Results(response.results || []);
    } catch (err: any) {
      setVlad2Error(err?.response?.data?.error || err.message || 'Failed to execute Vlad2 query');
      setVlad2Results([]);
    } finally {
      setVlad2Loading(false);
    }
  };

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-gray-900 mb-2">Vlad's Individual Queries</h1>
        <p className="text-gray-600">Analysis queries for sales and employee performance</p>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
        {/* Vlad1 Query Section */}
        <div className="bg-white rounded-lg shadow-md p-6">
          <div className="border-b border-gray-200 pb-4 mb-6">
            <h2 className="text-xl font-semibold text-gray-900 mb-2">
              Query 1: Most Sold Products in Category
            </h2>
            <p className="text-sm text-gray-600">
              Find the most sold products in a specific category within a time period
            </p>
          </div>

          {/* Vlad1 Parameters */}
          <div className="space-y-4 mb-6">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Category *
              </label>
              <select
                value={vlad1CategoryId}
                onChange={(e) => setVlad1CategoryId(parseInt(e.target.value))}
                className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              >
                <option value={0}>Select a category</option>
                {categories.map((category) => (
                  <option key={category.id} value={category.id}>
                    {category.name}
                  </option>
                ))}
              </select>
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Time Period (months) *
              </label>
              <input
                type="number"
                min="1"
                max="12"
                value={vlad1Months}
                onChange={(e) => setVlad1Months(parseInt(e.target.value) || 1)}
                className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                placeholder="Number of months (1-12)"
              />
            </div>

            <button
              onClick={executeVlad1}
              disabled={vlad1Loading}
              className="w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
            >
              {vlad1Loading ? 'Executing...' : 'Execute Query'}
            </button>
          </div>

          {/* Vlad1 Error */}
          {vlad1Error && (
            <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
              {vlad1Error}
            </div>
          )}

          {/* Vlad1 Results */}
          <div className="bg-gray-50 rounded-lg p-4">
            <h3 className="font-medium text-gray-900 mb-3">Results ({vlad1Results.length})</h3>
            {vlad1Results.length > 0 ? (
              <div className="overflow-x-auto">
                <table className="min-w-full bg-white border border-gray-200 rounded">
                  <thead className="bg-gray-100">
                    <tr>
                      <th className="px-3 py-2 text-left text-xs font-medium text-gray-500 uppercase border-b">Category</th>
                      <th className="px-3 py-2 text-left text-xs font-medium text-gray-500 uppercase border-b">Product</th>
                      <th className="px-3 py-2 text-right text-xs font-medium text-gray-500 uppercase border-b">Sales</th>
                      <th className="px-3 py-2 text-right text-xs font-medium text-gray-500 uppercase border-b">Units Sold</th>
                      <th className="px-3 py-2 text-right text-xs font-medium text-gray-500 uppercase border-b">Revenue</th>
                    </tr>
                  </thead>
                  <tbody>
                    {vlad1Results.map((result, index) => (
                      <tr key={index} className="border-b hover:bg-gray-50">
                        <td className="px-3 py-2 text-sm text-gray-900">{result.category_name}</td>
                        <td className="px-3 py-2 text-sm text-gray-900">{result.product_name}</td>
                        <td className="px-3 py-2 text-sm text-gray-900 text-right">{result.total_sales}</td>
                        <td className="px-3 py-2 text-sm text-gray-900 text-right">{result.total_units_sold}</td>
                        <td className="px-3 py-2 text-sm text-gray-900 text-right">${result.total_revenue.toFixed(2)}</td>
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

        {/* Vlad2 Query Section */}
        <div className="bg-white rounded-lg shadow-md p-6">
          <div className="border-b border-gray-200 pb-4 mb-6">
            <h2 className="text-xl font-semibold text-gray-900 mb-2">
              Query 2: Employees Without Promotional Sales
            </h2>
            <p className="text-sm text-gray-600">
              Find employees who have never sold any promotional products
            </p>
          </div>

          {/* Vlad2 Parameters */}
          <div className="mb-6">
            <p className="text-sm text-gray-600 mb-4">
              This query requires no parameters. Click the button below to execute.
            </p>

            <button
              onClick={executeVlad2}
              disabled={vlad2Loading}
              className="w-full bg-green-600 text-white py-2 px-4 rounded-md hover:bg-green-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
            >
              {vlad2Loading ? 'Executing...' : 'Execute Query'}
            </button>
          </div>

          {/* Vlad2 Error */}
          {vlad2Error && (
            <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
              {vlad2Error}
            </div>
          )}

          {/* Vlad2 Results */}
          <div className="bg-gray-50 rounded-lg p-4">
            <h3 className="font-medium text-gray-900 mb-3">Results ({vlad2Results.length})</h3>
            {vlad2Results.length > 0 ? (
              <div className="overflow-x-auto">
                <table className="min-w-full bg-white border border-gray-200 rounded">
                  <thead className="bg-gray-100">
                    <tr>
                      <th className="px-3 py-2 text-left text-xs font-medium text-gray-500 uppercase border-b">Employee ID</th>
                      <th className="px-3 py-2 text-left text-xs font-medium text-gray-500 uppercase border-b">Surname</th>
                      <th className="px-3 py-2 text-left text-xs font-medium text-gray-500 uppercase border-b">Name</th>
                    </tr>
                  </thead>
                  <tbody>
                    {vlad2Results.map((result, index) => (
                      <tr key={index} className="border-b hover:bg-gray-50">
                        <td className="px-3 py-2 text-sm text-gray-900 font-mono">{result.employee_id}</td>
                        <td className="px-3 py-2 text-sm text-gray-900">{result.surname}</td>
                        <td className="px-3 py-2 text-sm text-gray-900">{result.employee_name}</td>
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

export default IndividualsVlad;
