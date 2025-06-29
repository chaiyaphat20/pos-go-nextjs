// Statistics and analytics types

export interface OrderStatistics {
  totalOrders: number;
  totalRevenue: number;
  averageOrderValue: number;
  ordersByStatus: {
    pending: number;
    confirmed: number;
    shipped: number;
    delivered: number;
    cancelled: number;
  };
  recentOrders: {
    today: number;
    thisWeek: number;
    thisMonth: number;
    thisYear: number;
  };
  topProducts: Array<{
    productId: string;
    productName: string;
    quantity: number;
    revenue: number;
  }>;
  topCustomers: Array<{
    customerId: string;
    customerName: string;
    totalOrders: number;
    totalSpent: number;
  }>;
  salesTrend: Array<{
    date: string;
    orders: number;
    revenue: number;
  }>;
}

export interface ProductStatistics {
  totalProducts: number;
  lowStockCount: number;
  outOfStockCount: number;
  categoryBreakdown: Array<{
    category: string;
    count: number;
    totalValue: number;
  }>;
  topSellingProducts: Array<{
    productId: string;
    name: string;
    quantitySold: number;
    revenue: number;
  }>;
  inventoryValue: {
    total: number;
    lowStock: number;
    available: number;
  };
}

export interface UserStatistics {
  totalUsers: number;
  activeUsers: number;
  newUsersThisMonth: number;
  usersByRole: {
    admin: number;
    user: number;
  };
  userActivity: Array<{
    date: string;
    activeUsers: number;
    newUsers: number;
  }>;
}

export interface DashboardStatistics {
  orders: OrderStatistics;
  products: ProductStatistics;
  users: UserStatistics;
  revenue: {
    total: number;
    thisMonth: number;
    lastMonth: number;
    growthPercentage: number;
  };
  alerts: Array<{
    id: string;
    type: 'info' | 'warning' | 'error' | 'success';
    title: string;
    message: string;
    timestamp: string;
  }>;
}