import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';  // Import useNavigate from react-router-dom
import axios from 'axios';  // Import axios
import './Register.css';

function Register() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');

  const navigate = useNavigate();  // Hook เพื่อใช้การนำทาง

  const handleRegister = async (e) => {
    e.preventDefault();
    
    // ตรวจสอบว่า password กับ confirmPassword ตรงกันหรือไม่
    if (password !== confirmPassword) {
      alert("Passwords don't match");
      return;
    }

    // สร้าง object ข้อมูลที่จะส่งไป
    const userData = {
      email: email,
      password: password
    };

    try {
      // ส่งข้อมูลลงทะเบียนไปยังเซิร์ฟเวอร์
      const response = await axios.post('http://localhost:8082/register', userData);
      
      if (response.status === 200) {
        // ลงทะเบียนสำเร็จ ให้ไปที่หน้า Login
        console.log('Registration successful:', response.data);
        navigate('/login');
      }
    } catch (error) {
      console.error('Error during registration:', error);
      alert('Registration failed. Please try again.');
    }
  };

  return (
    <div className="register-container">
      <h2>Register</h2>
      <form onSubmit={handleRegister}>
        <input
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
        <input
          type="password"
          placeholder="Confirm Password"
          value={confirmPassword}
          onChange={(e) => setConfirmPassword(e.target.value)}
          required
        />
        <button type="submit">Register</button>
      </form>
    </div>
  );
}

export default Register;
