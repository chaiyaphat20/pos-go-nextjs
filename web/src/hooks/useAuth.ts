import { useState, useEffect } from 'react';
import { authService } from '@/services/api';
import type { User } from '@/types/database';
import type { LoginRequest, RegisterRequest, AuthResponse } from '@/types/api/auth';

export function useAuth() {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    checkAuthStatus();
  }, []);

  const checkAuthStatus = async () => {
    try {
      console.log('Checking auth status...');
      const isAuth = authService.isAuthenticated();
      console.log('Is authenticated:', isAuth);
      
      if (isAuth) {
        const currentUser = authService.getCurrentUser();
        console.log('Current user:', currentUser);
        setUser(currentUser);
      } else {
        console.log('Not authenticated, clearing user');
        setUser(null);
      }
    } catch (err) {
      console.error('Auth check failed:', err);
      authService.logout();
    } finally {
      setLoading(false);
    }
  };

  const login = async (credentials: LoginRequest): Promise<AuthResponse> => {
    setLoading(true);
    setError(null);
    
    try {
      const authResponse = await authService.login(credentials);
      setUser(authResponse.user);
      return authResponse;
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Login failed';
      setError(errorMessage);
      throw err;
    } finally {
      setLoading(false);
    }
  };

  const register = async (userData: RegisterRequest): Promise<AuthResponse> => {
    setLoading(true);
    setError(null);
    
    try {
      const authResponse = await authService.register(userData);
      setUser(authResponse.user);
      return authResponse;
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Registration failed';
      setError(errorMessage);
      throw err;
    } finally {
      setLoading(false);
    }
  };

  const logout = async (): Promise<void> => {
    setLoading(true);
    setError(null);
    
    try {
      await authService.logout();
      setUser(null);
    } catch (err) {
      console.error('Logout error:', err);
      // Still clear local state even if server logout fails
      setUser(null);
    } finally {
      setLoading(false);
    }
  };

  const refreshAuth = async (): Promise<void> => {
    try {
      const authResponse = await authService.refreshToken();
      setUser(authResponse.user);
    } catch (err) {
      console.error('Token refresh failed:', err);
      setUser(null);
      throw err;
    }
  };

  const isAuthenticated = !!user;
  console.log('useAuth return - user:', user);
  console.log('useAuth return - isAuthenticated:', isAuthenticated);
  
  return {
    user,
    loading,
    error,
    isAuthenticated,
    login,
    register,
    logout,
    refreshAuth,
    clearError: () => setError(null),
  };
}