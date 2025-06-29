import { httpClient } from '@/lib/httpClient';
import { TokenManager } from '@/lib/tokenManager';
import type { AuthResponse, LoginRequest, RegisterRequest } from '@/types/api/auth';

export class AuthService {
  async login(credentials: LoginRequest): Promise<AuthResponse> {
    const response = await httpClient.login(credentials.email, credentials.password);
    
    if (response.success && response.data) {
      TokenManager.setTokens(response.data);
      return response.data;
    }
    
    throw new Error(response.message || 'Login failed');
  }

  async register(userData: RegisterRequest): Promise<AuthResponse> {
    const response = await httpClient.register(userData);
    
    if (response.success && response.data) {
      TokenManager.setTokens(response.data);
      return response.data;
    }
    
    throw new Error(response.message || 'Registration failed');
  }

  async logout(): Promise<void> {
    try {
      await httpClient.post('/auth/logout');
    } catch (error) {
      // Continue with local logout even if server request fails
      console.warn('Server logout failed:', error);
    } finally {
      TokenManager.clearTokens();
    }
  }

  async refreshToken(): Promise<AuthResponse> {
    const refreshToken = TokenManager.getRefreshToken();
    
    if (!refreshToken) {
      throw new Error('No refresh token available');
    }

    const response = await httpClient.post<AuthResponse>('/auth/refresh', {
      refresh_token: refreshToken
    });

    if (response.success && response.data) {
      TokenManager.setTokens(response.data);
      return response.data;
    }

    throw new Error(response.message || 'Token refresh failed');
  }

  getCurrentUser() {
    return TokenManager.getUser();
  }

  isAuthenticated(): boolean {
    return TokenManager.hasValidTokens();
  }

  getAccessToken(): string | null {
    return TokenManager.getAccessToken();
  }
}