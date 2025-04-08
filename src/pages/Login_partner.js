import React, { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import './Login.css';

function Login_Partner() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const navigate = useNavigate();

  const handleLogin = (e) => {
    e.preventDefault();

    // 🔒 ตัวอย่างจำลองการล็อกอินสำหรับ partner
    console.log('Partner logging in with:', email, password);

    // TODO: ตรวจสอบสิทธิ์หรือเรียก API แยกสำหรับ partner login

    // นำทางไปยังหน้า dashboard สำหรับ partner
    navigate('/partner/dashboard');
  };

  return (
    <div className="login-container">
      <h2>Partner Login</h2>
      <form onSubmit={handleLogin}>
        <input
          type="email"
          placeholder="Partner Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
        />
        <input
          type="password"
          placeholder="Partner Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
        <button type="submit">Login as Partner</button>
      </form>

      <div className="register-link">
        <p>ยังไม่มีบัญชีพาร์ทเนอร์? <Link to="/partner/register">สมัครที่นี่</Link></p>
        <p><Link to="/">กลับหน้าแรก</Link></p>
      </div>
    </div>
  );
}

export default Login_Partner;
