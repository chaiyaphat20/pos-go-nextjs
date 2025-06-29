import axios, { AxiosInstance, AxiosResponse, InternalAxiosRequestConfig } from 'axios';
import { TokenManager } from './tokenManager';
import type { ApiResponse, AuthResponse, RegisterRequest } from '@/types/api';

class HttpClient {
  private axiosInstance: AxiosInstance;
  private isRefreshing = false;
  private refreshSubscribers: ((token: string) => void)[] = [];

  constructor() {
    this.axiosInstance = axios.create({
      baseURL: process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api/v1',
      timeout: 10000,
      headers: {
        'Content-Type': 'application/json',
      },
    });

    this.setupInterceptors();
  }

  private setupInterceptors(): void {
    // Request interceptor
    this.axiosInstance.interceptors.request.use(
      (config: InternalAxiosRequestConfig) => {
        const token = TokenManager.getAccessToken();
        if (token && !TokenManager.isTokenExpired(token)) {
          config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
      },
      (error) => Promise.reject(error)
    );

    // Response interceptor
    this.axiosInstance.interceptors.response.use(
      (response: AxiosResponse) => response,
      async (error) => {
        const originalRequest = error.config;

        // If error is 401 and we haven't already tried to refresh
        if (error.response?.status === 401 && !originalRequest._retry) {
          if (this.isRefreshing) {
            // If already refreshing, wait for the new token
            return new Promise((resolve) => {
              this.refreshSubscribers.push((token: string) => {
                originalRequest.headers.Authorization = `Bearer ${token}`;
                resolve(this.axiosInstance(originalRequest));
              });
            });
          }

          originalRequest._retry = true;
          this.isRefreshing = true;

          try {
            const newToken = await this.refreshAccessToken();
            this.isRefreshing = false;
            this.onTokenRefreshed(newToken);
            
            // Retry original request with new token
            originalRequest.headers.Authorization = `Bearer ${newToken}`;
            return this.axiosInstance(originalRequest);
          } catch (refreshError) {
            this.isRefreshing = false;
            this.onTokenRefreshFailed();
            return Promise.reject(refreshError);
          }
        }

        return Promise.reject(error);
      }
    );
  }

  private async refreshAccessToken(): Promise<string> {
    const refreshToken = TokenManager.getRefreshToken();
    
    if (!refreshToken || TokenManager.isTokenExpired(refreshToken)) {
      throw new Error('No valid refresh token');
    }

    try {
      const response = await axios.post(
        `${process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api/v1'}/auth/refresh`,
        { refresh_token: refreshToken },
        { headers: { 'Content-Type': 'application/json' } }
      );

      const authResponse = response.data.data;
      TokenManager.setTokens(authResponse);
      return authResponse.access_token;
    } catch (error) {
      TokenManager.clearTokens();
      // Redirect to login page
      if (typeof window !== 'undefined') {
        window.location.href = '/auth/signin';
      }
      throw error;
    }
  }

  private onTokenRefreshed(token: string): void {
    this.refreshSubscribers.forEach((callback) => callback(token));
    this.refreshSubscribers = [];
  }

  private onTokenRefreshFailed(): void {
    this.refreshSubscribers = [];
    TokenManager.clearTokens();
    if (typeof window !== 'undefined') {
      window.location.href = '/auth/signin';
    }
  }

  // Public methods
  async get<T = unknown>(url: string, params?: Record<string, unknown>): Promise<ApiResponse<T>> {
    const response = await this.axiosInstance.get(url, { params });
    return response.data;
  }

  async post<T = unknown>(url: string, data?: unknown): Promise<ApiResponse<T>> {
    const response = await this.axiosInstance.post(url, data);
    return response.data;
  }

  async put<T = unknown>(url: string, data?: unknown): Promise<ApiResponse<T>> {
    const response = await this.axiosInstance.put(url, data);
    return response.data;
  }

  async patch<T = unknown>(url: string, data?: unknown): Promise<ApiResponse<T>> {
    const response = await this.axiosInstance.patch(url, data);
    return response.data;
  }

  async delete<T = unknown>(url: string): Promise<ApiResponse<T>> {
    const response = await this.axiosInstance.delete(url);
    return response.data;
  }

  // Auth methods (no token required)
  async login(email: string, password: string): Promise<ApiResponse<AuthResponse>> {
    const response = await axios.post(
      `${this.axiosInstance.defaults.baseURL}/auth/login`,
      { email, password }
    );
    return response.data;
  }

  async register(userData: RegisterRequest): Promise<ApiResponse<AuthResponse>> {
    const response = await axios.post(
      `${this.axiosInstance.defaults.baseURL}/auth/register`,
      userData
    );
    return response.data;
  }
}

// Export singleton instance
export const httpClient = new HttpClient();