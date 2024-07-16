import React from "react";
import "../../css/Home.css";

function Channels() {
  const channels = [
    {
      id: 0,
      name: "Channel 1",
      imglink:
        "https://www.adobe.com/content/dam/cc/us/en/creativecloud/photography/discover/stock-photography/thumbnail.jpeg",
    },
    { id: 1, name: "Channel 2" },
    {
      id: 2,
      name: "Channel 3",
      imglink:
        "https://www.outboundengine.com/wp-content/uploads/2014/04/Top-10-Worst-Stock-Photos-for-Your-Marketing.jpg",
    },
    { id: 3, name: "Channel 4" },
    { id: 4, name: "Channel 5" },
    {
      id: 5,
      name: "Channel 6",
      imglink:
        "https://images01.military.com/sites/default/files/styles/full/public/2019-09/MightyStocklead1200.jpg",
    },
    { id: 6, name: "Channel 7" },
    {
      id: 7,
      name: "Channel 8",
      imglink:
        "https://images.pexels.com/photos/3314294/pexels-photo-3314294.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500",
    },
  ];

  return (
    <div className="channels-main">
      {channels.map((channel) => (
        <div
          key={channel.id}
          className="channel-container"
          style={{
            backgroundImage: channel.imglink
              ? `url(${channel.imglink})`
              : "none",
          }}
        >
          <div className="channel-img-container"></div>
          {!channel.imglink && <p>{channel.name}</p>}
        </div>
      ))}
    </div>
  );
}

export default Channels;
