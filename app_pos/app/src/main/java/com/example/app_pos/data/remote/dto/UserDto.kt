package com.example.app_pos.data.remote.dto

import com.example.app_pos.domain.entity.User
import com.example.app_pos.domain.entity.UserRole
import com.google.gson.annotations.SerializedName

data class UserDto(
    @SerializedName("id")
    val id: String,
    @SerializedName("username")
    val username: String,
    @SerializedName("email")
    val email: String,
    @SerializedName("role")
    val role: String,
    @SerializedName("is_active")
    val isActive: Boolean,
    @SerializedName("created_at")
    val createdAt: String
)

fun UserDto.toDomain(): User {
    return User(
        id = id,
        username = username,
        email = email,
        role = when (role.uppercase()) {
            "ADMIN" -> UserRole.ADMIN
            "MODERATOR" -> UserRole.MODERATOR
            else -> UserRole.USER
        },
        isActive = isActive,
        createdAt = createdAt
    )
}

fun User.toDto(): UserDto {
    return UserDto(
        id = id,
        username = username,
        email = email,
        role = role.name.lowercase(),
        isActive = isActive,
        createdAt = createdAt
    )
}