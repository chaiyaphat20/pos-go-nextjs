import type { OrderStatus } from '../database';
import type { QueryParams } from '../common';

// Order management interfaces
export interface CreateOrderRequest {
  customer_id?: string;
  items: OrderItemRequest[];
  shipping_address?: AddressRequest;
  billing_address?: AddressRequest;
  payment_method?: string;
  notes?: string;
}

export interface OrderItemRequest {
  product_id: string;
  quantity: number;
  unit_price?: number;
}

export interface AddressRequest {
  street: string;
  city: string;
  state: string;
  postal_code: string;
  country: string;
}

export interface UpdateOrderRequest {
  customer_id?: string;
  status?: OrderStatus;
  shipping_address?: AddressRequest;
  billing_address?: AddressRequest;
  payment_method?: string;
  notes?: string;
  tracking_number?: string;
}

export interface OrderQuery extends QueryParams {
  customer_id?: string;
  status?: OrderStatus;
  payment_status?: string;
  date_from?: string;
  date_to?: string;
  [key: string]: unknown;
}