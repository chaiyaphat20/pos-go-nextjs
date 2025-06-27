import { api } from '@/lib/api';
import { Sale, CreateSaleRequest } from '@/types/api';

export const getSales = async (): Promise<Sale[]> => {
  const response = await api.get<Sale[]>('/sales');
  return response.data;
};

export const getSaleById = async (id: number): Promise<Sale> => {
  const response = await api.get<Sale>(`/sales/${id}`);
  return response.data;
};

export const getSalesByCustomer = async (customerId: number): Promise<Sale[]> => {
  const response = await api.get<Sale[]>(`/sales/customer/${customerId}`);
  return response.data;
};

export const createSale = async (saleData: CreateSaleRequest): Promise<Sale> => {
  const response = await api.post<Sale>('/sales', saleData);
  return response.data;
};