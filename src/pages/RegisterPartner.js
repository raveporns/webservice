import React, { useState } from "react";
import { useNavigate, Link } from 'react-router-dom';
import axios from "axios";
import "./Register.css";

function RegisterPartner() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  const navigate = useNavigate();

  const handleRegister = async (e) => {
    e.preventDefault();

    if (password !== confirmPassword) {
      alert("Passwords don't match");
      return;
    }

    const userData = {
      email: email,
      password: password,
      role: "partner"
    };

    try {
      const response = await axios.post(
        "http://localhost:8082/register",
        userData
      );

      if (response.status === 200) {
        console.log("Partner registration successful:", response.data);
        navigate("/partner/login");
      }
    } catch (error) {
      console.error("Error during partner registration:", error);
      alert("ไม่สามารถสมัครสมาชิกได้ กรุณาลองใหม่อีกครั้ง");
    }
  };

  return (
    <div className="register-container">
      <h2>สมัครสมาชิกพาร์ทเนอร์</h2>
      <form onSubmit={handleRegister}>
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
        <input
          type="password"
          placeholder="ยืนยันรหัสผ่าน"
          value={confirmPassword}
          onChange={(e) => setConfirmPassword(e.target.value)}
          required
        />
        <button type="submit">สมัครสมาชิก</button>
      </form>
      <div className="login-link">
        <p>
          มีบัญชีอยู่แล้ว? <Link to="/partner/login">เข้าสู่ระบบ</Link>
        </p>
      </div>
    </div>
  );
}

export default RegisterPartner;
