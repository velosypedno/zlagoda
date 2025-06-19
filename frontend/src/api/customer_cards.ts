import type { CustomerCard, CustomerCardCreate, CustomerCardUpdate } from '../types/customer_card';
import { API_BASE_URL } from './config';

export async function getCustomerCards(): Promise<CustomerCard[]> {
  const response = await fetch(`${API_BASE_URL}/customer-cards`);
  if (!response.ok) {
    throw new Error('Failed to fetch customer cards');
  }
  return response.json();
}

export async function getCustomerCard(cardNumber: string): Promise<CustomerCard> {
  const response = await fetch(`${API_BASE_URL}/customer-cards/${cardNumber}`);
  if (!response.ok) {
    throw new Error('Failed to fetch customer card');
  }
  return response.json();
}

export async function createCustomerCard(card: CustomerCardCreate): Promise<string> {
  const response = await fetch(`${API_BASE_URL}/customer-cards`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(card),
  });
  if (!response.ok) {
    throw new Error('Failed to create customer card');
  }
  const data = await response.json();
  return data.id;
}

export async function updateCustomerCard(cardNumber: string, card: CustomerCardUpdate): Promise<void> {
  const response = await fetch(`${API_BASE_URL}/customer-cards/${cardNumber}`, {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(card),
  });
  if (!response.ok) {
    throw new Error('Failed to update customer card');
  }
}

export async function deleteCustomerCard(cardNumber: string): Promise<void> {
  const response = await fetch(`${API_BASE_URL}/customer-cards/${cardNumber}`, {
    method: 'DELETE',
  });
  if (!response.ok) {
    throw new Error('Failed to delete customer card');
  }
} 