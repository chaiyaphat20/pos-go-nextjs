package com.example.app_pos.domain.repository

import com.example.app_pos.domain.entity.User

interface UserRepository {
    suspend fun getUsers(): Result<List<User>>
    suspend fun getUserById(id: String): Result<User>
    suspend fun createUser(user: User): Result<User>
    suspend fun updateUser(user: User): Result<User>
    suspend fun deleteUser(id: String): Result<Boolean>
    suspend fun getUsersByRole(role: String): Result<List<User>>
}