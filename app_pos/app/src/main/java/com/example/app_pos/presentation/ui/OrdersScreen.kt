package com.example.app_pos.presentation.ui

import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import com.example.app_pos.domain.entity.Order
import com.example.app_pos.domain.entity.OrderItem
import com.example.app_pos.domain.entity.OrderStatus

@Composable
fun OrdersScreen() {
    // Mock data for demonstration
    val mockOrders = remember {
        listOf(
            Order(
                id = "ORD001",
                userId = "USER001",
                items = listOf(
                    OrderItem("1", "Coffee", 2, 45.0, 90.0),
                    OrderItem("2", "Sandwich", 1, 120.0, 120.0)
                ),
                totalAmount = 210.0,
                status = OrderStatus.COMPLETED,
                createdAt = "2024-01-15 10:30:00",
                updatedAt = "2024-01-15 10:45:00"
            ),
            Order(
                id = "ORD002",
                userId = "USER002",
                items = listOf(
                    OrderItem("3", "Smoothie", 1, 85.0, 85.0)
                ),
                totalAmount = 85.0,
                status = OrderStatus.PENDING,
                createdAt = "2024-01-15 11:00:00",
                updatedAt = "2024-01-15 11:00:00"
            )
        )
    }

    Column(
        modifier = Modifier
            .fillMaxSize()
            .padding(16.dp)
    ) {
        Text(
            text = "Orders",
            style = MaterialTheme.typography.headlineMedium,
            fontWeight = FontWeight.Bold
        )

        Spacer(modifier = Modifier.height(16.dp))

        if (mockOrders.isEmpty()) {
            Box(
                modifier = Modifier.fillMaxSize(),
                contentAlignment = Alignment.Center
            ) {
                Text(
                    text = "No orders found",
                    style = MaterialTheme.typography.titleMedium,
                    color = MaterialTheme.colorScheme.onSurfaceVariant
                )
            }
        } else {
            LazyColumn(
                verticalArrangement = Arrangement.spacedBy(12.dp)
            ) {
                items(mockOrders) { order ->
                    OrderCard(order = order)
                }
            }
        }
    }
}

@Composable
fun OrderCard(order: Order) {
    Card(
        modifier = Modifier.fillMaxWidth()
    ) {
        Column(
            modifier = Modifier.padding(16.dp)
        ) {
            Row(
                modifier = Modifier.fillMaxWidth(),
                horizontalArrangement = Arrangement.SpaceBetween,
                verticalAlignment = Alignment.Top
            ) {
                Column {
                    Text(
                        text = "Order #${order.id}",
                        style = MaterialTheme.typography.titleMedium,
                        fontWeight = FontWeight.Bold
                    )
                    
                    Text(
                        text = order.createdAt,
                        style = MaterialTheme.typography.bodySmall,
                        color = MaterialTheme.colorScheme.onSurfaceVariant
                    )
                }
                
                AssistChip(
                    onClick = { },
                    label = { Text(order.status.name) },
                    colors = AssistChipDefaults.assistChipColors(
                        containerColor = when (order.status) {
                            OrderStatus.COMPLETED -> MaterialTheme.colorScheme.primaryContainer
                            OrderStatus.PENDING -> MaterialTheme.colorScheme.secondaryContainer
                            OrderStatus.PROCESSING -> MaterialTheme.colorScheme.tertiaryContainer
                            OrderStatus.CANCELLED -> MaterialTheme.colorScheme.errorContainer
                            OrderStatus.CONFIRMED -> MaterialTheme.colorScheme.surfaceVariant
                        }
                    )
                )
            }
            
            Spacer(modifier = Modifier.height(12.dp))
            
            Text(
                text = "Items:",
                style = MaterialTheme.typography.titleSmall,
                fontWeight = FontWeight.Bold
            )
            
            Spacer(modifier = Modifier.height(4.dp))
            
            order.items.forEach { item ->
                Row(
                    modifier = Modifier.fillMaxWidth(),
                    horizontalArrangement = Arrangement.SpaceBetween
                ) {
                    Text(
                        text = "${item.productName} x${item.quantity}",
                        style = MaterialTheme.typography.bodyMedium
                    )
                    Text(
                        text = "฿${"%.2f".format(item.totalPrice)}",
                        style = MaterialTheme.typography.bodyMedium
                    )
                }
            }
            
            Spacer(modifier = Modifier.height(8.dp))
            
            Divider()
            
            Spacer(modifier = Modifier.height(8.dp))
            
            Row(
                modifier = Modifier.fillMaxWidth(),
                horizontalArrangement = Arrangement.SpaceBetween
            ) {
                Text(
                    text = "Total",
                    style = MaterialTheme.typography.titleMedium,
                    fontWeight = FontWeight.Bold
                )
                Text(
                    text = "฿${"%.2f".format(order.totalAmount)}",
                    style = MaterialTheme.typography.titleMedium,
                    fontWeight = FontWeight.Bold,
                    color = MaterialTheme.colorScheme.primary
                )
            }
        }
    }
}