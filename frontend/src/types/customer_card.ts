export interface CustomerCard {
  card_number: string;
  cust_surname: string;
  cust_name: string;
  cust_patronymic?: string;
  phone_number: string;
  city?: string;
  street?: string;
  zip_code?: string;
  percent: number;
}

export type CustomerCardCreate = Omit<CustomerCard, 'card_number'>;
export type CustomerCardUpdate = Partial<CustomerCardCreate>; 