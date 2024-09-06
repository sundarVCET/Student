// src/hoc/withAuthRedirect.js

import React, { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { setRedirectCallback } from '../../utils/axiosInstance';

const withAuthRedirect = (WrappedComponent) => {
  return (props) => {
    const navigate = useNavigate();

    useEffect(() => {
      const handleUnauthorized = () => {
        navigate('/Adminlogin');
      };

      setRedirectCallback(handleUnauthorized);

      return () => {
        setRedirectCallback(() => {}); // Clean up the callback
      };
    }, [navigate]);

    return <WrappedComponent {...props} />;
  };
};

export default withAuthRedirect;
