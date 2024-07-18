import { useState } from "react";
import { Routes, Route } from "react-router-dom";
import "../../css/WelcomeHome.css";
import Navbar from "../Common/Navbar";

function WelcomeHome() {
  return (
    <div className="welcome-home-main">
      <Navbar />
      <p>hi</p>
    </div>
  );
}

export default WelcomeHome;
