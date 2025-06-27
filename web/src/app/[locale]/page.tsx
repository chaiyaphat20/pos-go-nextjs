'use client';

import { useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { useLocale } from 'next-intl';

export default function Home() {
  const router = useRouter();
  const locale = useLocale();

  useEffect(() => {
    router.push(`/dashboard`);
  }, [router, locale]);

  return (
    <div className="min-h-screen flex items-center justify-center">
      <div className="animate-spin rounded-full h-32 w-32 border-b-2 border-blue-600"></div>
    </div>
  );
}
