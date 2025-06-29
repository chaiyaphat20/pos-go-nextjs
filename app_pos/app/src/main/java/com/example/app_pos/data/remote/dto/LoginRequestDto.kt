package com.example.app_pos.data.remote.dto

import com.example.app_pos.domain.entity.LoginRequest
import com.google.gson.annotations.SerializedName

data class LoginRequestDto(
    @SerializedName("username")
    val username: String,
    @SerializedName("password")
    val password: String
)

fun LoginRequest.toDto(): LoginRequestDto {
    return LoginRequestDto(
        username = username,
        password = password
    )
}