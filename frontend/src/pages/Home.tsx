import React from 'react';
import { Link } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';

const Home: React.FC = () => {
  const { isAuthenticated } = useAuth();

  return (
    <div className="min-h-screen bg-gray-100 flex flex-col items-center justify-center p-4">
      <div className="max-w-md w-full bg-white rounded-lg shadow-md p-6 text-center">
        <h1 className="text-3xl font-bold text-blue-600 mb-6">Bienvenido</h1>
        
        {isAuthenticated ? (
          <Link 
            to="/dashboard"
            className="px-6 py-3 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition"
          >
            Ir al Dashboard
          </Link>
        ) : (
          <div className="space-y-3">
            <Link 
              to="/login"
              className="block px-6 py-3 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition"
            >
              Iniciar sesión
            </Link>
            <Link 
              to="/dashboard"
              className="block px-6 py-3 bg-gray-200 text-gray-800 rounded-lg hover:bg-gray-300 transition"
            >
              Ver demo (sin login)
            </Link>
          </div>
        )}
      </div>
    </div>
  );
};

export default Home;
