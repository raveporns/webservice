import React from 'react';
import { Link } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faBars, faSearch, faShoppingCart, faUser } from '@fortawesome/free-solid-svg-icons';
import './Navbar.css';

function Navbar() {
    return (
        <header className="navbar-container">
            <div className="navbar">
                {/* ส่วนซ้าย: เมนูและโลโก้ */}
                <div className="navbar-left">
                    <button className="hamburger-menu">
                        <FontAwesomeIcon icon={faBars} size="lg" />
                    </button>
                    <Link to="/" className="logo">
                        <img src={'/assets/logo.png'} alt="" />
                    </Link>
                </div>

                {/* ส่วนกลาง: ช่องค้นหา */}
                <div className="navbar-center">
                    <div className="search-box">
                        <input type="text" placeholder="ค้นหาสินค้าที่ต้องการ..." />
                        <button type="submit">
                            <FontAwesomeIcon icon={faSearch} size="lg" />
                        </button>
                    </div>
                </div>

                {/* ส่วนขวา: รถเข็นและผู้ใช้ */}
                <div className="navbar-right">
                    <Link to="/cart" className="cart-icon">
                        <FontAwesomeIcon icon={faShoppingCart} size="lg" />
                    </Link>
                    <Link to="/login" className="user-icon">
                        <FontAwesomeIcon icon={faUser} size="lg" />
                    </Link>
                </div>
            </div>
            <nav className="secondary-menu">
                <Link to="/installation" className="font-bold">บริการติดตั้ง</Link>
                <Link to="/repair">บริการซ่อมแซม</Link>
                <Link to="/about">เกี่ยวกับเรา</Link>
            </nav>

        </header>
    );
}

export default Navbar;