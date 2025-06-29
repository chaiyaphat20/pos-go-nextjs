import CredentialsProvider from 'next-auth/providers/credentials';
import { apiClient } from './api';

export const authOptions = {
  providers: [
    CredentialsProvider({
      name: 'credentials',
      credentials: {
        email: { label: 'Email', type: 'email' },
        password: { label: 'Password', type: 'password' }
      },
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      async authorize(credentials: any) {
        if (!credentials?.email || !credentials?.password) {
          return null;
        }

        try {
          const authResponse = await apiClient.login({
            email: credentials.email,
            password: credentials.password,
          });

          if (authResponse && authResponse.user) {
            return {
              id: authResponse.user.id,
              email: authResponse.user.email,
              name: authResponse.user.username,
              role: authResponse.user.role,
              accessToken: authResponse.access_token,
              refreshToken: authResponse.refresh_token,
            };
          }
        } catch (error) {
          console.error('Authentication failed:', error);
        }

        return null;
      }
    })
  ],
  pages: {
    signIn: '/auth/signin',
  },
  session: {
    strategy: 'jwt' as const,
  },
  callbacks: {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    async jwt({ token, user }: any) {
      if (user) {
        token.id = user.id;
        token.role = user.role;
        token.accessToken = user.accessToken;
        token.refreshToken = user.refreshToken;
      }
      return token;
    },
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    async session({ session, token }: any) {
      if (token && session.user) {
        session.user.id = token.id as string;
        session.user.role = token.role as string;
        session.accessToken = token.accessToken as string;
        session.refreshToken = token.refreshToken as string;
      }
      return session;
    },
  },
};