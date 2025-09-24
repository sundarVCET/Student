import React, { useState,useRef } from 'react';
import axios from 'axios';
import { useProfile } from '../context/ProfileContext'; // Import useProfile
import { useSelector } from 'react-redux';

const ImageUploadPage = () => {
  const [imageName, setImageName] = useState('');
  const [description, setDescription] = useState('');
  const [fileType, setFileType] = useState('');
  const [data, setData] = useState('');
  const [selectedFile, setSelectedFile] = useState(null);
  const [previewUrl, setPreviewUrl] = useState('');
  const { updateProfileImage } = useProfile(); // Use the custom hook
  const { currentUser } = useSelector(state => state.user);
  const [successMessage, setSuccessMessage] = useState(''); // ✅ new state
  const fileInputRef = useRef(null); // ✅ ref for file input

  const userId = currentUser._id;
  const userRole = currentUser.role;

  const handleImageChange = (e) => {
    const file = e.target.files[0];
    if (file) {
      const reader = new FileReader();
      reader.onloadend = () => {
        const base64String = reader.result.replace("data:", "").replace(/^.+,/, "");
        setData(`data:${file.type};base64,${base64String}`);
      };
      reader.readAsDataURL(file);
      setImageName(file.name);
      setSelectedFile(file);
      setPreviewUrl(URL.createObjectURL(file));
      setFileType(file.type);
    }
  };

  const onFileUpload = async (e) => {
    e.preventDefault();
    const payload = { imageName, data, fileType, description, userId, userRole };
    const token = localStorage.getItem('token');

    try {
      const response = await axios.post(
        `${process.env.REACT_APP_BASE_URL}/AddImage`,
        payload,
        {
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
          }
        }
      );
      //const response = await axios.post('http://localhost:8080/AddImage', payload);
      console.log('Response:', response.data);

      if (response.data && response.data.userId) {

        updateProfileImage(response.data.userId);
        setSuccessMessage('Image uploaded successfully! ✅'); // ✅ show success
        // Option 1: reset form
        setImageName('');
        setDescription('');
        setFileType('');
        setData('');
        setSelectedFile(null);
        setPreviewUrl('');

        if (fileInputRef.current) {
          fileInputRef.current.value = ''; // clear file input
        }

        // Option 2 (if you need full reload):
        // window.location.reload();

      }
    } catch (error) {
      console.error('Error uploading image:', error);
    }
  };

  const fileData = () => {
    if (selectedFile) {
      return (
        <div>
          <h2>File Details:</h2>
          <p>File Name: {selectedFile.name}</p>
          <p>File Type: {selectedFile.type}</p>
          <p>Last Modified: {selectedFile.lastModifiedDate.toDateString()}</p>
        </div>
      );
    } else {
      return (
        <div>
          <br />
          <h4>Choose before Pressing the Upload button</h4>
        </div>
      );
    }
  };

  return (
    <div>
      <br />
      <h1>Upload the Image</h1>
      <br />
      <div>
        <input
          type="file"
          accept="image/png,image/jpeg"
          onChange={handleImageChange}
          ref={fileInputRef} // ✅ attach ref
        />
        <input
          type="text"
          placeholder="Description"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
        />
        <button onClick={onFileUpload} disabled={!selectedFile}>
          Upload!
        </button>
      </div>
      <br />
      {successMessage && <p style={{ color: 'green' }}>{successMessage}</p>} {/* ✅ success msg */}
      {previewUrl && (
        <div>
          <h3>Image Preview:</h3>
          <img src={previewUrl} alt="Selected" style={{ width: '300px', height: 'auto' }} />
        </div>
      )}
      {fileData()}
    </div>
  );
};

export default ImageUploadPage;
