import { httpClient } from '@/lib/httpClient';
import type { Order, OrderStatus } from '@/types/database';
import type { OrderStatistics } from '@/types/statistics';
import type { 
  CreateOrderRequest, 
  UpdateOrderRequest,
  OrderQuery
} from '@/types/api/order';

export class OrderService {
  async getOrders(params: OrderQuery = {}): Promise<Order[]> {
    const response = await httpClient.get<Order[]>('/orders', params);
    
    if (response.success) {
      return response.data || [];
    }
    
    throw new Error(response.message || 'Failed to fetch orders');
  }

  async getOrderById(id: string): Promise<Order> {
    const response = await httpClient.get<Order>(`/orders/${id}`);
    
    if (response.success && response.data) {
      return response.data;
    }
    
    throw new Error(response.message || 'Failed to fetch order');
  }

  async createOrder(orderData: CreateOrderRequest): Promise<Order> {
    const response = await httpClient.post<Order>('/orders', orderData);
    
    if (response.success && response.data) {
      return response.data;
    }
    
    throw new Error(response.message || 'Failed to create order');
  }

  async updateOrder(id: string, orderData: UpdateOrderRequest): Promise<Order> {
    const response = await httpClient.put<Order>(`/orders/${id}`, orderData);
    
    if (response.success && response.data) {
      return response.data;
    }
    
    throw new Error(response.message || 'Failed to update order');
  }

  async updateOrderStatus(id: string, status: OrderStatus): Promise<Order> {
    const response = await httpClient.patch<Order>(`/orders/${id}/status`, { status });
    
    if (response.success && response.data) {
      return response.data;
    }
    
    throw new Error(response.message || 'Failed to update order status');
  }

  async cancelOrder(id: string): Promise<void> {
    const response = await httpClient.delete(`/orders/${id}`);
    
    if (!response.success) {
      throw new Error(response.message || 'Failed to cancel order');
    }
  }

  async getOrdersByUser(userId: string): Promise<Order[]> {
    const response = await httpClient.get<Order[]>(`/users/${userId}/orders`);
    
    if (response.success) {
      return response.data || [];
    }
    
    throw new Error(response.message || 'Failed to fetch user orders');
  }

  async getMyOrders(): Promise<Order[]> {
    const response = await httpClient.get<Order[]>('/orders/me');
    
    if (response.success) {
      return response.data || [];
    }
    
    throw new Error(response.message || 'Failed to fetch my orders');
  }

  async getOrderStatistics(): Promise<OrderStatistics> {
    const response = await httpClient.get<OrderStatistics>('/orders/statistics');
    
    if (response.success && response.data) {
      return response.data;
    }
    
    throw new Error(response.message || 'Failed to fetch order statistics');
  }
}