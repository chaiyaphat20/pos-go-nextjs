package com.example.app_pos.domain.repository

import com.example.app_pos.domain.entity.Order

interface OrderRepository {
    suspend fun getOrders(): Result<List<Order>>
    suspend fun getOrderById(id: String): Result<Order>
    suspend fun createOrder(order: Order): Result<Order>
    suspend fun updateOrderStatus(orderId: String, status: String): Result<Order>
    suspend fun getUserOrders(userId: String): Result<List<Order>>
}