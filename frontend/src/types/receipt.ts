export interface ReceiptItem {
  upc: string;
  product_number: number;
  selling_price: number;
}

export interface ReceiptCreate {
  employee_id: string;
  card_number?: string;
  print_date: string;
  sum_total: number;
  vat: number;
}

export interface ReceiptCreateComplete {
  employee_id: string;
  card_number?: string | null;
  print_date: string;
  items: ReceiptItem[];
}

export interface ReceiptRetrieve {
  receipt_number: string;
  employee_id: string;
  card_number?: string | null;
  print_date: string;
  sum_total: number;
  vat: number;
}

export interface ReceiptUpdate {
  employee_id?: string;
  card_number?: string;
  print_date?: string;
  sum_total?: number;
  vat?: number;
}

export interface ReceiptCreateResponse {
  receipt_number: string;
  print_date: string;
  total_sum: number;
  vat: number;
}
