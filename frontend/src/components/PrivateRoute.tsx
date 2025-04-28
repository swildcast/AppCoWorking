import React from 'react';
import { Route, Redirect, RouteProps } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';

interface PrivateRouteProps extends RouteProps {
  component: React.ComponentType<any>;
  requiredRole?: string;
  requiredPermission?: string;
}


interface LocationState {
  from: {
    pathname: string;
  };
}

const PrivateRoute: React.FC<PrivateRouteProps> = ({ 
  component: Component,
  requiredRole,
  requiredPermission,
  ...rest 
}) => {
  const { isAuthenticated, hasRole, hasPermission } = useAuth();


  const hasRequiredAccess = () => {
    if (requiredRole && !hasRole(requiredRole)) return false;
    if (requiredPermission && !hasPermission(requiredPermission)) return false;
    return true;
  };



  return (
    <Route
      {...rest}
      render={(props) =>
        isAuthenticated ? (
          hasRequiredAccess() ? (

            <Component {...props} />
          ) : (
            <Redirect to={{
              pathname: '/unauthorized',
              state: { from: props.location }
            }} />
          )
        ) : (
          <Redirect to={{
            pathname: '/login',
            state: { from: props.location }
          }} />
        )
      }
    />
  );
};


export default PrivateRoute;
