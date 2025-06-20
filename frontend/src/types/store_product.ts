export interface StoreProduct {
  upc: string;
  upc_prom?: string;
  product_id: number;
  selling_price: number;
  products_number: number;
  promotional_product: boolean;
}

export interface StoreProductCreate {
  upc: string;
  upc_prom?: string;
  product_id: number;
  selling_price: number;
  products_number: number;
  promotional_product: boolean;
}

export interface StoreProductUpdate {
  upc_prom?: string;
  product_id?: number;
  selling_price?: number;
  products_number?: number;
  promotional_product?: boolean;
}

export interface StoreProductWithDetails extends StoreProduct {
  product_name: string;
  category_name: string;
} 