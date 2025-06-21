import React, { createContext, useContext, useState, useEffect } from "react";
import type { ReactNode } from "react";
import { getAccount } from "../api/auth";

interface User {
  employee_id: string;
  empl_surname: string;
  empl_name: string;
  empl_patronymic?: string;
  empl_role: string;
  salary: number;
  date_of_birth: string;
  date_of_start: string;
  phone_number: string;
  city: string;
  street: string;
  zip_code: string;
}

interface AuthContextType {
  user: User | null;
  loading: boolean;
  isAuthenticated: boolean;
  isManager: boolean;
  isCashier: boolean;
  login: (token: string) => void;
  logout: () => void;
  refreshUser: () => Promise<void>;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};

interface AuthProviderProps {
  children: ReactNode;
}

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);

  const isAuthenticated = !!user;
  const isManager = user?.empl_role === "Manager";
  const isCashier = user?.empl_role === "Cashier";

  const login = async (token: string) => {
    localStorage.setItem("token", token);
    await refreshUser();
  };

  const logout = () => {
    localStorage.removeItem("token");
    setUser(null);
  };

  const refreshUser = async () => {
    const token = localStorage.getItem("token");
    if (!token) {
      setLoading(false);
      return;
    }

    try {
      const response = await getAccount();
      // Convert AccountInfo to User, handling undefined values
      const user: User = {
        employee_id: response.employee_id || "",
        empl_surname: response.empl_surname || "",
        empl_name: response.empl_name || "",
        empl_patronymic: response.empl_patronymic || undefined,
        empl_role: response.empl_role || "",
        salary: response.salary || 0,
        date_of_birth: response.date_of_birth || "",
        date_of_start: response.date_of_start || "",
        phone_number: response.phone_number || "",
        city: response.city || "",
        street: response.street || "",
        zip_code: response.zip_code || "",
      };
      setUser(user);
    } catch (error) {
      console.error("Failed to refresh user:", error);
      logout();
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    refreshUser();
  }, []);

  const value: AuthContextType = {
    user,
    loading,
    isAuthenticated,
    isManager,
    isCashier,
    login,
    logout,
    refreshUser,
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};
