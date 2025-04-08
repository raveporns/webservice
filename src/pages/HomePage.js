import React, { useState, useEffect } from 'react';
import './HomePage.css';
import Navbar from '../components/Navbar';
import Footer from '../components/Footer';

function HomePage() {
    const bannerImages = [
        "/assets/banner1.png",
        "/assets/banner2.png"
    ];

    const [currentIndex, setCurrentIndex] = useState(0);
    const [showCartConfirmation, setShowCartConfirmation] = useState(false);
    const [selectedItem, setSelectedItem] = useState(null);

    useEffect(() => {
        const interval = setInterval(() => {
            setCurrentIndex((prev) => (prev + 1) % bannerImages.length);
        }, 3000);
        return () => clearInterval(interval); // cleanup
    }, [bannerImages.length]);

    const goToPrev = () => {
        setCurrentIndex((prev) =>
            prev === 0 ? bannerImages.length - 1 : prev - 1
        );
    };

    const goToNext = () => {
        setCurrentIndex((prev) => (prev + 1) % bannerImages.length);
    };

    const handleAddToCart = (item) => {
        setSelectedItem(item);
        setShowCartConfirmation(true);
    };

    const confirmAddToCart = () => {
        alert(`สินค้าหมวด ${selectedItem.title} ได้รับการเพิ่มลงในตะกร้าแล้ว!`);
        setShowCartConfirmation(false);
    };

    return (
        <div className="homepage-container">
            <Navbar />

            {/* Promo Banner */}
            <div className="promo-banner">
                <div className="custom-slider">
                    <button 
                        className="slider-btn left" 
                        onClick={goToPrev} 
                        aria-label="Previous banner">
                        &#8249;
                    </button>
                    <img
                        src={bannerImages[currentIndex]}
                        alt={`โปรโมชั่น ${currentIndex + 1}`}
                        className="promo-image"
                    />
                    <button 
                        className="slider-btn right" 
                        onClick={goToNext} 
                        aria-label="Next banner">
                        &#8250;
                    </button>
                </div>
                <div className="slider-dots">
                    {bannerImages.map((_, i) => (
                        <span
                            key={i}
                            className={`dot ${i === currentIndex ? 'active' : ''}`}
                            onClick={() => setCurrentIndex(i)}
                        />
                    ))}
                </div>
            </div>

            {/* Main Menu */}
            <div className="main-menu">
                <button className="menu-button">
                    <img src="/assets/ruler.png" alt="บริการติดตั้ง" className="menu-icon" />
                    บริการติดตั้ง
                </button>
                <button className="menu-button">
                    <img src="/assets/wrench.png" alt="บริการซ่อมแซม" className="menu-icon" />
                    บริการซ่อมแซม
                </button>
            </div>

            {/* Popular Services */}
            <section className="popular-section">
                <h3>บริการยอดนิยม</h3>
                <div className="popular-scroll">
                    {[
                        { title: "ติดตั้งเครื่องใช้ไฟฟ้า", imgSrc: "/assets/electric-appliance.jpg" },
                        { title: "ติดตั้งระบบไฟฟ้า", imgSrc: "/assets/electrical-system.jpg" },
                        { title: "ติดตั้งระบบประปา", imgSrc: "/assets/water-system.jpg" },
                        { title: "ติดตั้งเฟอร์นิเจอร์", imgSrc: "/assets/furniture.jpg" },
                    ].map((item, i) => (
                        <div key={i} className="popular-item">
                            <div className="img-box">
                                <img src={item.imgSrc} alt={item.title} />
                            </div>
                            <p>{item.title}</p>
                            <button className="popular-button">ดูรายละเอียด</button>
                            <button 
                                className="add-to-cart-button" 
                                onClick={() => handleAddToCart(item)}>
                                ใส่ตะกร้า
                            </button>
                        </div>
                    ))}
                </div>
            </section>

            {/* Categories */}
            <section className="category-section">
                <h3>หมวดสินค้า</h3>
                <div className="category-scroll">
                    {[
                        { title: "เครื่องใช้ไฟฟ้า", imgSrc: "/assets/electric-appliance.jpg" },
                        { title: "ประปา", imgSrc: "/assets/category2.png" },
                        { title: "ไฟฟ้า", imgSrc: "/assets/category3.png" },
                        { title: "เครื่องมือวัด", imgSrc: "/assets/category4.png" },
                        { title: "เฟอร์นิเจอร์", imgSrc: "/assets/category5.png" },
                        { title: "งานช่างทั่วไป", imgSrc: "/assets/category6.png" }
                    ].map((item, i) => (
                        <div key={i} className="category-item">
                            <div className="img-box">
                                <img src={item.imgSrc} alt={item.title} />
                            </div>
                            <p>{item.title}</p>
                            <button className="category-button">ดูรายละเอียด</button>
                            <button 
                                className="add-to-cart-button" 
                                onClick={() => handleAddToCart(item)}>
                                ใส่ตะกร้า
                            </button>
                        </div>
                    ))}
                </div>
            </section>

            {/* Cart Confirmation */}
            {showCartConfirmation && (
                <div className="cart-confirmation">
                    <p>คุณต้องการเพิ่มบริการหรือสินค้านี้ "{selectedItem.title}" ลงในตะกร้าหรือไม่?</p>
                    <button onClick={confirmAddToCart}>ใช่</button>
                    <button onClick={() => setShowCartConfirmation(false)}>ไม่</button>
                </div>
            )}

            {/* FAQ Section */}
            <section className="faq-section">
                <h3>คำถามที่พบบ่อย</h3>
                <div className="faq-box">รวมคำถามและคำตอบที่พบบ่อย</div>
            </section>

            {/* Map Section */}
            <section className="map-section">
                <h3>ตำแหน่ง</h3>
                <img src="/assets/map.png" alt="แผนที่" className="map-img" />
            </section>

            <Footer />
        </div>
    );
}

export default HomePage;
