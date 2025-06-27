import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { Sale, CreateSaleRequest } from '@/types/api';
import * as saleApi from '@/services/saleApi';

interface SaleState {
  sales: Sale[];
  loading: boolean;
  error: string | null;
}

const initialState: SaleState = {
  sales: [],
  loading: false,
  error: null,
};

export const fetchSales = createAsyncThunk(
  'sales/fetchSales',
  async () => {
    const response = await saleApi.getSales();
    return response;
  }
);

export const fetchSaleById = createAsyncThunk(
  'sales/fetchSaleById',
  async (id: number) => {
    const response = await saleApi.getSaleById(id);
    return response;
  }
);

export const fetchSalesByCustomer = createAsyncThunk(
  'sales/fetchSalesByCustomer',
  async (customerId: number) => {
    const response = await saleApi.getSalesByCustomer(customerId);
    return response;
  }
);

export const createSale = createAsyncThunk(
  'sales/createSale',
  async (saleData: CreateSaleRequest) => {
    const response = await saleApi.createSale(saleData);
    return response;
  }
);

const saleSlice = createSlice({
  name: 'sales',
  initialState,
  reducers: {
    clearError: (state) => {
      state.error = null;
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchSales.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(fetchSales.fulfilled, (state, action: PayloadAction<Sale[]>) => {
        state.loading = false;
        state.sales = action.payload;
      })
      .addCase(fetchSales.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error.message || 'Failed to fetch sales';
      })
      .addCase(createSale.fulfilled, (state, action: PayloadAction<Sale>) => {
        state.sales.push(action.payload);
      })
      .addCase(fetchSalesByCustomer.fulfilled, (state, action: PayloadAction<Sale[]>) => {
        state.sales = action.payload;
      });
  },
});

export const { clearError } = saleSlice.actions;
export default saleSlice.reducer;