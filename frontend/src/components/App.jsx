import { useState } from "react";
import { Routes, Route } from 'react-router-dom'
import "../css/App.css";
import Home from "./Home/Home";

function App() {
  return (
    <div className="main">
      <Routes>
        <Route path="/username/*" element={<Home />} />
      </Routes>
    </div>
  );
}

export default App;
