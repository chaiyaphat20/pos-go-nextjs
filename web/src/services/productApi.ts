import { api } from '@/lib/api';
import { Product, CreateProductRequest, UpdateProductRequest, UpdateStockRequest } from '@/types/api';

export const getProducts = async (): Promise<Product[]> => {
  const response = await api.get<Product[]>('/products');
  return response.data;
};

export const getProductById = async (id: number): Promise<Product> => {
  const response = await api.get<Product>(`/products/${id}`);
  return response.data;
};

export const createProduct = async (productData: CreateProductRequest): Promise<Product> => {
  const response = await api.post<Product>('/products', productData);
  return response.data;
};

export const updateProduct = async (id: number, productData: UpdateProductRequest): Promise<Product> => {
  const response = await api.put<Product>(`/products/${id}`, productData);
  return response.data;
};

export const updateProductStock = async (id: number, stockData: UpdateStockRequest): Promise<void> => {
  await api.patch(`/products/${id}/stock`, stockData);
};

export const deleteProduct = async (id: number): Promise<void> => {
  await api.delete(`/products/${id}`);
};