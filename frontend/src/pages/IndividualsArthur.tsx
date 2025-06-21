import { useState } from 'react';
import { queryArthur1, queryArthur2, type Arthur1Response, type Arthur2Response } from '../api/individuals';

const IndividualsArthur = () => {
  // Arthur1 state
  const [arthur1Results, setArthur1Results] = useState<Arthur1Response[]>([]);
  const [arthur1Loading, setArthur1Loading] = useState(false);
  const [arthur1Error, setArthur1Error] = useState<string | null>(null);
  const [arthur1StartDate, setArthur1StartDate] = useState<string>('2024-01-01');
  const [arthur1EndDate, setArthur1EndDate] = useState<string>('2024-12-31');

  // Arthur2 state
  const [arthur2Results, setArthur2Results] = useState<Arthur2Response[]>([]);
  const [arthur2Loading, setArthur2Loading] = useState(false);
  const [arthur2Error, setArthur2Error] = useState<string | null>(null);

  // Arthur1 query execution
  const executeArthur1 = async () => {
    if (!arthur1StartDate || !arthur1EndDate) {
      setArthur1Error('Please provide both start and end dates');
      return;
    }

    if (arthur1StartDate > arthur1EndDate) {
      setArthur1Error('Start date must be before end date');
      return;
    }

    setArthur1Loading(true);
    setArthur1Error(null);

    try {
      const response = await queryArthur1(arthur1StartDate, arthur1EndDate);
      setArthur1Results(response.results || []);
    } catch (err: any) {
      setArthur1Error(err?.response?.data?.error || err.message || 'Failed to execute Arthur1 query');
      setArthur1Results([]);
    } finally {
      setArthur1Loading(false);
    }
  };

  // Arthur2 query execution
  const executeArthur2 = async () => {
    setArthur2Loading(true);
    setArthur2Error(null);

    try {
      const response = await queryArthur2();
      setArthur2Results(response.results || []);
    } catch (err: any) {
      setArthur2Error(err?.response?.data?.error || err.message || 'Failed to execute Arthur2 query');
      setArthur2Results([]);
    } finally {
      setArthur2Loading(false);
    }
  };

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-gray-900 mb-2">Arthur's Individual Queries</h1>
        <p className="text-gray-600">Category analysis and product inventory queries</p>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
        {/* Arthur1 Query Section */}
        <div className="bg-white rounded-lg shadow-md p-6">
          <div className="border-b border-gray-200 pb-4 mb-6">
            <h2 className="text-xl font-semibold text-gray-900 mb-2">
              Query 1: Category Sales Statistics
            </h2>
            <p className="text-sm text-gray-600">
              Analyze sales statistics by category within a specific date range
            </p>
          </div>

          {/* Arthur1 Parameters */}
          <div className="space-y-4 mb-6">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Start Date *
              </label>
              <input
                type="date"
                value={arthur1StartDate}
                onChange={(e) => setArthur1StartDate(e.target.value)}
                className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                End Date *
              </label>
              <input
                type="date"
                value={arthur1EndDate}
                onChange={(e) => setArthur1EndDate(e.target.value)}
                className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>

            <button
              onClick={executeArthur1}
              disabled={arthur1Loading}
              className="w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
            >
              {arthur1Loading ? 'Executing...' : 'Execute Query'}
            </button>
          </div>

          {/* Arthur1 Error */}
          {arthur1Error && (
            <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
              {arthur1Error}
            </div>
          )}

          {/* Arthur1 Results */}
          <div className="bg-gray-50 rounded-lg p-4">
            <h3 className="font-medium text-gray-900 mb-3">Results ({arthur1Results.length})</h3>
            {arthur1Results.length > 0 ? (
              <div className="overflow-x-auto">
                <table className="min-w-full bg-white border border-gray-200 rounded">
                  <thead className="bg-gray-100">
                    <tr>
                      <th className="px-3 py-2 text-left text-xs font-medium text-gray-500 uppercase border-b">Category</th>
                      <th className="px-3 py-2 text-right text-xs font-medium text-gray-500 uppercase border-b">Units Sold</th>
                      <th className="px-3 py-2 text-right text-xs font-medium text-gray-500 uppercase border-b">Revenue</th>
                    </tr>
                  </thead>
                  <tbody>
                    {arthur1Results.map((result, index) => (
                      <tr key={index} className="border-b hover:bg-gray-50">
                        <td className="px-3 py-2 text-sm text-gray-900">{result.category_name}</td>
                        <td className="px-3 py-2 text-sm text-gray-900 text-right">{result.units_sold}</td>
                        <td className="px-3 py-2 text-sm text-gray-900 text-right">${result.revenue.toFixed(2)}</td>
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

        {/* Arthur2 Query Section */}
        <div className="bg-white rounded-lg shadow-md p-6">
          <div className="border-b border-gray-200 pb-4 mb-6">
            <h2 className="text-xl font-semibold text-gray-900 mb-2">
              Query 2: Unsold Non-Promotional Products
            </h2>
            <p className="text-sm text-gray-600">
              Find products in store that have never been sold and are not promotional
            </p>
          </div>

          {/* Arthur2 Parameters */}
          <div className="mb-6">
            <p className="text-sm text-gray-600 mb-4">
              This query requires no parameters. Click the button below to execute.
            </p>

            <button
              onClick={executeArthur2}
              disabled={arthur2Loading}
              className="w-full bg-green-600 text-white py-2 px-4 rounded-md hover:bg-green-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
            >
              {arthur2Loading ? 'Executing...' : 'Execute Query'}
            </button>
          </div>

          {/* Arthur2 Error */}
          {arthur2Error && (
            <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
              {arthur2Error}
            </div>
          )}

          {/* Arthur2 Results */}
          <div className="bg-gray-50 rounded-lg p-4">
            <h3 className="font-medium text-gray-900 mb-3">Results ({arthur2Results.length})</h3>
            {arthur2Results.length > 0 ? (
              <div className="overflow-x-auto">
                <table className="min-w-full bg-white border border-gray-200 rounded">
                  <thead className="bg-gray-100">
                    <tr>
                      <th className="px-3 py-2 text-left text-xs font-medium text-gray-500 uppercase border-b">UPC</th>
                      <th className="px-3 py-2 text-left text-xs font-medium text-gray-500 uppercase border-b">Product Name</th>
                      <th className="px-3 py-2 text-right text-xs font-medium text-gray-500 uppercase border-b">Stock</th>
                      <th className="px-3 py-2 text-left text-xs font-medium text-gray-500 uppercase border-b">Category</th>
                    </tr>
                  </thead>
                  <tbody>
                    {arthur2Results.map((result, index) => (
                      <tr key={index} className="border-b hover:bg-gray-50">
                        <td className="px-3 py-2 text-sm text-gray-900 font-mono">{result.upc}</td>
                        <td className="px-3 py-2 text-sm text-gray-900">{result.product_name}</td>
                        <td className="px-3 py-2 text-sm text-gray-900 text-right">{result.products_number}</td>
                        <td className="px-3 py-2 text-sm text-gray-900">{result.category_name}</td>
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

export default IndividualsArthur;
