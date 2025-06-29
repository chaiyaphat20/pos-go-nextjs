// Database entity types - mirrors backend models

export interface BaseEntity {
  id: string;
  created_at: string;
  updated_at: string;
  deleted_at?: string | null;
}

export interface User extends BaseEntity {
  username: string;
  email: string;
  role: UserRole;
  is_active: boolean;
  last_login?: string | null;
  email_verified?: boolean;
  profile?: UserProfile | null;
}

export interface UserProfile extends BaseEntity {
  user_id: string;
  first_name?: string;
  last_name?: string;
  phone?: string;
  address?: Address | null;
  preferences?: UserPreferences;
}

export interface Address {
  street: string;
  city: string;
  state: string;
  postal_code: string;
  country: string;
}

export interface UserPreferences {
  language: string;
  timezone: string;
  notifications: {
    email: boolean;
    sms: boolean;
    push: boolean;
  };
  theme: 'light' | 'dark' | 'auto';
}

export interface Product extends BaseEntity {
  name: string;
  description: string;
  price: number;
  stock: number;
  sku?: string;
  category_id?: string | null;
  category?: ProductCategory | null;
  images?: ProductImage[];
  attributes?: ProductAttribute[];
  is_active: boolean;
  weight?: number;
  dimensions?: ProductDimensions;
}

export interface ProductCategory extends BaseEntity {
  name: string;
  description?: string;
  parent_id?: string | null;
  parent?: ProductCategory | null;
  children?: ProductCategory[];
  is_active: boolean;
}

export interface ProductImage extends BaseEntity {
  product_id: string;
  url: string;
  alt_text?: string;
  sort_order: number;
  is_primary: boolean;
}

export interface ProductAttribute extends BaseEntity {
  product_id: string;
  name: string;
  value: string;
  type: 'text' | 'number' | 'boolean' | 'date';
}

export interface ProductDimensions {
  length: number;
  width: number;
  height: number;
  unit: 'cm' | 'inch';
}

export interface Order extends BaseEntity {
  order_number: string;
  customer_id?: string | null;
  customer?: User | null;
  status: OrderStatus;
  total_amount: number;
  subtotal: number;
  tax_amount: number;
  shipping_amount: number;
  discount_amount: number;
  currency: string;
  payment_status: PaymentStatus;
  payment_method?: string;
  shipping_address?: Address;
  billing_address?: Address;
  items: OrderItem[];
  notes?: string;
  tracking_number?: string;
  shipped_at?: string | null;
  delivered_at?: string | null;
}

export interface OrderItem extends BaseEntity {
  order_id: string;
  product_id: string;
  product?: Product;
  quantity: number;
  unit_price: number;
  total_price: number;
  product_snapshot?: ProductSnapshot;
}

export interface ProductSnapshot {
  name: string;
  description: string;
  sku?: string;
  price: number;
  image_url?: string;
}

// Enums
export type UserRole = 'admin' | 'user' | 'manager' | 'staff';

export type OrderStatus = 
  | 'pending' 
  | 'confirmed' 
  | 'processing' 
  | 'shipped' 
  | 'delivered' 
  | 'cancelled' 
  | 'refunded';

export type PaymentStatus = 
  | 'pending' 
  | 'paid' 
  | 'failed' 
  | 'refunded' 
  | 'partial';

// Legacy compatibility (can be removed later)
// @deprecated Use Order instead
export type Sale = Order;
export interface SaleItem extends OrderItem {
  sale_id: string;
  subtotal: number;
}