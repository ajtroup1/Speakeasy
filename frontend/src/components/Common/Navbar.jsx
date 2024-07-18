import { useState } from "react";
import "../../css/App.css";

function Navbar() {
  return (
    <div className="navbar-main">
      <div className=""></div>
      <div className="navbar-title-container">
        <p>Speakeasy</p>
        <img
          src="https://cdn-icons-png.flaticon.com/512/2790/2790087.png"
          className="nav-hat-icon"
        />
      </div>
    </div>
  );
}

export default Navbar;
