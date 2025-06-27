# URL Routing Debug

## สิ่งที่ควรใช้งานได้:

### Thai (Default - no prefix):
- `/` → redirect to `/dashboard`
- `/dashboard` → Dashboard in Thai
- `/products` → Products in Thai  
- `/sales` → Sales in Thai
- `/users` → Users in Thai
- `/auth/signin` → Sign in in Thai

### English (with /en prefix):
- `/en` → redirect to `/en/dashboard`
- `/en/dashboard` → Dashboard in English
- `/en/products` → Products in English
- `/en/sales` → Sales in English  
- `/en/users` → Users in English
- `/en/auth/signin` → Sign in in English

## Current Configuration:
- Default locale: `th`
- Supported locales: `['th', 'en']`
- Locale prefix: `as-needed` with `/en` prefix for English only
- Middleware matcher: All routes except API, static files

## Potential Issues:
1. Check if middleware is correctly handling the routing
2. Make sure [locale] folder structure is correct
3. Verify that NextIntlClientProvider is wrapping the content properly

## Testing URLs:
1. Go to `http://localhost:3000` → should redirect to dashboard in Thai
2. Go to `http://localhost:3000/dashboard` → should show dashboard in Thai
3. Go to `http://localhost:3000/en/dashboard` → should show dashboard in English
4. Use language switcher to change between languages