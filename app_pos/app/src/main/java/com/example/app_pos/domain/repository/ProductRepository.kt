package com.example.app_pos.domain.repository

import com.example.app_pos.domain.entity.Product

interface ProductRepository {
    suspend fun getProducts(): Result<List<Product>>
    suspend fun getProductById(id: String): Result<Product>
    suspend fun createProduct(product: Product): Result<Product>
    suspend fun updateProduct(product: Product): Result<Product>
    suspend fun deleteProduct(id: String): Result<Boolean>
    suspend fun getProductsByCategory(category: String): Result<List<Product>>
    suspend fun searchProducts(query: String): Result<List<Product>>
}