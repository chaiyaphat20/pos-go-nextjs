import { useState, useCallback } from 'react';
import { orderService } from '@/services/api';
import type { Order, OrderStatus } from '@/types/database';
import type { OrderStatistics } from '@/types/statistics';
import type { CreateOrderRequest, UpdateOrderRequest, OrderQuery } from '@/types/api/order';
import type { UseOrdersReturn } from '@/types/api/hooks';

export function useOrders(): UseOrdersReturn {
  const [orders, setOrders] = useState<Order[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchOrders = useCallback(async (params: OrderQuery = {}): Promise<void> => {
    setLoading(true);
    setError(null);
    
    try {
      const fetchedOrders = await orderService.getOrders(params);
      setOrders(fetchedOrders);
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to fetch orders';
      setError(errorMessage);
    } finally {
      setLoading(false);
    }
  }, []);

  const createOrder = useCallback(async (orderData: CreateOrderRequest): Promise<Order> => {
    setLoading(true);
    setError(null);
    
    try {
      const newOrder = await orderService.createOrder(orderData);
      setOrders(prev => [newOrder, ...prev]);
      return newOrder;
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to create order';
      setError(errorMessage);
      throw err;
    } finally {
      setLoading(false);
    }
  }, []);

  const updateOrder = useCallback(async (id: string, data: UpdateOrderRequest): Promise<Order> => {
    setLoading(true);
    setError(null);
    
    try {
      const updatedOrder = await orderService.updateOrder(id, data);
      setOrders(prev => prev.map(order => order.id === id ? updatedOrder : order));
      return updatedOrder;
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to update order';
      setError(errorMessage);
      throw err;
    } finally {
      setLoading(false);
    }
  }, []);

  const updateOrderStatus = useCallback(async (id: string, status: OrderStatus): Promise<Order> => {
    setLoading(true);
    setError(null);
    
    try {
      const updatedOrder = await orderService.updateOrderStatus(id, status);
      setOrders(prev => prev.map(order => order.id === id ? updatedOrder : order));
      return updatedOrder;
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to update order status';
      setError(errorMessage);
      throw err;
    } finally {
      setLoading(false);
    }
  }, []);

  const cancelOrder = useCallback(async (id: string): Promise<boolean> => {
    setLoading(true);
    setError(null);
    
    try {
      await orderService.cancelOrder(id);
      setOrders(prev => prev.filter(order => order.id !== id));
      return true;
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to cancel order';
      setError(errorMessage);
      throw err;
    } finally {
      setLoading(false);
    }
  }, []);

  const getOrderById = useCallback(async (id: string): Promise<Order> => {
    setLoading(true);
    setError(null);
    
    try {
      return await orderService.getOrderById(id);
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to fetch order';
      setError(errorMessage);
      throw err;
    } finally {
      setLoading(false);
    }
  }, []);

  const getMyOrders = useCallback(async (): Promise<Order[]> => {
    setLoading(true);
    setError(null);
    
    try {
      return await orderService.getMyOrders();
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to fetch my orders';
      setError(errorMessage);
      throw err;
    } finally {
      setLoading(false);
    }
  }, []);

  const getOrdersByUser = useCallback(async (userId: string): Promise<Order[]> => {
    setLoading(true);
    setError(null);
    
    try {
      return await orderService.getOrdersByUser(userId);
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to fetch user orders';
      setError(errorMessage);
      throw err;
    } finally {
      setLoading(false);
    }
  }, []);

  const getOrderStatistics = useCallback(async (): Promise<OrderStatistics> => {
    setLoading(true);
    setError(null);
    
    try {
      return await orderService.getOrderStatistics();
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to fetch order statistics';
      setError(errorMessage);
      throw err;
    } finally {
      setLoading(false);
    }
  }, []);

  return {
    orders,
    loading,
    error,
    fetchOrders,
    createOrder,
    updateOrder,
    updateOrderStatus,
    cancelOrder,
    getOrderById,
    getMyOrders,
    getOrdersByUser,
    getOrderStatistics,
    clearError: () => setError(null),
  };
}