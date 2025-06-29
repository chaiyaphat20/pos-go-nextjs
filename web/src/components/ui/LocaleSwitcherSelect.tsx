'use client'
import {  routing, usePathname } from '@/i18n/routing'
import { ReactNode } from 'react'

type Props = {
  children: ReactNode
  defaultValue: string
  label: string
}

export default function LocaleSwitcherSelect({ defaultValue, label }: Props) {
  const pathname = usePathname()
  function onSelectChange(event: React.ChangeEvent<HTMLSelectElement>) {
    const nextLocale = event.target.value
    console.log('Switching to locale:', nextLocale)
    console.log('Current pathname:', pathname)
    
    // Get current path and remove any locale prefix
    let pathWithoutLocale = pathname
    if (pathname.startsWith('/th')) {
      pathWithoutLocale = pathname.replace('/th', '')
    } else if (pathname.startsWith('/en')) {
      pathWithoutLocale = pathname.replace('/en', '')
    }
    
    // Build new path with new locale
    const newPath = `/${nextLocale}${pathWithoutLocale}`
    
    console.log('Path without locale:', pathWithoutLocale)
    console.log('New path:', newPath)
    
    // Use window.location for force navigation
    if (typeof window !== 'undefined') {
      window.location.href = newPath
    }
  }

  return (
    <select
      value={defaultValue}
      onChange={onSelectChange}
      aria-label={label}
      className="h-8 w-[120px] border border-gray-300 rounded px-2 bg-white text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
    >
      {routing.locales.map((locale) => (
        <option key={locale} value={locale}>
          {locale === 'th' ? '🇹🇭 ไทย' : '🇺🇸 English'}
        </option>
      ))}
    </select>
  )
}