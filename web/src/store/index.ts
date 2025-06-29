import { configureStore, createSlice } from '@reduxjs/toolkit';
import { TypedUseSelectorHook, useDispatch, useSelector } from 'react-redux';

// Dummy slice since we're using custom hooks instead of Redux
const dummySlice = createSlice({
  name: 'dummy',
  initialState: {},
  reducers: {},
});

export const store = configureStore({
  reducer: {
    dummy: dummySlice.reducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export const useAppDispatch = () => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;