export interface CheckSaleItem {
  upc: string;
  product_number: number;
  selling_price: number;
}

export interface CheckCreate {
  employee_id: string;
  card_number?: string;
  print_date: string;
  items: CheckSaleItem[];
}

export interface CheckCreateResponse {
  receipt_number: string;
  print_date: string;
  total_sum: number;
  vat: number;
} 