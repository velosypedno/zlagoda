import { useState } from "react";
import type { Employee } from "../types/employee";

interface Props {
  employee: Employee;
  onDelete: () => void;
  onUpdate: (employee: Partial<Employee>) => void;
}

const EmployeeCard = ({ employee, onDelete, onUpdate }: Props) => {
  const [isEditing, setIsEditing] = useState(false);
  const [form, setForm] = useState<Partial<Employee>>({ ...employee });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSave = () => {
    onUpdate(form);
    setIsEditing(false);
  };

  const handleCancel = () => {
    setForm({ ...employee });
    setIsEditing(false);
  };

  return (
    <div className="bg-white border border-gray-200 shadow-sm rounded-lg p-4">
      <div className="flex items-center justify-between mb-2">
        <div className="font-semibold text-lg text-gray-800">
          {employee.empl_surname || ''} {employee.empl_name || ''}
        </div>
        <span className="text-xs px-2 py-1 rounded bg-gray-100 text-gray-600 font-medium">
          {(employee.empl_role || '').charAt(0).toUpperCase() + (employee.empl_role || '').slice(1)}
        </span>
      </div>
      {isEditing ? (
        <div className="grid grid-cols-1 gap-2 mb-2">
          <input name="empl_surname" value={form.empl_surname || ""} onChange={handleChange} className="border p-1 rounded" placeholder="Surname" />
          <input name="empl_name" value={form.empl_name || ""} onChange={handleChange} className="border p-1 rounded" placeholder="Name" />
          <input name="empl_patronymic" value={form.empl_patronymic || ""} onChange={handleChange} className="border p-1 rounded" placeholder="Patronymic" />
          <select name="empl_role" value={form.empl_role || ""} onChange={handleChange} className="border p-1 rounded">
            <option value="">Role</option>
            <option value="manager">Manager</option>
            <option value="cashier">Cashier</option>
          </select>
          <input name="salary" type="number" value={form.salary || ""} onChange={handleChange} className="border p-1 rounded" placeholder="Salary" />
          <input name="date_of_birth" type="date" value={form.date_of_birth || ""} onChange={handleChange} className="border p-1 rounded" placeholder="Date of Birth" />
          <input name="date_of_start" type="date" value={form.date_of_start || ""} onChange={handleChange} className="border p-1 rounded" placeholder="Date of Start" />
          <input name="phone_number" value={form.phone_number || ""} onChange={handleChange} className="border p-1 rounded" placeholder="Phone Number" />
          <input name="city" value={form.city || ""} onChange={handleChange} className="border p-1 rounded" placeholder="City" />
          <input name="street" value={form.street || ""} onChange={handleChange} className="border p-1 rounded" placeholder="Street" />
          <input name="zip_code" value={form.zip_code || ""} onChange={handleChange} className="border p-1 rounded" placeholder="Zip Code" />
        </div>
      ) : (
        <div className="text-sm text-gray-700 space-y-1 mb-2">
          <div><span className="font-medium">ID:</span> {employee.employee_id || 'N/A'}</div>
          <div><span className="font-medium">Salary:</span> {employee.salary || 'N/A'}</div>
          <div><span className="font-medium">DOB:</span> {employee.date_of_birth || 'N/A'}</div>
          <div><span className="font-medium">Start:</span> {employee.date_of_start || 'N/A'}</div>
          <div><span className="font-medium">Phone:</span> {employee.phone_number || 'N/A'}</div>
          <div><span className="font-medium">City:</span> {employee.city || 'N/A'}</div>
          <div><span className="font-medium">Street:</span> {employee.street || 'N/A'}</div>
          <div><span className="font-medium">Zip:</span> {employee.zip_code || 'N/A'}</div>
        </div>
      )}
      <div className="flex gap-2 mt-2 justify-end">
        {isEditing ? (
          <>
            <button className="bg-green-500 text-white px-3 py-1 rounded" onClick={handleSave}>Save</button>
            <button className="bg-gray-300 text-gray-800 px-3 py-1 rounded" onClick={handleCancel}>Cancel</button>
            <button className="bg-red-500 text-white px-3 py-1 rounded" onClick={onDelete}>Delete</button>
          </>
        ) : (
          <>
            <button className="bg-blue-500 text-white px-3 py-1 rounded" onClick={() => setIsEditing(true)}>Edit</button>
            <button className="bg-red-500 text-white px-3 py-1 rounded" onClick={onDelete}>Delete</button>
          </>
        )}
      </div>
    </div>
  );
};

export default EmployeeCard; 