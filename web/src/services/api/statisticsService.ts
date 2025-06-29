import { httpClient } from '@/lib/httpClient';
import type { 
  OrderStatistics, 
  ProductStatistics, 
  UserStatistics, 
  DashboardStatistics 
} from '@/types/statistics';

export class StatisticsService {
  async getDashboardStatistics(): Promise<DashboardStatistics> {
    const response = await httpClient.get<DashboardStatistics>('/statistics/dashboard');
    
    if (response.success && response.data) {
      return response.data;
    }
    
    throw new Error(response.message || 'Failed to fetch dashboard statistics');
  }

  async getOrderStatistics(): Promise<OrderStatistics> {
    const response = await httpClient.get<OrderStatistics>('/statistics/orders');
    
    if (response.success && response.data) {
      return response.data;
    }
    
    throw new Error(response.message || 'Failed to fetch order statistics');
  }

  async getProductStatistics(): Promise<ProductStatistics> {
    const response = await httpClient.get<ProductStatistics>('/statistics/products');
    
    if (response.success && response.data) {
      return response.data;
    }
    
    throw new Error(response.message || 'Failed to fetch product statistics');
  }

  async getUserStatistics(): Promise<UserStatistics> {
    const response = await httpClient.get<UserStatistics>('/statistics/users');
    
    if (response.success && response.data) {
      return response.data;
    }
    
    throw new Error(response.message || 'Failed to fetch user statistics');
  }

  async getRevenueStatistics(params?: { 
    period?: 'day' | 'week' | 'month' | 'year';
    start_date?: string;
    end_date?: string;
  }) {
    const response = await httpClient.get('/statistics/revenue', params);
    
    if (response.success) {
      return response.data;
    }
    
    throw new Error(response.message || 'Failed to fetch revenue statistics');
  }

  async getSalesTrend(params?: {
    period?: 'day' | 'week' | 'month';
    limit?: number;
  }) {
    const response = await httpClient.get('/statistics/sales-trend', params);
    
    if (response.success) {
      return response.data;
    }
    
    throw new Error(response.message || 'Failed to fetch sales trend');
  }
}