import axios from "./axios";
import type {
  ReceiptCreate,
  ReceiptCreateComplete,
  ReceiptCreateResponse,
  ReceiptRetrieve,
  ReceiptUpdate
} from "../types/receipt";
import type { SaleWithDetails } from "../types/sale";

export const createReceipt = (receipt: ReceiptCreate) =>
  axios.post<{ id: string }>("/api/receipts", receipt);

export const createReceiptComplete = (receipt: ReceiptCreateComplete) =>
  axios.post<{ id: string }>("/api/receipts/complete", receipt);

export const fetchReceipts = () =>
  axios.get<ReceiptRetrieve[]>("/api/receipts");

export const fetchReceiptDetails = (receiptNumber: string) =>
  axios.get<ReceiptRetrieve>(`/api/receipts/${receiptNumber}`);

export const updateReceipt = (receiptNumber: string, receipt: ReceiptUpdate) =>
  axios.patch(`/api/receipts/${receiptNumber}`, receipt);

export const deleteReceipt = (receiptNumber: string) =>
  axios.delete(`/api/receipts/${receiptNumber}`);

export const fetchReceiptSalesWithDetails = (receiptNumber: string) =>
  axios.get<SaleWithDetails[]>(`/api/sales/by-receipt/${receiptNumber}/details`);
