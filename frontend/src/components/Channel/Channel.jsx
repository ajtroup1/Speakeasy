import { useState } from "react";
import "../../css/Channel.css";

function Channel() {
  return (
    <div className="channel-main">
        <div className="channel-content-main"></div>
      <div className="messaging-container">
        <img src="https://cdn-icons-png.freepik.com/512/3756/3756616.png" className="attachment-icon"/>
        <textarea
          className="message-input"
          placeholder="Type your message..."
        ></textarea>
        <img
          className="send-button"
          src="https://www.iconninja.com/files/148/341/227/telegram-send-chat-media-message-icon.svg"
        />
      </div>
    </div>
  );
}

export default Channel;
