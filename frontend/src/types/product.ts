export interface Product {
  id: number;
  name: string;
  characteristics: string;
  category_id: number;
}

export interface ProductCreate {
  name: string;
  characteristics: string;
  category_id: number;
}

export interface ProductUpdate {
  name?: string;
  characteristics?: string;
  category_id?: number;
} 