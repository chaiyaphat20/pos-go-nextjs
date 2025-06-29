package com.example.app_pos.data.remote.dto

import com.example.app_pos.domain.entity.Order
import com.example.app_pos.domain.entity.OrderItem
import com.example.app_pos.domain.entity.OrderStatus
import com.google.gson.annotations.SerializedName

data class OrderDto(
    @SerializedName("id")
    val id: String,
    @SerializedName("user_id")
    val userId: String,
    @SerializedName("items")
    val items: List<OrderItemDto>,
    @SerializedName("total_amount")
    val totalAmount: Double,
    @SerializedName("status")
    val status: String,
    @SerializedName("created_at")
    val createdAt: String,
    @SerializedName("updated_at")
    val updatedAt: String
)

data class OrderItemDto(
    @SerializedName("product_id")
    val productId: String,
    @SerializedName("product_name")
    val productName: String,
    @SerializedName("quantity")
    val quantity: Int,
    @SerializedName("unit_price")
    val unitPrice: Double,
    @SerializedName("total_price")
    val totalPrice: Double
)

fun OrderDto.toDomain(): Order {
    return Order(
        id = id,
        userId = userId,
        items = items.map { it.toDomain() },
        totalAmount = totalAmount,
        status = when (status.uppercase()) {
            "PENDING" -> OrderStatus.PENDING
            "CONFIRMED" -> OrderStatus.CONFIRMED
            "PROCESSING" -> OrderStatus.PROCESSING
            "COMPLETED" -> OrderStatus.COMPLETED
            "CANCELLED" -> OrderStatus.CANCELLED
            else -> OrderStatus.PENDING
        },
        createdAt = createdAt,
        updatedAt = updatedAt
    )
}

fun OrderItemDto.toDomain(): OrderItem {
    return OrderItem(
        productId = productId,
        productName = productName,
        quantity = quantity,
        unitPrice = unitPrice,
        totalPrice = totalPrice
    )
}