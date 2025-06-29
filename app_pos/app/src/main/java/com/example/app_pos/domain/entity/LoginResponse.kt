package com.example.app_pos.domain.entity

data class LoginResponse(
    val token: String,
    val user: User,
    val expiresAt: String
)