package com.example.app_pos.data.remote

import com.example.app_pos.data.remote.dto.LoginRequestDto
import com.example.app_pos.data.remote.dto.LoginResponseDto
import com.example.app_pos.data.remote.dto.OrderDto
import com.example.app_pos.data.remote.dto.ProductDto
import com.example.app_pos.data.remote.dto.UserDto
import retrofit2.Response
import retrofit2.http.*

interface ApiService {
    @POST("auth/login")
    suspend fun login(@Body request: LoginRequestDto): Response<LoginResponseDto>

    @POST("auth/logout")
    suspend fun logout(): Response<Unit>

    // Products
    @GET("products")
    suspend fun getProducts(): Response<List<ProductDto>>

    @GET("products/{id}")
    suspend fun getProductById(@Path("id") id: String): Response<ProductDto>

    @POST("products")
    suspend fun createProduct(@Body product: ProductDto): Response<ProductDto>

    @PUT("products/{id}")
    suspend fun updateProduct(
        @Path("id") id: String,
        @Body product: ProductDto
    ): Response<ProductDto>

    @DELETE("products/{id}")
    suspend fun deleteProduct(@Path("id") id: String): Response<Unit>

    @GET("products/category/{category}")
    suspend fun getProductsByCategory(@Path("category") category: String): Response<List<ProductDto>>

    @GET("products/search")
    suspend fun searchProducts(@Query("q") query: String): Response<List<ProductDto>>

    // Orders
    @GET("orders")
    suspend fun getOrders(): Response<List<OrderDto>>

    @GET("orders/{id}")
    suspend fun getOrderById(@Path("id") id: String): Response<OrderDto>

    @POST("orders")
    suspend fun createOrder(@Body order: OrderDto): Response<OrderDto>

    @PUT("orders/{id}/status")
    suspend fun updateOrderStatus(
        @Path("id") orderId: String,
        @Body status: Map<String, String>
    ): Response<OrderDto>

    @GET("orders/user/{userId}")
    suspend fun getUserOrders(@Path("userId") userId: String): Response<List<OrderDto>>
    @GET("users")
    suspend fun getUsers(): Response<List<UserDto>>

    @GET("users/{id}")
    suspend fun getUserById(@Path("id") id: String): Response<UserDto>

    @POST("users")
    suspend fun createUser(@Body user: UserDto): Response<UserDto>

    @PUT("users/{id}")
    suspend fun updateUser(
        @Path("id") id: String,
        @Body user: UserDto
    ): Response<UserDto>

    @DELETE("users/{id}")
    suspend fun deleteUser(@Path("id") id: String): Response<Unit>

    @GET("users/role/{role}")
    suspend fun getUsersByRole(@Path("role") role: String): Response<List<UserDto>>
}