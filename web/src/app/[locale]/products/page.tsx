'use client';

import { useEffect, useState } from 'react';
import { useSession } from 'next-auth/react';
import { useRouter } from 'next/navigation';
import { useAppDispatch, useAppSelector } from '@/store';
import { 
  fetchProducts, 
  createProduct, 
  updateProduct, 
  deleteProduct 
} from '@/store/slices/productSlice';
import DashboardLayout from '@/components/layout/DashboardLayout';
import Modal from '@/components/ui/Modal';
import ProductForm from '@/components/forms/ProductForm';
import { Product } from '@/types/api';
import { PencilIcon, TrashIcon, PlusIcon, ShoppingCartIcon } from '@heroicons/react/24/outline';
import { useTranslations } from 'next-intl';

export default function Products() {
  const { data: session } = useSession();
  const router = useRouter();
  const dispatch = useAppDispatch();
  const { products, loading, error } = useAppSelector((state) => state.products);
  const [modalOpen, setModalOpen] = useState(false);
  const [editingProduct, setEditingProduct] = useState<Product | null>(null);
  const [formLoading, setFormLoading] = useState(false);
  const [cart, setCart] = useState<{[key: number]: number}>({});
  const [tempCartItems, setTempCartItems] = useState<Array<{id: number, name: string, price: number, quantity: number, stock: number}>>([]);
  const [showConfirmButton, setShowConfirmButton] = useState(false);
  const t = useTranslations('products');
  const tShopping = useTranslations('shopping');
  const tMessages = useTranslations('messages');
  const tCart = useTranslations('cart');

  const isAdmin = (session?.user as any)?.role === 'admin';

  useEffect(() => {
    dispatch(fetchProducts());
  }, [dispatch]);

  const handleCreateProduct = async (data: {
    name: string;
    description: string;
    price: number;
    stock: number;
  }) => {
    setFormLoading(true);
    try {
      await dispatch(createProduct(data)).unwrap();
      setModalOpen(false);
      setEditingProduct(null);
    } catch (error) {
      console.error('Failed to create product:', error);
    } finally {
      setFormLoading(false);
    }
  };

  const handleUpdateProduct = async (data: {
    name: string;
    description: string;
    price: number;
    stock: number;
  }) => {
    if (!editingProduct) return;
    setFormLoading(true);
    try {
      await dispatch(updateProduct({ id: editingProduct.id, productData: data })).unwrap();
      setModalOpen(false);
      setEditingProduct(null);
    } catch (error) {
      console.error('Failed to update product:', error);
    } finally {
      setFormLoading(false);
    }
  };

  const handleDeleteProduct = async (id: number) => {
    if (confirm('Are you sure you want to delete this product?')) {
      try {
        await dispatch(deleteProduct(id)).unwrap();
      } catch (error) {
        console.error('Failed to delete product:', error);
      }
    }
  };

  const openCreateModal = () => {
    setEditingProduct(null);
    setModalOpen(true);
  };

  const openEditModal = (product: Product) => {
    setEditingProduct(product);
    setModalOpen(true);
  };

  const closeModal = () => {
    setModalOpen(false);
    setEditingProduct(null);
  };

  // Shopping cart functions (สำหรับ user)
  const addToCart = (productId: number, quantity: number = 1) => {
    const product = products.find(p => p.id === productId);
    if (!product) return;

    // เพิ่มลง temporary cart ก่อน (ยังไม่บันทึกลง localStorage)
    setTempCartItems(prev => {
      const existingItemIndex = prev.findIndex(item => item.id === productId);
      
      if (existingItemIndex > -1) {
        // เพิ่มจำนวนถ้ามีอยู่แล้ว
        const updated = [...prev];
        updated[existingItemIndex].quantity += quantity;
        return updated;
      } else {
        // เพิ่มรายการใหม่
        return [...prev, {
          id: product.id,
          name: product.name,
          price: product.price,
          quantity: quantity,
          stock: product.stock
        }];
      }
    });

    // อัพเดท local state เพื่อแสดงจำนวน
    setCart(prev => ({
      ...prev,
      [productId]: (prev[productId] || 0) + quantity
    }));

    // แสดงปุ่มยืนยัน
    setShowConfirmButton(true);
  };

  const confirmPurchase = () => {
    if (tempCartItems.length === 0) return;

    // ย้ายข้อมูลจาก temporary cart ไป localStorage
    const savedCart = localStorage.getItem('cart');
    let cartItems = savedCart ? JSON.parse(savedCart) : [];
    
    // รวม temporary items เข้ากับ cart ที่มีอยู่
    tempCartItems.forEach(tempItem => {
      const existingItemIndex = cartItems.findIndex((item: any) => item.id === tempItem.id);
      
      if (existingItemIndex > -1) {
        cartItems[existingItemIndex].quantity += tempItem.quantity;
      } else {
        cartItems.push(tempItem);
      }
    });
    
    localStorage.setItem('cart', JSON.stringify(cartItems));
    
    // ล้าง temporary state
    setTempCartItems([]);
    setCart({});
    setShowConfirmButton(false);
    
    // ไปหน้าตะกร้า
    router.push(`/cart`);
  };

  const cancelPurchase = () => {
    // ล้าง temporary cart
    setTempCartItems([]);
    setCart({});
    setShowConfirmButton(false);
  };

  const getTempCartTotal = () => {
    return tempCartItems.reduce((sum, item) => sum + item.quantity, 0);
  };

  const getCartQuantity = (productId: number) => {
    return cart[productId] || 0;
  };

  return (
    <DashboardLayout>
      <div className="space-y-6">
        <div className="sm:flex sm:items-center sm:justify-between">
          <div>
            <h1 className="text-2xl font-semibold text-gray-900">
              {isAdmin ? t('title') : tShopping('title')}
            </h1>
            <p className="mt-1 text-sm text-gray-600">
              {isAdmin ? t('subtitle') : tShopping('subtitle')}
            </p>
          </div>
          {isAdmin && (
            <div className="mt-4 sm:mt-0">
              <button
                onClick={openCreateModal}
                className="inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
              >
                <PlusIcon className="-ml-1 mr-2 h-5 w-5" />
                {t('addProduct')}
              </button>
            </div>
          )}
        </div>

        {error && (
          <div className="bg-red-50 border border-red-200 rounded-md p-4">
            <div className="text-sm text-red-700">{error}</div>
          </div>
        )}

        <div className="bg-white shadow overflow-hidden sm:rounded-md">
          {loading && products.length === 0 ? (
            <div className="text-center py-12">
              <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600 mx-auto"></div>
              <p className="mt-2 text-gray-700">Loading products...</p>
            </div>
          ) : products.length === 0 ? (
            <div className="text-center py-12">
              <p className="text-gray-700">No products found. Create your first product!</p>
            </div>
          ) : (
            <ul className="divide-y divide-gray-200">
              {products.map((product) => (
                <li key={product.id}>
                  <div className="px-4 py-4 sm:px-6">
                    <div className="flex items-center justify-between">
                      <div className="flex-1 min-w-0">
                        <div className="flex items-center space-x-3">
                          <h3 className="text-lg font-medium text-gray-900 truncate">
                            {product.name}
                          </h3>
                          <span className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
                            product.stock < 10 
                              ? 'bg-red-100 text-red-800' 
                              : 'bg-green-100 text-green-800'
                          }`}>
                            Stock: {product.stock}
                          </span>
                        </div>
                        <p className="mt-1 text-sm text-gray-700">
                          {product.description}
                        </p>
                        <div className="mt-2 flex items-center space-x-4 text-sm text-gray-700">
                          <span>Price: ${product.price.toFixed(2)}</span>
                          <span>•</span>
                          <span>ID: {product.id}</span>
                        </div>
                      </div>
                      <div className="flex items-center space-x-2">
                        {isAdmin ? (
                          // Admin controls - edit และ delete
                          <>
                            <button
                              onClick={() => openEditModal(product)}
                              className="p-2 text-gray-400 hover:text-blue-600"
                            >
                              <PencilIcon className="h-5 w-5" />
                            </button>
                            <button
                              onClick={() => handleDeleteProduct(product.id)}
                              className="p-2 text-gray-400 hover:text-red-600"
                            >
                              <TrashIcon className="h-5 w-5" />
                            </button>
                          </>
                        ) : (
                          // User controls - add to cart
                          <div className="flex items-center space-x-3">
                            {getCartQuantity(product.id) > 0 && (
                              <span className="text-sm text-gray-700 font-medium">
                                {tShopping('quantity')}: {getCartQuantity(product.id)}
                              </span>
                            )}
                            <button
                              onClick={() => addToCart(product.id)}
                              disabled={product.stock === 0}
                              className={`inline-flex items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md ${
                                product.stock === 0
                                  ? 'text-gray-400 bg-gray-100 cursor-not-allowed'
                                  : 'text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500'
                              }`}
                            >
                              <ShoppingCartIcon className="h-4 w-4 mr-1" />
                              {product.stock === 0 ? tShopping('outOfStock') : tShopping('addToCart')}
                            </button>
                          </div>
                        )}
                      </div>
                    </div>
                  </div>
                </li>
              ))}
            </ul>
          )}
        </div>

        {/* ปุ่มยืนยันการซื้อ */}
        {!isAdmin && showConfirmButton && tempCartItems.length > 0 && (
          <div className="fixed bottom-0 left-0 right-0 bg-white border-t shadow-lg p-4 z-50">
            <div className="max-w-7xl mx-auto flex items-center justify-between">
              <div className="flex items-center space-x-4">
                <ShoppingCartIcon className="h-6 w-6 text-gray-700" />
                <div>
                  <p className="text-sm text-gray-700 font-medium">
                    {tShopping('selectedItems')}: {getTempCartTotal()} {tShopping('items')}
                  </p>
                  <p className="text-lg font-semibold text-gray-900">
                    {tShopping('total')}: ${tempCartItems.reduce((sum, item) => sum + (item.price * item.quantity), 0).toFixed(2)}
                  </p>
                </div>
              </div>
              <div className="flex items-center space-x-3">
                <button
                  onClick={cancelPurchase}
                  className="px-4 py-2 border border-gray-300 rounded-md text-gray-800 font-medium hover:bg-gray-50"
                >
                  {tShopping('cancel')}
                </button>
                <button
                  onClick={confirmPurchase}
                  className="px-6 py-2 bg-green-600 hover:bg-green-700 text-white rounded-md font-medium"
                >
                  {tShopping('confirmPurchase')}
                </button>
              </div>
            </div>
          </div>
        )}
      </div>

      <Modal
        open={modalOpen}
        onClose={closeModal}
        title={editingProduct ? 'Edit Product' : 'Create Product'}
      >
        <ProductForm
          product={editingProduct || undefined}
          onSubmit={editingProduct ? handleUpdateProduct : handleCreateProduct}
          onCancel={closeModal}
          loading={formLoading}
        />
      </Modal>
    </DashboardLayout>
  );
}