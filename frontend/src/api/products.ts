import axios from "./axios";
import { type Product, type ProductCreate, type ProductUpdate } from "../types/product";

export const fetchProducts = () => axios.get<Product[]>("/api/products");

export const fetchProductsByCategory = (categoryId: number) => 
  axios.get<Product[]>(`/api/products/by-category/${categoryId}`);

export const fetchProductsByName = (name: string) => 
  axios.get<Product[]>(`/api/products/search?name=${encodeURIComponent(name)}`);

export const fetchProduct = (id: number) => axios.get<Product>(`/api/products/${id}`);

export const createProduct = (product: ProductCreate) =>
  axios.post<Product>(`/api/products`, product);

export const updateProduct = (id: number, product: ProductUpdate) =>
  axios.patch<Product>(`/api/products/${id}`, product);

export const deleteProduct = (id: number) => axios.delete(`/api/products/${id}`); 