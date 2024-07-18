import { useState } from "react";
import "../../css/Home.css";

function Navbar() {
  return (
    <div className="home-navbar-main">
      <div className="home-navbar-title-container">
        <p>Speakeasy</p>
        <img src="https://cdn-icons-png.flaticon.com/512/2790/2790087.png" className="home-nav-hat-icon"/>
      </div>
    </div>
  );
}

export default Navbar;
