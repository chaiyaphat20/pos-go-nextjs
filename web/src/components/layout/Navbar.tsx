'use client';

import { useSession, signOut } from 'next-auth/react';
import type { Session } from 'next-auth';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { Bars3Icon, XMarkIcon } from '@heroicons/react/24/outline';
import { useState } from 'react';
import { useTranslations } from 'next-intl';
import LocaleSwitcher from '../ui/LocaleSwitcher';

export default function Navbar() {
  const { data: session } = useSession();
  const router = useRouter();
  const t = useTranslations('navigation');
  const tButtons = useTranslations('buttons');
  const tMessages = useTranslations('messages');
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false);

  // สร้าง navigation menu ตาม role
  const getNavigation = () => {
    const baseNav = [
      { name: t('dashboard'), href: '/dashboard' },
      { name: t('products'), href: '/products' },
    ];

    if ((session?.user as any)?.role === 'admin') {
      // Admin: Dashboard, Products (full access), Users
      return [
        ...baseNav,
        { name: t('users'), href: '/users' },
      ];
    } else {
      // User: Dashboard, Products (read-only + buy), Shopping Cart, Orders
      return [
        ...baseNav,
        { name: t('cart'), href: '/cart' },
        { name: t('orders'), href: '/orders' },
      ];
    }
  };

  const navigation = getNavigation();

  const handleSignOut = async () => {
    await signOut({ redirect: false });
    router.push('/auth/signin');
  };

  if (!session) return null;

  return (
    <nav className="bg-white shadow-lg">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between h-16">
          <div className="flex">
            <div className="flex-shrink-0 flex items-center">
              <Link href="/dashboard" className="text-xl font-bold text-gray-800">
                Go Clean POS
              </Link>
            </div>
            <div className="hidden sm:ml-6 sm:flex sm:space-x-8">
              {navigation.map((item) => (
                <Link
                  key={item.name}
                  href={item.href}
                  className="border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium"
                >
                  {item.name}
                </Link>
              ))}
            </div>
          </div>
          <div className="hidden sm:ml-6 sm:flex sm:items-center">
            <div className="ml-3 relative">
              <div className="flex items-center space-x-4">
                <LocaleSwitcher />
                <span className="text-gray-700">{tMessages('welcomeUser', { name: session.user?.name || '' })}</span>
                <button
                  onClick={handleSignOut}
                  className="bg-red-600 hover:bg-red-700 text-white px-4 py-2 rounded-md text-sm font-medium"
                >
                  {tButtons('signout')}
                </button>
              </div>
            </div>
          </div>
          <div className="-mr-2 flex items-center sm:hidden">
            <button
              onClick={() => setMobileMenuOpen(!mobileMenuOpen)}
              className="bg-white inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-gray-500 hover:bg-gray-100"
            >
              {mobileMenuOpen ? (
                <XMarkIcon className="block h-6 w-6" />
              ) : (
                <Bars3Icon className="block h-6 w-6" />
              )}
            </button>
          </div>
        </div>
      </div>

      {mobileMenuOpen && (
        <div className="sm:hidden">
          <div className="pt-2 pb-3 space-y-1">
            {navigation.map((item) => (
              <Link
                key={item.name}
                href={item.href}
                className="border-transparent text-gray-500 hover:bg-gray-50 hover:border-gray-300 hover:text-gray-700 block pl-3 pr-4 py-2 border-l-4 text-base font-medium"
                onClick={() => setMobileMenuOpen(false)}
              >
                {item.name}
              </Link>
            ))}
          </div>
          <div className="pt-4 pb-3 border-t border-gray-200">
            <div className="flex items-center px-4">
              <div className="flex-shrink-0">
                <span className="text-gray-800 font-medium">{session.user?.name}</span>
              </div>
            </div>
            <div className="mt-3 space-y-1">
              <div className="px-4 py-2">
                <LocaleSwitcher />
              </div>
              <button
                onClick={handleSignOut}
                className="block px-4 py-2 text-base font-medium text-red-600 hover:text-red-800 hover:bg-gray-50 w-full text-left"
              >
                {tButtons('signout')}
              </button>
            </div>
          </div>
        </div>
      )}
    </nav>
  );
}