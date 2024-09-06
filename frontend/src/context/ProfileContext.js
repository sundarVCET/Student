import React, { createContext, useContext, useState, useCallback } from 'react';
import axios from 'axios';

// Create a context
const ProfileContext = createContext();

// Create a provider component
const ProfileProvider = ({ children }) => {
  const [profileImage, setProfileImage] = useState('');

  const fetchProfileImage = useCallback(async (imageId) => {
    const token = localStorage.getItem('token');
  try {
      const result = await axios.get(
        `${process.env.REACT_APP_BASE_URL}/GetImage/${imageId}`,
        {
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
          }
        }
      );
      //const response = await axios.get(`http://localhost:8080/GetImage/${imageId}`);
  
      const base64Data = result.data.data; // Ensure this is a complete Base64 string
  
      // Check if `base64Data` starts with `data:image/...`
      if (base64Data.startsWith('data:')) {
        // Directly use the Base64 string as the image URL
        const imageUrl = base64Data;
        setProfileImage(imageUrl);
      } else {
       // throw new Error('Base64 data does not include MIME type prefix');
      }
    } catch (error) {
      //console.error('Error fetching profile image:', error);
    }
  }, []);
  

  const updateProfileImage = useCallback((imageUrl) => {
    //setProfileImage(imageUrl);
    fetchProfileImage(imageUrl)
  }, []);

  return (
    <ProfileContext.Provider value={{ profileImage, fetchProfileImage, updateProfileImage }}>
      {children}
    </ProfileContext.Provider>
  );
};

// Custom hook to use the ProfileContext
const useProfile = () => {
  const context = useContext(ProfileContext);
  if (!context) {
    throw new Error('useProfile must be used within a ProfileProvider');
  }
  return context;
};

export { ProfileProvider, useProfile };
