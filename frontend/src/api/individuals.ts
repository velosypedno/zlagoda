import axios from './axios';

// Response types for individual queries
export interface Vlad1Response {
  category_id: number;
  category_name: string;
  product_id: number;
  product_name: string;
  total_sales: number;
  total_units_sold: number;
  total_revenue: number;
}

export interface Vlad2Response {
  employee_id: string;
  employee_name: string;
  surname: string;
}

export interface Arthur1Response {
  category_name: string;
  units_sold: number;
  revenue: number;
}

export interface Arthur2Response {
  upc: string;
  product_name: string;
  products_number: number;
  category_name: string;
}

export interface Oleksii1Response {
  employee_id: string;
  employee_surname: string;
  employee_name: string;
  high_discount_customers: number;
  total_receipts_high_discount: number;
  total_revenue_high_discount: number;
  avg_receipt_amount: number;
  avg_customer_discount: number;
}

export interface Oleksii2Response {
  card_number: string;
  cust_surname: string;
  cust_name: string;
  phone_number: string;
}

// Generic response wrapper
export interface IndividualQueryResponse<T> {
  description: string;
  parameters?: Record<string, any>;
  results: T[];
}

// API functions
export const queryVlad1 = async (categoryId: number, months: number = 1): Promise<IndividualQueryResponse<Vlad1Response>> => {
  const response = await axios.get(`/api/vlad1?category_id=${categoryId}&months=${months}`);
  return response.data;
};

export const queryVlad2 = async (): Promise<IndividualQueryResponse<Vlad2Response>> => {
  const response = await axios.get('/api/vlad2');
  return response.data;
};

export const queryArthur1 = async (startDate: string, endDate: string): Promise<IndividualQueryResponse<Arthur1Response>> => {
  const response = await axios.get(`/api/arthur1?start_date=${startDate}&end_date=${endDate}`);
  return response.data;
};

export const queryArthur2 = async (): Promise<IndividualQueryResponse<Arthur2Response>> => {
  const response = await axios.get('/api/arthur2');
  return response.data;
};

export const queryOleksii1 = async (discountThreshold: number = 10): Promise<IndividualQueryResponse<Oleksii1Response>> => {
  const response = await axios.get(`/api/oleksii1?discount_threshold=${discountThreshold}`);
  return response.data;
};

export const queryOleksii2 = async (): Promise<IndividualQueryResponse<Oleksii2Response>> => {
  const response = await axios.get('/api/oleksii2');
  return response.data;
};
