import type { CustomerCard, CustomerCardCreate, CustomerCardUpdate } from '../types/customer_card';
import axios from './axios';

export async function getCustomerCards(): Promise<CustomerCard[]> {
  const response = await axios.get('/api/customer-cards');
  return response.data;
}

export async function getCustomerCard(cardNumber: string): Promise<CustomerCard> {
  const response = await axios.get(`/api/customer-cards/${cardNumber}`);
  return response.data;
}

export async function createCustomerCard(card: CustomerCardCreate): Promise<string> {
  const response = await axios.post('/api/customer-cards', card);
  return response.data.id;
}

export async function updateCustomerCard(cardNumber: string, card: CustomerCardUpdate): Promise<void> {
  await axios.patch(`/api/customer-cards/${cardNumber}`, card);
}

export async function deleteCustomerCard(cardNumber: string): Promise<void> {
  await axios.delete(`/api/customer-cards/${cardNumber}`);
} 