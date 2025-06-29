package com.example.app_pos.data.remote.dto

import com.example.app_pos.domain.entity.LoginResponse
import com.google.gson.annotations.SerializedName

data class LoginResponseDto(
    @SerializedName("token")
    val token: String,
    @SerializedName("user")
    val user: UserDto,
    @SerializedName("expires_at")
    val expiresAt: String
)

fun LoginResponseDto.toDomain(): LoginResponse {
    return LoginResponse(
        token = token,
        user = user.toDomain(),
        expiresAt = expiresAt
    )
}