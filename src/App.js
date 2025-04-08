import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import HomePage from './pages/HomePage';
import Services from './pages/Services';
import Login from './pages/Login';
import Register from './pages/Register';
import Login_Partner from './pages/Login_partner';
import Register_Partner from './pages/RegisterPartner';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/services" element={<Services />} />
        <Route path="/user/login" element={<Login />} />
        <Route path="/user/register" element={<Register />} />
        <Route path="/partner/login" element={<Login_Partner />} />
        <Route path="/partner/register" element={<Register_Partner />} />
        
      </Routes>
    </Router>
  );
}

export default App;
