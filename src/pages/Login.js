import React, { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import axios from 'axios'; // Import axios
import './Login.css';

function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const navigate = useNavigate();  // Hook เพื่อใช้การนำทาง

  const handleLogin = async (e) => {
    e.preventDefault();
    console.log('Logging in with:', email, password);  // เพิ่มการ log ข้อมูลที่ส่งไป
  
    try {
      const response = await axios.post('http://localhost:8082/login', { email, password });
      const token = response.data.token;
  
      if (token) {
        localStorage.setItem('authToken', token);
        navigate('/');
      }
    } catch (error) {
      console.error('Error during login:', error);
      alert('Invalid credentials. Please try again.');
    }
  };
  

  return (
    <div className="login-container">
      <h2>เข้าสู่ระบบ</h2>
      <form onSubmit={handleLogin}>
        <input
          type="email"
          placeholder="อีเมล"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
        />
        <input
          type="password"
          placeholder="รหัสผ่าน"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
        <button type="submit">Login</button>
      </form>

      <div className="register-link">
        <p>ยังไม่มีบัญชีผู้ใช้ ? <Link to="/user/register">สมัครสมาชิกที่นี่</Link></p>
      </div>
    </div>
  );
}

export default Login;
