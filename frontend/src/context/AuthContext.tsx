import React, { createContext, useContext, ReactNode, useState } from 'react';

interface User {
  email: string;
  role: string;
  permissions?: string[];
}

interface AuthContextType {
  isAuthenticated: boolean;
  user: User | null;
  login: (email: string, password: string) => Promise<void>;
  logout: () => void;
  hasRole: (role: string) => boolean;
  hasPermission: (permission: string) => boolean;
}


const AuthContext = createContext<AuthContextType | null>(null);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [user, setUser] = useState<User | null>(null);


  const login = async (email: string, password: string) => {
    try {
      const response = await fetch('/api/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password })
      });

      if (!response.ok) {
        throw new Error('Login failed');
      }

      const userData = await response.json();
      const authData = {
        isAuthenticated: true,
        user: {
          email: userData.email,
          role: userData.role || 'user',
          permissions: userData.permissions || []
        }
      };

      localStorage.setItem('auth', JSON.stringify(authData));
      setIsAuthenticated(true);
      setUser(authData.user);
    } catch (error) {
      console.error('Login error:', error);
      throw error;
    }
  };


  const logout = () => {
    fetch('/api/auth/logout', { method: 'POST' });
    localStorage.removeItem('auth');
    setIsAuthenticated(false);
    setUser(null);
  };

  const hasRole = (role: string) => {
    return user?.role === role;
  };

  const hasPermission = (permission: string) => {
    return user?.permissions?.includes(permission) || false;
  };



  return (
    <AuthContext.Provider value={{ 
      isAuthenticated, 
      user, 
      login, 
      logout,
      hasRole,
      hasPermission 
    }}>

      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
}
