import { httpClient } from '@/lib/httpClient';
import type { Product } from '@/types/database';
import type { 
  CreateProductRequest, 
  UpdateProductRequest,
  ProductQuery,
  UpdateStockRequest 
} from '@/types/api/product';

export class ProductService {
  async getProducts(params: ProductQuery = {}): Promise<Product[]> {
    const response = await httpClient.get<Product[]>('/products', params);
    console.log({response})
    if (response.success) {
      return response.data || [];
    }
    
    throw new Error(response.message || 'Failed to fetch products');
  }

  async getProductById(id: string): Promise<Product> {
    const response = await httpClient.get<Product>(`/products/${id}`);
    
    if (response.success && response.data) {
      return response.data;
    }
    
    throw new Error(response.message || 'Failed to fetch product');
  }

  async createProduct(productData: CreateProductRequest): Promise<Product> {
    const response = await httpClient.post<Product>('/products', productData);
    
    if (response.success && response.data) {
      return response.data;
    }
    
    throw new Error(response.message || 'Failed to create product');
  }

  async updateProduct(id: string, productData: UpdateProductRequest): Promise<Product> {
    const response = await httpClient.put<Product>(`/products/${id}`, productData);
    
    if (response.success && response.data) {
      return response.data;
    }
    
    throw new Error(response.message || 'Failed to update product');
  }

  async deleteProduct(id: string): Promise<void> {
    const response = await httpClient.delete(`/products/${id}`);
    
    if (!response.success) {
      throw new Error(response.message || 'Failed to delete product');
    }
  }

  async updateProductStock(id: string, data: UpdateStockRequest): Promise<Product> {
    const response = await httpClient.patch<Product>(`/products/${id}/stock`, data);
    
    if (response.success && response.data) {
      return response.data;
    }
    
    throw new Error(response.message || 'Failed to update product stock');
  }

  async getLowStockProducts(threshold = 10): Promise<Product[]> {
    const response = await httpClient.get<Product[]>('/products/low-stock', { threshold });
    
    if (response.success) {
      return response.data || [];
    }
    
    throw new Error(response.message || 'Failed to fetch low stock products');
  }

  async searchProducts(query: string, limit = 10): Promise<Product[]> {
    const response = await httpClient.get<Product[]>('/products/search', { q: query, limit });
    
    if (response.success) {
      return response.data || [];
    }
    
    throw new Error(response.message || 'Failed to search products');
  }
}