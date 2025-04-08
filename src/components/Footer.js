import React from 'react';
import './Footer.css';

function Footer() {
    return (
        <footer className="footer">
            <div className="footer-container">
                <div className="footer-section">
                    <h4>เกี่ยวกับเรา</h4>
                    <p>บริการติดตั้งและซ่อมแซมครบวงจร โดยทีมงานมืออาชีพที่คุณวางใจได้</p>
                </div>

                <div className="footer-section">
                    <h4>ติดต่อเรา</h4>
                    <p>โทร: 02-123-4567</p>
                    <p>อีเมล: support@fixnserve.com</p>
                    <p>เวลาทำการ: จันทร์ - เสาร์ 08:00 - 18:00</p>
                </div>

                <div className="footer-section">
                    <h4>ติดตามเรา</h4>
                    <div className="social-icons">
                        <a href="#"><img src="/assets/facebook.png" alt="Facebook" /></a>
                        <a href="#"><img src="/assets/instagram.png" alt="Instagram" /></a>
                        <a href="#"><img src="/assets/line.png" alt="LINE" /></a>
                    </div>
                </div>
            </div>

            <div className="footer-bottom">
                <p>&copy; {new Date().getFullYear()} Fix & Serve. All Rights Reserved.</p>
            </div>
        </footer>
    );
}

export default Footer;
