import axios from "./axios";
import {type Category} from "../types/category";

export const fetchCategories = () => axios.get<Category[]>("/api/categories");

export const deleteCategory = (id: number) => axios.delete(`/api/categories/${id}`);

export const updateCategory = (id: number, name: string) =>
  axios.patch(`/api/categories/${id}`, { name });

export const createCategory = (name: string) =>
  axios.post<Category>(`/api/categories`, { name });