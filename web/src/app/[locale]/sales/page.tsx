'use client';

import { useEffect, useState } from 'react';
import { useAppDispatch, useAppSelector } from '@/store';
import { fetchSales, createSale } from '@/store/slices/saleSlice';
import { fetchProducts } from '@/store/slices/productSlice';
import { fetchUsers } from '@/store/slices/userSlice';
import DashboardLayout from '@/components/layout/DashboardLayout';
import Modal from '@/components/ui/Modal';
import { Sale } from '@/types/api';
import { PlusIcon, EyeIcon } from '@heroicons/react/24/outline';

export default function Sales() {
  const dispatch = useAppDispatch();
  const { sales, loading, error } = useAppSelector((state) => state.sales);
  const { products } = useAppSelector((state) => state.products);
  const { users } = useAppSelector((state) => state.users);
  const [modalOpen, setModalOpen] = useState(false);
  const [viewModalOpen, setViewModalOpen] = useState(false);
  const [selectedSale, setSelectedSale] = useState<Sale | null>(null);
  const [saleItems, setSaleItems] = useState([{ product_id: 0, quantity: 1 }]);
  const [customerId, setCustomerId] = useState<number | undefined>();
  const [formLoading, setFormLoading] = useState(false);

  useEffect(() => {
    dispatch(fetchSales());
    dispatch(fetchProducts());
    dispatch(fetchUsers());
  }, [dispatch]);

  const handleCreateSale = async (e: React.FormEvent) => {
    e.preventDefault();
    setFormLoading(true);
    try {
      const validItems = saleItems.filter(item => item.product_id > 0 && item.quantity > 0);
      if (validItems.length === 0) {
        alert('Please add at least one valid item');
        return;
      }

      await dispatch(createSale({
        ...(customerId ? { customer_id: customerId } : {}),
        items: validItems
      })).unwrap();

      setModalOpen(false);
      setSaleItems([{ product_id: 0, quantity: 1 }]);
      setCustomerId(undefined);
    } catch (error) {
      console.error('Failed to create sale:', error);
      alert('Failed to create sale. Please check product availability.');
    } finally {
      setFormLoading(false);
    }
  };

  const addSaleItem = () => {
    setSaleItems([...saleItems, { product_id: 0, quantity: 1 }]);
  };

  const updateSaleItem = (index: number, field: string, value: number) => {
    const updatedItems = [...saleItems];
    updatedItems[index] = { ...updatedItems[index], [field]: value };
    setSaleItems(updatedItems);
  };

  const removeSaleItem = (index: number) => {
    setSaleItems(saleItems.filter((_, i) => i !== index));
  };

  const openCreateModal = () => {
    setSaleItems([{ product_id: 0, quantity: 1 }]);
    setCustomerId(undefined);
    setModalOpen(true);
  };

  const openViewModal = (sale: Sale) => {
    setSelectedSale(sale);
    setViewModalOpen(true);
  };

  return (
    <DashboardLayout>
      <div className="space-y-6">
        <div className="sm:flex sm:items-center sm:justify-between">
          <div>
            <h1 className="text-2xl font-semibold text-gray-900">Sales</h1>
            <p className="mt-1 text-sm text-gray-600">View and create sales transactions</p>
          </div>
          <div className="mt-4 sm:mt-0">
            <button
              onClick={openCreateModal}
              className="inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            >
              <PlusIcon className="-ml-1 mr-2 h-5 w-5" />
              New Sale
            </button>
          </div>
        </div>

        {error && (
          <div className="bg-red-50 border border-red-200 rounded-md p-4">
            <div className="text-sm text-red-700">{error}</div>
          </div>
        )}

        <div className="bg-white shadow overflow-hidden sm:rounded-md">
          {loading && sales.length === 0 ? (
            <div className="text-center py-12">
              <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600 mx-auto"></div>
              <p className="mt-2 text-gray-500">Loading sales...</p>
            </div>
          ) : sales.length === 0 ? (
            <div className="text-center py-12">
              <p className="text-gray-500">No sales found. Create your first sale!</p>
            </div>
          ) : (
            <ul className="divide-y divide-gray-200">
              {sales.map((sale) => (
                <li key={sale.id}>
                  <div className="px-4 py-4 sm:px-6">
                    <div className="flex items-center justify-between">
                      <div className="flex-1 min-w-0">
                        <div className="flex items-center space-x-3">
                          <h3 className="text-lg font-medium text-gray-900">
                            Sale #{sale.id}
                          </h3>
                          <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
                            ${sale.total.toFixed(2)}
                          </span>
                        </div>
                        <p className="mt-1 text-sm text-gray-500">
                          Customer: {sale.customer?.name || 'Walk-in Customer'}
                        </p>
                        <p className="mt-1 text-xs text-gray-400">
                          {new Date(sale.created_at).toLocaleString()}
                        </p>
                      </div>
                      <div className="flex items-center space-x-2">
                        <button
                          onClick={() => openViewModal(sale)}
                          className="p-2 text-gray-400 hover:text-blue-600"
                        >
                          <EyeIcon className="h-5 w-5" />
                        </button>
                      </div>
                    </div>
                  </div>
                </li>
              ))}
            </ul>
          )}
        </div>
      </div>

      <Modal
        open={modalOpen}
        onClose={() => setModalOpen(false)}
        title="Create New Sale"
      >
        <form onSubmit={handleCreateSale} className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-gray-700">
              Customer (Optional)
            </label>
            <select
              className="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500 sm:text-sm px-3 py-2 border text-gray-500 bg-white"
              value={customerId || ''}
              onChange={(e) => setCustomerId(e.target.value ? parseInt(e.target.value) : undefined)}
            >
              <option value="">Walk-in Customer</option>
              {users.map((user) => (
                <option key={user.id} value={user.id}>
                  {user.name} ({user.email})
                </option>
              ))}
            </select>
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">
              Items
            </label>
            {saleItems.map((item, index) => (
              <div key={index} className="flex space-x-2 mb-2">
                <select
                  className="flex-1 border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500 sm:text-sm px-3 py-2 border text-gray-500 bg-white"
                  value={item.product_id}
                  onChange={(e) => updateSaleItem(index, 'product_id', parseInt(e.target.value))}
                  required
                >
                  <option value={0}>Select Product</option>
                  {products.map((product) => (
                    <option key={product.id} value={product.id}>
                      {product.name} - ${product.price.toFixed(2)} (Stock: {product.stock})
                    </option>
                  ))}
                </select>
                <input
                  type="number"
                  min="1"
                  className="w-20 border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500 sm:text-sm px-3 py-2 border text-gray-500 bg-white"
                  value={item.quantity}
                  onChange={(e) => updateSaleItem(index, 'quantity', parseInt(e.target.value))}
                  required
                />
                <button
                  type="button"
                  onClick={() => removeSaleItem(index)}
                  className="px-3 py-2 text-red-600 hover:text-red-800"
                  disabled={saleItems.length === 1}
                >
                  Remove
                </button>
              </div>
            ))}
            <button
              type="button"
              onClick={addSaleItem}
              className="text-blue-600 hover:text-blue-800 text-sm"
            >
              + Add Item
            </button>
          </div>

          <div className="flex justify-end space-x-3 pt-4">
            <button
              type="button"
              onClick={() => setModalOpen(false)}
              className="px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50"
            >
              Cancel
            </button>
            <button
              type="submit"
              disabled={formLoading}
              className="px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 disabled:opacity-50"
            >
              {formLoading ? 'Creating...' : 'Create Sale'}
            </button>
          </div>
        </form>
      </Modal>

      <Modal
        open={viewModalOpen}
        onClose={() => setViewModalOpen(false)}
        title={`Sale #${selectedSale?.id} Details`}
      >
        {selectedSale && (
          <div className="space-y-4">
            <div>
              <p><strong>Customer:</strong> {selectedSale.customer?.name || 'Walk-in Customer'}</p>
              <p><strong>Date:</strong> {new Date(selectedSale.created_at).toLocaleString()}</p>
              <p><strong>Total:</strong> ${selectedSale.total.toFixed(2)}</p>
            </div>
            <div>
              <h4 className="font-medium text-gray-900 mb-2">Items:</h4>
              <ul className="space-y-2">
                {selectedSale.sale_items.map((item) => (
                  <li key={item.id} className="flex justify-between">
                    <span>{item.product.name} x {item.quantity}</span>
                    <span>${item.subtotal.toFixed(2)}</span>
                  </li>
                ))}
              </ul>
            </div>
          </div>
        )}
      </Modal>
    </DashboardLayout>
  );
}