import type { QueryParams } from '../common';

// Product management interfaces
export interface CreateProductRequest {
  name: string;
  description: string;
  price: number;
  stock: number;
  category_id?: string;
  sku?: string;
  is_active?: boolean;
  images?: string[];
  attributes?: Record<string, unknown>;
  dimensions?: {
    length: number;
    width: number;
    height: number;
    unit: 'cm' | 'inch';
  };
}

export interface UpdateProductRequest {
  name?: string;
  description?: string;
  price?: number;
  stock?: number;
  category_id?: string;
  sku?: string;
  is_active?: boolean;
  images?: string[];
  attributes?: Record<string, unknown>;
  dimensions?: {
    length: number;
    width: number;
    height: number;
    unit: 'cm' | 'inch';
  };
}

export interface ProductQuery extends QueryParams {
  category_id?: string;
  is_active?: boolean;
  min_price?: number;
  max_price?: number;
  low_stock?: boolean;
  out_of_stock?: boolean;
  [key: string]: unknown;
}

export interface UpdateStockRequest {
  stock: number;
  reason?: string;
}