package com.example.app_pos.data.repository

import com.example.app_pos.data.remote.ApiService
import com.example.app_pos.data.remote.dto.toDomain
import com.example.app_pos.data.remote.dto.toDto
import com.example.app_pos.domain.entity.User
import com.example.app_pos.domain.repository.UserRepository
import javax.inject.Inject

class UserRepositoryImpl @Inject constructor(
    private val apiService: ApiService
) : UserRepository {

    override suspend fun getUsers(): Result<List<User>> {
        return try {
            val response = apiService.getUsers()
            if (response.isSuccessful) {
                val users = response.body()?.map { it.toDomain() } ?: emptyList()
                Result.success(users)
            } else {
                Result.failure(Exception("Failed to fetch users: ${response.message()}"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    override suspend fun getUserById(id: String): Result<User> {
        return try {
            val response = apiService.getUserById(id)
            if (response.isSuccessful) {
                response.body()?.let { userDto ->
                    Result.success(userDto.toDomain())
                } ?: Result.failure(Exception("User not found"))
            } else {
                Result.failure(Exception("Failed to fetch user: ${response.message()}"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    override suspend fun createUser(user: User): Result<User> {
        return try {
            val response = apiService.createUser(user.toDto())
            if (response.isSuccessful) {
                response.body()?.let { userDto ->
                    Result.success(userDto.toDomain())
                } ?: Result.failure(Exception("Failed to create user"))
            } else {
                Result.failure(Exception("Failed to create user: ${response.message()}"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    override suspend fun updateUser(user: User): Result<User> {
        return try {
            val response = apiService.updateUser(user.id, user.toDto())
            if (response.isSuccessful) {
                response.body()?.let { userDto ->
                    Result.success(userDto.toDomain())
                } ?: Result.failure(Exception("Failed to update user"))
            } else {
                Result.failure(Exception("Failed to update user: ${response.message()}"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    override suspend fun deleteUser(id: String): Result<Boolean> {
        return try {
            val response = apiService.deleteUser(id)
            if (response.isSuccessful) {
                Result.success(true)
            } else {
                Result.failure(Exception("Failed to delete user: ${response.message()}"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    override suspend fun getUsersByRole(role: String): Result<List<User>> {
        return try {
            val response = apiService.getUsersByRole(role)
            if (response.isSuccessful) {
                val users = response.body()?.map { it.toDomain() } ?: emptyList()
                Result.success(users)
            } else {
                Result.failure(Exception("Failed to fetch users by role: ${response.message()}"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }
}