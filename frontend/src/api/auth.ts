import axios from './axios';

export interface LoginRequest {
  login: string;
  password: string;
}

export interface RegisterRequest {
  login: string;
  password: string;
  surname: string;
  name: string;
  patronymic?: string;
  role: string;
  salary: number;
  date_of_birth: string;
  date_of_start: string;
  phone_number: string;
  city: string;
  street: string;
  zip_code: string;
}

export interface AuthResponse {
  token: string;
}

export interface AccountInfo {
  employee_id?: string;
  login?: string;
  empl_surname?: string;
  empl_name?: string;
  empl_patronymic?: string;
  empl_role?: string;
  salary?: number;
  date_of_birth?: string;
  date_of_start?: string;
  phone_number?: string;
  city?: string;
  street?: string;
  zip_code?: string;
}

export const login = async (data: LoginRequest): Promise<AuthResponse> => {
  const response = await axios.post('/api/login', data);
  return response.data;
};

export const register = async (data: RegisterRequest): Promise<AuthResponse> => {
  const response = await axios.post('/api/register', data);
  return response.data;
};

export const getAccount = async (): Promise<AccountInfo> => {
  const response = await axios.get('/api/account');
  return response.data;
}; 