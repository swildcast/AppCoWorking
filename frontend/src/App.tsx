import React from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import { AuthProvider } from '@context/AuthContext';
import PrivateRoute from '@components/PrivateRoute';
import Home from '@pages/Home';
import Login from '@pages/Login';
import Dashboard from '@pages/Dashboard';
import AdminDashboard from '@pages/AdminDashboard';
import Unauthorized from '@pages/Unauthorized';



function App() {
  return (
    <Router>
      <AuthProvider>
        <Switch>
          <Route exact path="/" component={Home} />
          <Route path="/login" component={Login} />
          <Route path="/unauthorized" component={Unauthorized} />
          
          <PrivateRoute path="/dashboard" component={Dashboard} />
          <PrivateRoute 
            path="/admin" 
            component={AdminDashboard} 
            requiredRole="admin" 
          />
        </Switch>
      </AuthProvider>
    </Router>

  );
}

export default App;
