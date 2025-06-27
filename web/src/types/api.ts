export interface User {
  id: number;
  name: string;
  email: string;
  role: string;
  created_at: string;
  updated_at: string;
}

export interface Product {
  id: number;
  name: string;
  description: string;
  price: number;
  stock: number;
  created_at: string;
  updated_at: string;
}

export interface SaleItem {
  id: number;
  sale_id: number;
  product_id: number;
  product: Product;
  quantity: number;
  price: number;
  subtotal: number;
}

export interface Sale {
  id: number;
  total: number;
  customer_id?: number;
  customer?: User;
  sale_items: SaleItem[];
  created_at: string;
  updated_at: string;
}

export interface CreateUserRequest {
  name: string;
  email: string;
  password: string;
  role: string;
}

export interface UpdateUserRequest {
  name: string;
  email: string;
  password?: string;
  role: string;
}

export interface CreateProductRequest {
  name: string;
  description: string;
  price: number;
  stock: number;
}

export interface UpdateProductRequest {
  name: string;
  description: string;
  price: number;
  stock: number;
}

export interface UpdateStockRequest {
  stock: number;
}

export interface CreateSaleRequest {
  customer_id?: number;
  items: {
    product_id: number;
    quantity: number;
  }[];
}

export interface ApiResponse<T> {
  data: T;
  message?: string;
}

export interface ApiError {
  error: string;
}