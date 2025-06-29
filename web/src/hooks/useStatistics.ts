import { useState, useCallback } from 'react';
import { statisticsService } from '@/services/api';
import type { 
  DashboardStatistics, 
  OrderStatistics, 
  ProductStatistics, 
  UserStatistics
} from '@/types/statistics';
import type { UseStatisticsReturn } from '@/types/api/hooks';

export function useStatistics(): UseStatisticsReturn {
  const [statistics, setStatistics] = useState<DashboardStatistics | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchDashboardStats = useCallback(async (): Promise<void> => {
    setLoading(true);
    setError(null);
    
    try {
      const stats = await statisticsService.getDashboardStatistics();
      setStatistics(stats);
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to fetch dashboard statistics';
      setError(errorMessage);
    } finally {
      setLoading(false);
    }
  }, []);

  const fetchOrderStats = useCallback(async (): Promise<OrderStatistics> => {
    setLoading(true);
    setError(null);
    
    try {
      return await statisticsService.getOrderStatistics();
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to fetch order statistics';
      setError(errorMessage);
      throw err;
    } finally {
      setLoading(false);
    }
  }, []);

  const fetchProductStats = useCallback(async (): Promise<ProductStatistics> => {
    setLoading(true);
    setError(null);
    
    try {
      return await statisticsService.getProductStatistics();
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to fetch product statistics';
      setError(errorMessage);
      throw err;
    } finally {
      setLoading(false);
    }
  }, []);

  const fetchUserStats = useCallback(async (): Promise<UserStatistics> => {
    setLoading(true);
    setError(null);
    
    try {
      return await statisticsService.getUserStatistics();
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to fetch user statistics';
      setError(errorMessage);
      throw err;
    } finally {
      setLoading(false);
    }
  }, []);

  return {
    statistics,
    loading,
    error,
    fetchDashboardStats,
    fetchOrderStats,
    fetchProductStats,
    fetchUserStats,
    clearError: () => setError(null),
  };
}