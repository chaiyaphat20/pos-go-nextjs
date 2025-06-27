import { routing } from '@/i18n/routing'
import { useLocale } from 'next-intl'
import { useParams, usePathname } from 'next/navigation'
import { GlobeAltIcon } from '@heroicons/react/24/outline'
import LocaleSwitcherSelect from './LocaleSwitcherSelect'

const localeNames = {
  th: '🇹🇭 ไทย',
  en: '🇺🇸 English'
}

export default function LocaleSwitcher() {
  const locale = useLocale()
  const params = useParams()
  const pathname = usePathname()
  
  // Get current locale from URL path
  const windowPath = typeof window !== 'undefined' ? window.location.pathname : pathname
  const currentLocale = windowPath.startsWith('/en') ? 'en' : 'th'
  
  console.log('useLocale():', locale)
  console.log('params.locale:', params.locale)
  console.log('pathname:', pathname)
  console.log('currentLocale:', currentLocale)

  return (
    <div className="flex items-center gap-2">
      <GlobeAltIcon className="text-gray-500 h-4 w-4" />
      <LocaleSwitcherSelect defaultValue={currentLocale} label="Select a locale">
        {routing.locales.map((cur) => (
          <option key={cur} value={cur}>
            {localeNames[cur as keyof typeof localeNames]}
          </option>
        ))}
      </LocaleSwitcherSelect>
    </div>
  )
}