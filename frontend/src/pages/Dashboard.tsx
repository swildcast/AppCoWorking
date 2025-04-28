import React from 'react'
import { useAuth } from '../context/AuthContext'

const DashboardPage: React.FC = () => {
  const { user, logout } = useAuth()

  return (
    <div className="min-h-screen bg-gray-50">
      <header className="bg-white shadow">
        <div className="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8 flex justify-between items-center">
          <h1 className="text-2xl font-bold text-gray-900">Panel de Control</h1>
          <div className="flex items-center space-x-4">
            <span className="text-gray-700">Hola, {user?.email}</span>
            <button
              onClick={logout}
              className="text-red-600 hover:text-red-800 px-3 py-2 rounded-md text-sm font-medium"
            >
              Cerrar Sesión
            </button>
          </div>
        </div>
      </header>
      <main className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        <div className="px-4 py-6 sm:px-0">
          <div className="border-4 border-dashed border-gray-200 rounded-lg h-96 p-8 text-center">
            <h2 className="text-xl font-semibold mb-4">Contenido del Dashboard</h2>
            <p className="text-gray-500">Aquí irán los módulos principales de la aplicación</p>
          </div>
        </div>
      </main>
    </div>
  )
}

export default DashboardPage
