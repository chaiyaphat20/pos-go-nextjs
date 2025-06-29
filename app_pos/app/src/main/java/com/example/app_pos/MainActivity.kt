package com.example.app_pos

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Surface
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import com.example.app_pos.domain.usecase.CartUseCase
import com.example.app_pos.presentation.ui.*
import com.example.app_pos.ui.theme.PosAppTheme
import dagger.hilt.android.AndroidEntryPoint
import javax.inject.Inject

@AndroidEntryPoint
class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            PosAppTheme {
                Surface(
                    modifier = Modifier.fillMaxSize(),
                    color = MaterialTheme.colorScheme.background
                ) {
                    PosNavigation()
                }
            }
        }
    }
}

@Composable
fun PosNavigation() {
    val navController = rememberNavController()
    
    NavHost(
        navController = navController,
        startDestination = "login"
    ) {
        composable("login") {
            LoginScreen(
                onLoginSuccess = {
                    navController.navigate("dashboard") {
                        popUpTo("login") { inclusive = true }
                    }
                }
            )
        }
        
        composable("dashboard") {
            DashboardScreen(
                onNavigateToProducts = { navController.navigate("products") },
                onNavigateToCart = { navController.navigate("cart") },
                onNavigateToOrders = { navController.navigate("orders") },
                onNavigateToUsers = { navController.navigate("users") }
            )
        }
        
        composable("products") {
            ProductsScreen(
                onAddToCart = { product ->
                    // This will be handled by CartUseCase injection
                },
                onAddProduct = {
                    // Navigate to add product screen (can be implemented later)
                }
            )
        }
        
        composable("cart") {
            CartScreen(
                onCheckout = {
                    navController.navigate("orders")
                }
            )
        }
        
        composable("orders") {
            OrdersScreen()
        }
        
        composable("users") {
            UserListScreen()
        }
    }
}