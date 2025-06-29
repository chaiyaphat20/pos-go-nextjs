package com.example.app_pos.domain.entity

data class User(
    val id: String,
    val username: String,
    val email: String,
    val role: UserRole,
    val isActive: Boolean,
    val createdAt: String
)

enum class UserRole {
    ADMIN,
    USER,
    MODERATOR
}