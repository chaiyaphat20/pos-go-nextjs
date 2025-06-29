package com.example.app_pos.domain.entity

data class Product(
    val id: String,
    val name: String,
    val description: String,
    val price: Double,
    val stockQuantity: Int,
    val category: String,
    val imageUrl: String? = null,
    val isActive: Boolean = true,
    val createdAt: String,
    val updatedAt: String
)