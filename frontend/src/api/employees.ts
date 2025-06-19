import axios from "./axios";
import type { Employee } from "../types/employee";

export const fetchEmployees = () => axios.get<Employee[]>("/api/employees");

export const fetchEmployee = (id: string) => axios.get<Employee>(`/api/employees/${id}`);

export const createEmployee = (employee: Omit<Employee, 'employee_id'>) =>
  axios.post<Employee>("/api/employees", employee);

export const updateEmployee = (id: string, employee: Partial<Omit<Employee, 'employee_id'>>) =>
  axios.patch(`/api/employees/${id}`, employee);

export const deleteEmployee = (id: string) =>
  axios.delete(`/api/employees/${id}`); 