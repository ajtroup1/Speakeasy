import { useState } from "react";
import { Routes, Route } from 'react-router-dom'
import "../css/App.css";
import Home from "./Home/Home";
import Login from "./Login/Login";
import WelcomeHome from "./WelcomeHome/WelcomeHome"

function App() {
  return (
    <div className="main">
      <Routes>
        <Route path="/username/*" element={<Home />} />
        <Route path="/" element={<WelcomeHome />} />
      </Routes>
    </div>
  );
}

export default App;
