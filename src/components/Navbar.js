import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faBars, faSearch, faShoppingCart, faUser } from '@fortawesome/free-solid-svg-icons';
import './Navbar.css';

function Navbar() {
    const [isDropdownOpen, setIsDropdownOpen] = useState(false);

    const toggleDropdown = () => {
        setIsDropdownOpen(!isDropdownOpen);
    };

    return (
        <header className="navbar-container">
            <div className="navbar">
                {/* ซ้าย */}
                <div className="navbar-left">
                    <button className="hamburger-menu">
                        <FontAwesomeIcon icon={faBars} size="lg" />
                    </button>
                    <Link to="/" className="logo">
                        <img src={'/assets/logo.png'} alt="" />
                    </Link>
                </div>

                {/* กลาง */}
                <div className="navbar-center">
                    <div className="search-box">
                        <input type="text" placeholder="ค้นหาสินค้าที่ต้องการ..." />
                        <button type="submit">
                            <FontAwesomeIcon icon={faSearch} size="lg" />
                        </button>
                    </div>
                </div>

                {/* ขวา */}
                <div className="navbar-right">
                    <Link to="/cart" className="cart-icon">
                        <FontAwesomeIcon icon={faShoppingCart} size="lg" />
                    </Link>

                    <div className="user-dropdown" onClick={toggleDropdown}>
                        <FontAwesomeIcon icon={faUser} size="lg" className="user-icon" />
                        {isDropdownOpen && (
                            <div className="dropdown-menu">
                                <Link to="/user/login">เข้าสู่ระบบ</Link>
                                <Link to="/user/register">สมัครสมาชิก</Link>
                                <Link to="/partner/register">สมัครเป็นหุ้นส่วน</Link>
                                <Link to="/user/login">ออกจากระบบ</Link>
                            </div>
                        )}
                    </div>
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
