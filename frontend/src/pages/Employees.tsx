import { useEffect, useState } from "react";
import { fetchEmployees, deleteEmployee, updateEmployee, createEmployee } from "../api/employees";
import type { Employee } from "../types/employee";
import EmployeeCard from "../components/EmployeeCard";

const EmployeesPage = () => {
  const [employees, setEmployees] = useState<Employee[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [newEmployee, setNewEmployee] = useState<Omit<Employee, 'employee_id'>>({
    empl_surname: "",
    empl_name: "",
    empl_patronymic: "",
    empl_role: "",
    salary: 0,
    date_of_birth: "",
    date_of_start: "",
    phone_number: "",
    city: "",
    street: "",
    zip_code: "",
  });

  const loadEmployees = async () => {
    try {
      const res = await fetchEmployees();
      const data = Array.isArray(res.data) ? res.data : [];
      setEmployees(data);
    } catch (err) {
      setEmployees([]);
      setError("Failed to load employees");
    }
  };

  useEffect(() => {
    loadEmployees();
  }, []);

  const handleDelete = async (id: string) => {
    try {
      await deleteEmployee(id);
      setEmployees((prev) => prev.filter((emp) => emp.employee_id !== id));
    } catch {
      setError("Failed to delete employee");
    }
  };

  const handleUpdate = async (id: string, data: Partial<Employee>) => {
    try {
      await updateEmployee(id, data);
      setEmployees((prev) =>
        prev.map((emp) => (emp.employee_id === id ? { ...emp, ...data } : emp))
      );
    } catch {
      setError("Failed to update employee");
    }
  };

  const handleCreate = async () => {
    // Basic validation
    if (!newEmployee.empl_surname || !newEmployee.empl_name || !newEmployee.empl_role || !newEmployee.salary || !newEmployee.date_of_birth || !newEmployee.date_of_start || !newEmployee.phone_number || !newEmployee.city || !newEmployee.street || !newEmployee.zip_code) {
      setError("Please fill all required fields");
      return;
    }
    try {
      const res = await createEmployee(newEmployee);
      setEmployees((prev) => [...prev, res.data]);
      setNewEmployee({
        empl_surname: "",
        empl_name: "",
        empl_patronymic: "",
        empl_role: "",
        salary: 0,
        date_of_birth: "",
        date_of_start: "",
        phone_number: "",
        city: "",
        street: "",
        zip_code: "",
      });
      setError(null);
      await loadEmployees();
    } catch {
      setError("Failed to create employee");
    }
  };

  const handleNewChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value, type } = e.target;
    setNewEmployee({
      ...newEmployee,
      [name]: name === "salary" ? Number(value) : value,
    });
  };

  return (
    <div className="p-6 max-w-6xl mx-auto">
      <h1 className="text-3xl font-bold mb-6">Employees</h1>
      {error && <div className="mb-4 text-red-500">{error}</div>}
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-8">
        <input name="empl_surname" value={newEmployee.empl_surname} onChange={handleNewChange} placeholder="Surname" className="border rounded px-3 py-2" />
        <input name="empl_name" value={newEmployee.empl_name} onChange={handleNewChange} placeholder="Name" className="border rounded px-3 py-2" />
        <input name="empl_patronymic" value={newEmployee.empl_patronymic} onChange={handleNewChange} placeholder="Patronymic" className="border rounded px-3 py-2" />
        <select name="empl_role" value={newEmployee.empl_role} onChange={handleNewChange} className="border rounded px-3 py-2">
          <option value="">Role</option>
          <option value="manager">Manager</option>
          <option value="cashier">Cashier</option>
        </select>
        <input name="salary" type="number" value={newEmployee.salary} onChange={handleNewChange} placeholder="Salary" className="border rounded px-3 py-2" />
        <input name="date_of_birth" type="date" value={newEmployee.date_of_birth} onChange={handleNewChange} placeholder="Date of Birth" className="border rounded px-3 py-2" />
        <input name="date_of_start" type="date" value={newEmployee.date_of_start} onChange={handleNewChange} placeholder="Date of Start" className="border rounded px-3 py-2" />
        <input name="phone_number" value={newEmployee.phone_number} onChange={handleNewChange} placeholder="Phone Number" className="border rounded px-3 py-2" />
        <input name="city" value={newEmployee.city} onChange={handleNewChange} placeholder="City" className="border rounded px-3 py-2" />
        <input name="street" value={newEmployee.street} onChange={handleNewChange} placeholder="Street" className="border rounded px-3 py-2" />
        <input name="zip_code" value={newEmployee.zip_code} onChange={handleNewChange} placeholder="Zip Code" className="border rounded px-3 py-2" />
      </div>
      <button onClick={handleCreate} className="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700 disabled:opacity-50 mb-8">Add Employee</button>
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
        {Array.isArray(employees) && employees.map((employee) => (
          <EmployeeCard
            key={employee.employee_id}
            employee={employee}
            onDelete={() => handleDelete(employee.employee_id)}
            onUpdate={(data) => handleUpdate(employee.employee_id, data)}
          />
        ))}
      </div>
    </div>
  );
};

export default EmployeesPage; 