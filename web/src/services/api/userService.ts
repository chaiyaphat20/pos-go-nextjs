import { httpClient } from '@/lib/httpClient';
import type { User } from '@/types/database';
import type { CreateUserRequest, UpdateUserRequest, UserQuery } from '@/types/api/user';

export class UserService {
  async getUsers(params: UserQuery = {}): Promise<User[]> {
    const response = await httpClient.get<User[]>('/users', params);
    
    if (response.success) {
      return response.data || [];
    }
    
    throw new Error(response.message || 'Failed to fetch users');
  }

  async getUserById(id: string): Promise<User> {
    const response = await httpClient.get<User>(`/users/${id}`);
    
    if (response.success && response.data) {
      return response.data;
    }
    
    throw new Error(response.message || 'Failed to fetch user');
  }

  async createUser(userData: CreateUserRequest): Promise<User> {
    const response = await httpClient.post<User>('/users', userData);
    
    if (response.success && response.data) {
      return response.data;
    }
    
    throw new Error(response.message || 'Failed to create user');
  }

  async updateUser(id: string, userData: UpdateUserRequest): Promise<User> {
    const response = await httpClient.put<User>(`/users/${id}`, userData);
    
    if (response.success && response.data) {
      return response.data;
    }
    
    throw new Error(response.message || 'Failed to update user');
  }

  async deleteUser(id: string): Promise<void> {
    const response = await httpClient.delete(`/users/${id}`);
    
    if (!response.success) {
      throw new Error(response.message || 'Failed to delete user');
    }
  }

  async getCurrentUser(): Promise<User> {
    const response = await httpClient.get<User>('/users/me');
    
    if (response.success && response.data) {
      return response.data;
    }
    
    throw new Error(response.message || 'Failed to fetch current user');
  }
}