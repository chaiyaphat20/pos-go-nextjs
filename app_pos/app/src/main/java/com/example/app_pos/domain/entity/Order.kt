package com.example.app_pos.domain.entity

data class Order(
    val id: String,
    val userId: String,
    val items: List<OrderItem>,
    val totalAmount: Double,
    val status: OrderStatus,
    val createdAt: String,
    val updatedAt: String
)

data class OrderItem(
    val productId: String,
    val productName: String,
    val quantity: Int,
    val unitPrice: Double,
    val totalPrice: Double
)

enum class OrderStatus {
    PENDING,
    CONFIRMED,
    PROCESSING,
    COMPLETED,
    CANCELLED
}