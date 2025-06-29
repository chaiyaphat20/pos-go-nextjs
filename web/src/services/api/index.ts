// Service classes
export { AuthService } from './authService';
export { UserService } from './userService';
export { ProductService } from './productService';
export { OrderService } from './orderService';
export { StatisticsService } from './statisticsService';

// Create singleton instances
import { AuthService } from './authService';
import { UserService } from './userService';
import { ProductService } from './productService';
import { OrderService } from './orderService';
import { StatisticsService } from './statisticsService';

export const authService = new AuthService();
export const userService = new UserService();
export const productService = new ProductService();
export const orderService = new OrderService();
export const statisticsService = new StatisticsService();

// Export types for convenience
export type { ApiResponse } from '@/types/api/response';