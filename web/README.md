# Go Clean POS - Frontend

A modern Point of Sale (POS) system frontend built with Next.js, TypeScript, and Tailwind CSS.

## Features

- **Authentication**: NextAuth.js with credentials provider
- **State Management**: Redux Toolkit for application state
- **Internationalization**: Multi-language support (English & Thai)
- **UI Components**: Modern React components with Tailwind CSS
- **API Integration**: Complete integration with Go Clean POS backend
- **Responsive Design**: Mobile-first responsive design

## Tech Stack

- **Framework**: Next.js 15 with App Router
- **Language**: TypeScript
- **Styling**: Tailwind CSS
- **State Management**: Redux Toolkit
- **Authentication**: NextAuth.js
- **Icons**: Heroicons
- **UI Components**: Headless UI
- **HTTP Client**: Axios

## Project Structure

```
src/
├── app/                    # Next.js App Router pages
│   ├── auth/              # Authentication pages
│   ├── dashboard/         # Dashboard page
│   ├── products/          # Products management
│   ├── sales/             # Sales management
│   ├── users/             # Users management
│   └── api/               # API routes
├── components/            # React components
│   ├── forms/             # Form components
│   ├── layout/            # Layout components
│   ├── providers/         # Context providers
│   └── ui/                # UI components
├── lib/                   # Utility libraries
├── services/              # API service functions
├── store/                 # Redux store and slices
├── types/                 # TypeScript type definitions
└── utils/                 # Utility functions
```

## Getting Started

### Prerequisites

- Node.js 18+ 
- npm or yarn
- Go Clean POS Backend running on `http://localhost:8080`

### Installation

1. **Install dependencies:**
   ```bash
   npm install
   ```

2. **Set up environment variables:**
   Update `.env.local` if needed:
   ```bash
   NEXT_PUBLIC_API_URL=http://localhost:8080
   NEXTAUTH_URL=http://localhost:3000
   NEXTAUTH_SECRET=your-nextauth-secret-key-here
   ```

3. **Start the development server:**
   ```bash
   npm run dev
   ```

4. **Open your browser:**
   Navigate to `http://localhost:3000`

### Default Login Credentials

For demo purposes, use these credentials:
- **Email**: `admin@example.com`
- **Password**: `admin`

## Available Scripts

- `npm run dev` - Start development server
- `npm run build` - Build for production
- `npm run start` - Start production server
- `npm run lint` - Run ESLint

## API Integration

The frontend integrates with all backend endpoints:

### User Management
- GET `/api/v1/users` - Fetch all users
- POST `/api/v1/users` - Create user
- PUT `/api/v1/users/{id}` - Update user
- DELETE `/api/v1/users/{id}` - Delete user

### Product Management
- GET `/api/v1/products` - Fetch all products
- POST `/api/v1/products` - Create product
- PUT `/api/v1/products/{id}` - Update product
- PATCH `/api/v1/products/{id}/stock` - Update stock
- DELETE `/api/v1/products/{id}` - Delete product

### Sales Management
- GET `/api/v1/sales` - Fetch all sales
- POST `/api/v1/sales` - Create sale
- GET `/api/v1/sales/{id}` - Get sale by ID
- GET `/api/v1/sales/customer/{id}` - Get sales by customer

## Features Overview

### Dashboard
- Overview statistics (total products, revenue, low stock alerts)
- Recent sales summary
- Quick navigation to all modules

### Product Management
- Add, edit, delete products
- Stock level monitoring
- Price management
- Low stock alerts

### Sales Management
- Create new sales with multiple items
- Customer assignment (optional)
- Automatic stock deduction
- Sales history and details view

### User Management
- Add, edit, delete users
- Customer database for sales

### Authentication
- Secure login system
- Session management
- Protected routes

## Internationalization

The app supports multiple languages:
- **English** (default)
- **Thai**

Language files are located in `public/locales/`

## State Management

Using Redux Toolkit with organized slices:
- `userSlice` - User management state
- `productSlice` - Product management state  
- `saleSlice` - Sales management state

## Best Practices Implemented

- **TypeScript**: Full type safety
- **Component Architecture**: Reusable components
- **Error Handling**: Comprehensive error handling
- **Loading States**: Loading indicators throughout
- **Form Validation**: Client-side validation
- **Responsive Design**: Mobile-first approach
- **Code Organization**: Clear folder structure
- **Performance**: Optimized components and API calls
