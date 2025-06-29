import type { UserRole } from '../database';
import type { QueryParams } from '../common';

// User management interfaces
export interface CreateUserRequest {
  username: string;
  email: string;
  password: string;
  password_confirmation: string;
  role: UserRole;
  is_active?: boolean;
  profile?: {
    first_name?: string;
    last_name?: string;
    phone?: string;
  };
}

export interface UpdateUserRequest {
  username?: string;
  email?: string;
  password?: string;
  password_confirmation?: string;
  role?: UserRole;
  is_active?: boolean;
  profile?: {
    first_name?: string;
    last_name?: string;
    phone?: string;
  };
}

export interface UserQuery extends QueryParams {
  role?: UserRole;
  is_active?: boolean;
  email_verified?: boolean;
  [key: string]: unknown;
}