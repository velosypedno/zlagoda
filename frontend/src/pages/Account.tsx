import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { getAccount, type AccountInfo } from '../api/auth';

const Account = () => {
  const [account, setAccount] = useState<AccountInfo | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    const fetchAccount = async () => {
      try {
        const accountData = await getAccount();
        setAccount(accountData);
      } catch (err: any) {
        setError(err.response?.data?.error || 'Failed to load account information');
      } finally {
        setLoading(false);
      }
    };

    fetchAccount();
  }, []);

  const handleLogout = () => {
    localStorage.removeItem('token');
    navigate('/login');
  };

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString();
  };

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="animate-spin rounded-full h-32 w-32 border-b-2 border-blue-600"></div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="text-center">
          <div className="text-red-600 text-lg mb-4">{error}</div>
          <button
            onClick={() => window.location.reload()}
            className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
          >
            Try Again
          </button>
        </div>
      </div>
    );
  }

  if (!account) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="text-center">
          <div className="text-gray-600 text-lg">No account information found</div>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-3xl mx-auto">
        <div className="bg-white shadow rounded-lg">
          {/* Header */}
          <div className="px-6 py-4 border-b border-gray-200">
            <div className="flex items-center justify-between">
              <h1 className="text-2xl font-bold text-gray-900">My Account</h1>
              <button
                onClick={handleLogout}
                className="bg-red-600 text-white px-4 py-2 rounded-md hover:bg-red-700 transition-colors"
              >
                Logout
              </button>
            </div>
          </div>

          {/* Account Information */}
          <div className="px-6 py-6">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
              {/* Personal Information */}
              <div>
                <h2 className="text-lg font-semibold text-gray-900 mb-4">Personal Information</h2>
                <div className="space-y-3">
                  <div>
                    <label className="block text-sm font-medium text-gray-600">Employee ID</label>
                    <p className="mt-1 text-sm text-gray-900">{account.employee_id || 'Not specified'}</p>
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-gray-600">Name</label>
                    <p className="mt-1 text-sm text-gray-900">
                      {[account.empl_surname, account.empl_name, account.empl_patronymic]
                        .filter(Boolean)
                        .join(' ') || 'Not specified'}
                    </p>
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-gray-600">Role</label>
                    <p className="mt-1 text-sm text-gray-900">{account.empl_role || 'Not specified'}</p>
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-gray-600">Salary</label>
                    <p className="mt-1 text-sm text-gray-900">
                      {account.salary ? `$${account.salary.toFixed(2)}` : 'Not specified'}
                    </p>
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-gray-600">Date of Birth</label>
                    <p className="mt-1 text-sm text-gray-900">
                      {account.date_of_birth ? formatDate(account.date_of_birth) : 'Not specified'}
                    </p>
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-gray-600">Date of Hire</label>
                    <p className="mt-1 text-sm text-gray-900">
                      {account.date_of_start ? formatDate(account.date_of_start) : 'Not specified'}
                    </p>
                  </div>
                </div>
              </div>

              {/* Contact Information */}
              <div>
                <h2 className="text-lg font-semibold text-gray-900 mb-4">Contact Information</h2>
                <div className="space-y-3">
                  <div>
                    <label className="block text-sm font-medium text-gray-600">Phone Number</label>
                    <p className="mt-1 text-sm text-gray-900">{account.phone_number || 'Not specified'}</p>
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-gray-600">Address</label>
                    <p className="mt-1 text-sm text-gray-900">
                      {[account.street, account.city, account.zip_code]
                        .filter(Boolean)
                        .join(', ') || 'Not specified'}
                    </p>
                  </div>
                </div>
              </div>
            </div>

            {/* Account Actions */}
            <div className="mt-8 pt-6 border-t border-gray-200">
              <h2 className="text-lg font-semibold text-gray-900 mb-4">Account Actions</h2>
              <div className="flex flex-wrap gap-3">
                <button
                  onClick={handleLogout}
                  className="bg-red-600 text-white px-4 py-2 rounded-md hover:bg-red-700 transition-colors"
                >
                  Logout
                </button>
                <button
                  onClick={() => navigate('/')}
                  className="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors"
                >
                  Back to Dashboard
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Account; 