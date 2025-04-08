import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import HomePage from './pages/HomePage';
import Services from './pages/Services'; // คุณจะเพิ่มในภายหลัง
import Login from './pages/Login'; // คุณจะเพิ่มในภายหลัง
import Register from './pages/Register'; // คุณจะเพิ่มในภายหลัง

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/services" element={<Services />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
      </Routes>
    </Router>
  );
}

export default App;
