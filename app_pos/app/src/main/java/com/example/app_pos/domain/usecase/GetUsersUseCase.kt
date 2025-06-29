package com.example.app_pos.domain.usecase

import com.example.app_pos.domain.entity.User
import com.example.app_pos.domain.repository.UserRepository
import javax.inject.Inject

class GetUsersUseCase @Inject constructor(
    private val repository: UserRepository
) {
    suspend operator fun invoke(): Result<List<User>> {
        return repository.getUsers()
    }
}