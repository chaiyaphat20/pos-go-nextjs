package com.example.app_pos.domain.usecase

import com.example.app_pos.domain.repository.AuthRepository
import javax.inject.Inject

class LogoutUseCase @Inject constructor(
    private val authRepository: AuthRepository
) {
    suspend operator fun invoke(): Result<Boolean> {
        return authRepository.logout()
    }
}