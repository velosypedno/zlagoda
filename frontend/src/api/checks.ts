import axios from "./axios";
import type { CheckCreate, CheckCreateResponse, Check } from "../types/check";
import type { SaleWithDetails } from "../types/sale";

export const createCheck = (check: CheckCreate) =>
  axios.post<CheckCreateResponse>("/api/checks", check);

export const fetchChecks = () =>
  axios.get<Check[]>("/api/checks");

export const fetchCheckDetails = (receiptNumber: string) =>
  axios.get<Check>(`/api/checks/${receiptNumber}`);

export const fetchCheckSalesWithDetails = (receiptNumber: string) =>
  axios.get<SaleWithDetails[]>(`/api/sales/by-receipt/${receiptNumber}/details`); 