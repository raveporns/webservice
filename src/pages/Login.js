import React, { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';  // เพิ่มการนำเข้า Link
import './Login.css';

function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const navigate = useNavigate();  // Hook เพื่อใช้การนำทาง

  const handleLogin = (e) => {
    e.preventDefault();
    // Add login logic here
    console.log('Logging in with:', email, password);

    // หลังจากล็อกอินเสร็จให้ไปหน้า HomePage
    navigate('/');  // นำทางไปยังหน้า HomePage
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
        <p>ยังไม่มีบัญชีผู้ใช้ ? <Link to="/register">สมัครสมาชิกที่นี่</Link></p>
      </div>
    </div>
  );
}

export default Login;
