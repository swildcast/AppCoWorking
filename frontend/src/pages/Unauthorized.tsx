import React from 'react';
import { Link, useLocation } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';

const Unauthorized: React.FC = () => {
  const location = useLocation<{ from: { pathname: string } }>();
  const { logout } = useAuth();

  return (
    <div className="min-h-screen bg-gray-100 flex flex-col justify-center items-center p-4">
      <div className="max-w-md w-full bg-white rounded-lg shadow-md p-6 text-center">
        <h1 className="text-2xl font-bold text-red-600 mb-4">Acceso no autorizado</h1>
        <p className="mb-6">
          No tienes los permisos necesarios para acceder a: 
          <span className="font-semibold ml-1">
            {location.state?.from?.pathname || 'esta página'}
          </span>
        </p>
        
        <div className="space-y-3">
          <Link 
            to="/"
            className="block w-full px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 transition"
          >
            Ir al inicio
          </Link>
          
          <button
            onClick={logout}
            className="w-full px-4 py-2 bg-gray-200 text-gray-800 rounded hover:bg-gray-300 transition"
          >
            Cerrar sesión
          </button>
        </div>
      </div>
    </div>
  );
};

export default Unauthorized;
