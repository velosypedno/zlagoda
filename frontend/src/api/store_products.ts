import axios from "./axios";
import { 
  type StoreProduct, 
  type StoreProductCreate, 
  type StoreProductUpdate,
  type StoreProductWithDetails 
} from "../types/store_product";

export const fetchStoreProducts = () => axios.get<StoreProduct[]>("/api/store-products");

export const fetchStoreProductsWithDetails = () => 
  axios.get<StoreProductWithDetails[]>("/api/store-products/details");

export const fetchStoreProduct = (upc: string) => 
  axios.get<StoreProduct>(`/api/store-products/${upc}`);

export const fetchStoreProductsByProduct = (productId: number) =>
  axios.get<StoreProduct[]>(`/api/store-products/by-product/${productId}`);

export const fetchPromotionalProducts = () =>
  axios.get<StoreProduct[]>("/api/store-products/promotional");

export const createStoreProduct = (storeProduct: StoreProductCreate) =>
  axios.post<StoreProduct>(`/api/store-products`, storeProduct);

export const updateStoreProduct = (upc: string, storeProduct: StoreProductUpdate) =>
  axios.patch<StoreProduct>(`/api/store-products/${upc}`, storeProduct);

export const deleteStoreProduct = (upc: string) => 
  axios.delete(`/api/store-products/${upc}`);

export const updateProductQuantity = (upc: string, quantityChange: number) =>
  axios.patch(`/api/store-products/${upc}/quantity`, { quantity_change: quantityChange });

export const checkStockAvailability = (upc: string, quantity: number) =>
  axios.get(`/api/store-products/${upc}/stock-check?quantity=${quantity}`); 