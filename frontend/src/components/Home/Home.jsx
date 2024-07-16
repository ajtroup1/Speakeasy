import { useState } from "react";
import { Routes, Route } from "react-router-dom";
import "../../css/Home.css";
import ChannelsList from "./ChannelsList"
import Navbar from "./Navbar";
import Channel from "../Channel/Channel";

function Home() {

  return (
    <div className="home-main">
      <div className="left-home-main">
        <ChannelsList />
      </div>
      <div className="right-home-main">
        <Navbar />
        <img
          src="../../src/assets/speakeasy-logo.webp"
          className="right-home-background-img"
        />
        <div className="right-home-content">
          <Routes>
            <Route path="/channelname" element={<Channel />} />
          </Routes>
        </div>
      </div>
    </div>
  );
}

export default Home;
