package com.example.app_pos.data.repository

import com.example.app_pos.data.local.TokenManager
import com.example.app_pos.data.remote.ApiService
import com.example.app_pos.data.remote.dto.toDomain
import com.example.app_pos.data.remote.dto.toDto
import com.example.app_pos.domain.entity.LoginRequest
import com.example.app_pos.domain.entity.LoginResponse
import com.example.app_pos.domain.repository.AuthRepository
import javax.inject.Inject

class AuthRepositoryImpl @Inject constructor(
    private val apiService: ApiService,
    private val tokenManager: TokenManager
) : AuthRepository {

    override suspend fun login(request: LoginRequest): Result<LoginResponse> {
        return try {
            val response = apiService.login(request.toDto())
            if (response.isSuccessful) {
                response.body()?.let { loginResponseDto ->
                    val loginResponse = loginResponseDto.toDomain()
                    tokenManager.saveToken(loginResponse.token)
                    Result.success(loginResponse)
                } ?: Result.failure(Exception("Login failed: Empty response"))
            } else {
                Result.failure(Exception("Login failed: ${response.message()}"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    override suspend fun logout(): Result<Boolean> {
        return try {
            val response = apiService.logout()
            tokenManager.clearToken()
            if (response.isSuccessful) {
                Result.success(true)
            } else {
                Result.success(true) // Clear token even if API call fails
            }
        } catch (e: Exception) {
            tokenManager.clearToken() // Clear token even if API call fails
            Result.success(true)
        }
    }

    override suspend fun isLoggedIn(): Boolean {
        return tokenManager.isLoggedIn()
    }

    override suspend fun getToken(): String? {
        return tokenManager.getToken()
    }

    override suspend fun saveToken(token: String) {
        tokenManager.saveToken(token)
    }

    override suspend fun clearToken() {
        tokenManager.clearToken()
    }
}