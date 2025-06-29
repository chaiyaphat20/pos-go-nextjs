package com.example.app_pos.domain.repository

import com.example.app_pos.domain.entity.LoginRequest
import com.example.app_pos.domain.entity.LoginResponse

interface AuthRepository {
    suspend fun login(request: LoginRequest): Result<LoginResponse>
    suspend fun logout(): Result<Boolean>
    suspend fun isLoggedIn(): Boolean
    suspend fun getToken(): String?
    suspend fun saveToken(token: String)
    suspend fun clearToken()
}