package com.example.app_pos.data.remote.dto

import com.example.app_pos.domain.entity.Product
import com.google.gson.annotations.SerializedName

data class ProductDto(
    @SerializedName("id")
    val id: String,
    @SerializedName("name")
    val name: String,
    @SerializedName("description")
    val description: String,
    @SerializedName("price")
    val price: Double,
    @SerializedName("stock_quantity")
    val stockQuantity: Int,
    @SerializedName("category")
    val category: String,
    @SerializedName("image_url")
    val imageUrl: String? = null,
    @SerializedName("is_active")
    val isActive: Boolean = true,
    @SerializedName("created_at")
    val createdAt: String,
    @SerializedName("updated_at")
    val updatedAt: String
)

fun ProductDto.toDomain(): Product {
    return Product(
        id = id,
        name = name,
        description = description,
        price = price,
        stockQuantity = stockQuantity,
        category = category,
        imageUrl = imageUrl,
        isActive = isActive,
        createdAt = createdAt,
        updatedAt = updatedAt
    )
}

fun Product.toDto(): ProductDto {
    return ProductDto(
        id = id,
        name = name,
        description = description,
        price = price,
        stockQuantity = stockQuantity,
        category = category,
        imageUrl = imageUrl,
        isActive = isActive,
        createdAt = createdAt,
        updatedAt = updatedAt
    )
}