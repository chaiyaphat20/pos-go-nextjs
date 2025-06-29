import type { User, Product, Order, OrderStatus } from '../database';
import type { UserStatistics, ProductStatistics, OrderStatistics, DashboardStatistics } from '../statistics';
import type { 
  CreateUserRequest, 
  UpdateUserRequest, 
  UserQuery,
  CreateProductRequest,
  UpdateProductRequest,
  ProductQuery,
  UpdateStockRequest,
  CreateOrderRequest,
  UpdateOrderRequest,
  OrderQuery,
  LoginRequest,
  RegisterRequest,
  AuthResponse
} from './index';

// Auth Hook
export interface UseAuthReturn {
  user: User | null;
  loading: boolean;
  error: string | null;
  isAuthenticated: boolean;
  login: (credentials: LoginRequest) => Promise<AuthResponse>;
  register: (userData: RegisterRequest) => Promise<AuthResponse>;
  logout: () => Promise<void>;
  refreshAuth: () => Promise<void>;
  clearError: () => void;
}

// Users Hook
export interface UseUsersReturn {
  users: User[];
  loading: boolean;
  error: string | null;
  fetchUsers: (params?: UserQuery) => Promise<void>;
  createUser: (userData: CreateUserRequest) => Promise<User>;
  updateUser: (id: string, data: UpdateUserRequest) => Promise<User>;
  deleteUser: (id: string) => Promise<void>;
  getUserById: (id: string) => Promise<User>;
  clearError: () => void;
}

// Products Hook
export interface UseProductsReturn {
  products: Product[];
  loading: boolean;
  error: string | null;
  fetchProducts: (params?: ProductQuery) => Promise<void>;
  createProduct: (productData: CreateProductRequest) => Promise<Product>;
  updateProduct: (id: string, data: UpdateProductRequest) => Promise<Product>;
  deleteProduct: (id: string) => Promise<boolean>;
  updateStock: (id: string, data: UpdateStockRequest) => Promise<Product>;
  getProductById: (id: string) => Promise<Product>;
  searchProducts: (query: string, limit?: number) => Promise<Product[]>;
  getLowStockProducts: (threshold?: number) => Promise<Product[]>;
  clearError: () => void;
}

// Orders Hook
export interface UseOrdersReturn {
  orders: Order[];
  loading: boolean;
  error: string | null;
  fetchOrders: (params?: OrderQuery) => Promise<void>;
  createOrder: (orderData: CreateOrderRequest) => Promise<Order>;
  updateOrder: (id: string, data: UpdateOrderRequest) => Promise<Order>;
  updateOrderStatus: (id: string, status: OrderStatus) => Promise<Order>;
  cancelOrder: (id: string) => Promise<boolean>;
  getOrderById: (id: string) => Promise<Order>;
  getMyOrders: () => Promise<Order[]>;
  getOrdersByUser: (userId: string) => Promise<Order[]>;
  getOrderStatistics: () => Promise<OrderStatistics>;
  clearError: () => void;
}

// Statistics Hook
export interface UseStatisticsReturn {
  statistics: DashboardStatistics | null;
  loading: boolean;
  error: string | null;
  fetchDashboardStats: () => Promise<void>;
  fetchOrderStats: () => Promise<OrderStatistics>;
  fetchProductStats: () => Promise<ProductStatistics>;
  fetchUserStats: () => Promise<UserStatistics>;
  clearError: () => void;
}