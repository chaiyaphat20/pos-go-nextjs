package com.example.app_pos.domain.usecase

import com.example.app_pos.domain.entity.LoginRequest
import com.example.app_pos.domain.entity.LoginResponse
import com.example.app_pos.domain.repository.AuthRepository
import javax.inject.Inject

class LoginUseCase @Inject constructor(
    private val authRepository: AuthRepository
) {
    suspend operator fun invoke(username: String, password: String): Result<LoginResponse> {
        if (username.isBlank()) {
            return Result.failure(Exception("Username cannot be empty"))
        }
        if (password.isBlank()) {
            return Result.failure(Exception("Password cannot be empty"))
        }
        
        val request = LoginRequest(username = username.trim(), password = password)
        return authRepository.login(request)
    }
}