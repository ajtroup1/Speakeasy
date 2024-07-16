import React from "react";
import "../../css/Home.css";

function Channels() {
  const channels = [
    { id: 0, name: "Channel 1" },
    { id: 1, name: "Channel 2" },
    { id: 2, name: "Channel 3" },
    { id: 3, name: "Channel 4" },
    { id: 4, name: "Channel 5" },
    { id: 5, name: "Channel 6" },
    { id: 6, name: "Channel 7" },
    { id: 7, name: "Channel 8" },
  ];

  return (
    <div className="channels-main">
      {channels.map((channel) => (
        <div key={channel.id} className="channel-container">
          <div className="channel-img-container"></div>
          <p>{channel.name}</p>
        </div>
      ))}
    </div>
  );
}

export default Channels;
