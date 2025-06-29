import { useState, useCallback } from 'react';
import { productService } from '@/services/api';
import type { Product } from '@/types/database';
import type { CreateProductRequest, UpdateProductRequest, ProductQuery, UpdateStockRequest } from '@/types/api/product';
import type { UseProductsReturn } from '@/types/api/hooks';

export function useProducts(): UseProductsReturn {
  const [products, setProducts] = useState<Product[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchProducts = useCallback(async (params: ProductQuery = {}): Promise<void> => {
    setLoading(true);
    setError(null);
    
    try {
      const fetchedProducts = await productService.getProducts(params);
      setProducts(fetchedProducts);
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to fetch products';
      setError(errorMessage);
    } finally {
      setLoading(false);
    }
  }, []);

  const createProduct = useCallback(async (productData: CreateProductRequest): Promise<Product> => {
    setLoading(true);
    setError(null);
    
    try {
      const newProduct = await productService.createProduct(productData);
      setProducts(prev => [...prev, newProduct]);
      return newProduct;
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to create product';
      setError(errorMessage);
      throw err;
    } finally {
      setLoading(false);
    }
  }, []);

  const updateProduct = useCallback(async (id: string, productData: UpdateProductRequest): Promise<Product> => {
    setLoading(true);
    setError(null);
    
    try {
      const updatedProduct = await productService.updateProduct(id, productData);
      setProducts(prev => prev.map(product => product.id === id ? updatedProduct : product));
      return updatedProduct;
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to update product';
      setError(errorMessage);
      throw err;
    } finally {
      setLoading(false);
    }
  }, []);

  const deleteProduct = useCallback(async (id: string): Promise<boolean> => {
    setLoading(true);
    setError(null);
    
    try {
      await productService.deleteProduct(id);
      setProducts(prev => prev.filter(product => product.id !== id));
      return true;
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to delete product';
      setError(errorMessage);
      throw err;
    } finally {
      setLoading(false);
    }
  }, []);

  const updateStock = useCallback(async (id: string, data: UpdateStockRequest): Promise<Product> => {
    setLoading(true);
    setError(null);
    
    try {
      const updatedProduct = await productService.updateProductStock(id, data);
      setProducts(prev => prev.map(product => product.id === id ? updatedProduct : product));
      return updatedProduct;
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to update stock';
      setError(errorMessage);
      throw err;
    } finally {
      setLoading(false);
    }
  }, []);

  const getProductById = useCallback(async (id: string): Promise<Product> => {
    setLoading(true);
    setError(null);
    
    try {
      return await productService.getProductById(id);
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to fetch product';
      setError(errorMessage);
      throw err;
    } finally {
      setLoading(false);
    }
  }, []);

  const searchProducts = useCallback(async (query: string, limit = 10): Promise<Product[]> => {
    setLoading(true);
    setError(null);
    
    try {
      return await productService.searchProducts(query, limit);
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to search products';
      setError(errorMessage);
      throw err;
    } finally {
      setLoading(false);
    }
  }, []);

  const getLowStockProducts = useCallback(async (threshold = 10): Promise<Product[]> => {
    setLoading(true);
    setError(null);
    
    try {
      return await productService.getLowStockProducts(threshold);
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to fetch low stock products';
      setError(errorMessage);
      throw err;
    } finally {
      setLoading(false);
    }
  }, []);

  return {
    products,
    loading,
    error,
    fetchProducts,
    createProduct,
    updateProduct,
    deleteProduct,
    updateStock,
    getProductById,
    searchProducts,
    getLowStockProducts,
    clearError: () => setError(null),
  };
}