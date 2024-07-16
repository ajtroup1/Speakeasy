import { useState } from "react";
import "../../css/Home.css";
import ChannelsList from "./ChannelsList"
import Navbar from "./Navbar";

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
      </div>
    </div>
  );
}

export default Home;
