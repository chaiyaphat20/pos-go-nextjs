import type { ApiError as CommonApiError } from '../common';

// Standard API Response interface
export interface ApiResponse<T = unknown> {
  success: boolean;
  message: string;
  data?: T;
  error?: string | CommonApiError;
  timestamp?: string;
}

// HTTP Error Response
export interface ErrorResponse {
  success: false;
  message: string;
  error: string | CommonApiError;
  timestamp: string;
}

// Success Response
export interface SuccessResponse<T> {
  success: true;
  message: string;
  data: T;
  timestamp?: string;
}