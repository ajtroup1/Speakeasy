import { useState } from "react";
import "../../css/Channel.css";

function Channel() {
  const [username, setUsername] = useState("adamjtroup");
  const messages = [
    {
      id: 0,
      content: "Message 1",
      username: "User 1",
      fileAttachments: [],
      createdAt: "7/16/2024",
      userimg:
        "https://www.zdnet.com/a/img/resize/65c53a5f93470920c32524d907dc0f1554f9e767/2022/10/01/4dbd71ee-49de-4b9b-a78e-261b4d2c728c/dall-e-person-explaining-to-coworkers.png?auto=webp&width=1280",
    },
    {
      id: 1,
      content:
        "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sit amet mattis vulputate enim. Turpis nunc eget lorem dolor sed. Diam quis enim lobortis scelerisque fermentum dui faucibus. Tellus rutrum tellus pellentesque eu tincidunt tortor aliquam nulla. Orci ac auctor augue mauris augue neque gravida in fermentum. Cursus vitae congue mauris rhoncus aenean. Ullamcorper a lacus vestibulum sed. In fermentum posuere urna nec tincidunt. Dignissim cras tincidunt lobortis feugiat vivamus. At erat pellentesque adipiscing commodo elit at imperdiet dui accumsan. Et tortor consequat id porta nibh venenatis. Quis viverra nibh cras pulvinar mattis nunc sed blandit libero.Sit amet tellus cras adipiscing enim eu turpis egestas. Lectus proin nibh nisl condimentum id. Non nisi est sit amet facilisis magna etiam tempor. Aliquet bibendum enim facilisis gravida. Gravida in fermentum et sollicitudin ac. Nisi vitae suscipit tellus mauris a. Varius vel pharetra vel turpis nunc eget lorem dolor. Nam aliquam sem et tortor consequat. Posuere lorem ipsum dolor sit amet consectetur adipiscing. Ut morbi tincidunt augue interdum velit euismod in pellentesque massa. Aliquam ultrices sagittis orci a scelerisque. Tortor at auctor urna nunc id cursus metus aliquam eleifend. Nunc sed augue lacus viverra.Mattis molestie a iaculis at erat. Nisl purus in mollis nunc. Velit euismod in pellentesque massa placerat. Nisi lacus sed viverra tellus in hac habitasse platea dictumst. Malesuada fames ac turpis egestas integer eget. Leo in vitae turpis massa sed elementum tempus egestas. Pellentesque adipiscing commodo elit at imperdiet dui accumsan sit amet. Penatibus et magnis dis parturient montes. Etiam tempor orci eu lobortis elementum. Feugiat in ante metus dictum at. Amet volutpat consequat mauris nunc congue nisi vitae suscipit tellus. Mauris pellentesque pulvinar pellentesque habitant morbi tristique senectus. Convallis posuere morbi leo urna. Senectus et netus et malesuada. Quis risus sed vulputate odio ut enim blandit. Imperdiet massa tincidunt nunc pulvinar sapien et. Et odio pellentesque diam volutpat. Vitae et leo duis ut. Tincidunt tortor aliquam nulla facilisi. Ullamcorper eget nulla facilisi etiam dignissim diam.",
      username: "User 2",
      fileAttachments: [],
      createdAt: "7/16/2024",
      userimg:
        "https://www.zdnet.com/a/img/resize/65c53a5f93470920c32524d907dc0f1554f9e767/2022/10/01/4dbd71ee-49de-4b9b-a78e-261b4d2c728c/dall-e-person-explaining-to-coworkers.png?auto=webp&width=1280",
    },
    {
      id: 2,
      content: "Message 3",
      username: "User 3",
      fileAttachments: [],
      createdAt: "7/16/2024",
      userimg:
        "https://www.zdnet.com/a/img/resize/65c53a5f93470920c32524d907dc0f1554f9e767/2022/10/01/4dbd71ee-49de-4b9b-a78e-261b4d2c728c/dall-e-person-explaining-to-coworkers.png?auto=webp&width=1280",
    },
    {
      id: 3,
      content: "Message 4",
      username: "User 4",
      fileAttachments: [],
      createdAt: "7/16/2024",
      userimg:
        "https://www.zdnet.com/a/img/resize/65c53a5f93470920c32524d907dc0f1554f9e767/2022/10/01/4dbd71ee-49de-4b9b-a78e-261b4d2c728c/dall-e-person-explaining-to-coworkers.png?auto=webp&width=1280",
    },
    {
      id: 4,
      content: "Message 5",
      username: "User 5",
      fileAttachments: [],
      createdAt: "7/16/2024",
      userimg:
        "https://www.zdnet.com/a/img/resize/65c53a5f93470920c32524d907dc0f1554f9e767/2022/10/01/4dbd71ee-49de-4b9b-a78e-261b4d2c728c/dall-e-person-explaining-to-coworkers.png?auto=webp&width=1280",
    },
    {
      id: 5,
      content: "Message 6",
      username: "User 6",
      fileAttachments: [],
      createdAt: "7/16/2024",
      userimg:
        "https://www.zdnet.com/a/img/resize/65c53a5f93470920c32524d907dc0f1554f9e767/2022/10/01/4dbd71ee-49de-4b9b-a78e-261b4d2c728c/dall-e-person-explaining-to-coworkers.png?auto=webp&width=1280",
    },
    {
      id: 6,
      content: "Message 7",
      username: "User 7",
      fileAttachments: [],
      createdAt: "7/16/2024",
      userimg:
        "https://www.zdnet.com/a/img/resize/65c53a5f93470920c32524d907dc0f1554f9e767/2022/10/01/4dbd71ee-49de-4b9b-a78e-261b4d2c728c/dall-e-person-explaining-to-coworkers.png?auto=webp&width=1280",
    },
    {
      id: 7,
      content: "Message 8",
      username: "User 8",
      fileAttachments: [],
      createdAt: "7/16/2024",
      userimg:
        "https://www.zdnet.com/a/img/resize/65c53a5f93470920c32524d907dc0f1554f9e767/2022/10/01/4dbd71ee-49de-4b9b-a78e-261b4d2c728c/dall-e-person-explaining-to-coworkers.png?auto=webp&width=1280",
    },
    {
      id: 8,
      content: "Message 9",
      username: "User 9",
      fileAttachments: [],
      createdAt: "7/16/2024",
      userimg:
        "https://www.zdnet.com/a/img/resize/65c53a5f93470920c32524d907dc0f1554f9e767/2022/10/01/4dbd71ee-49de-4b9b-a78e-261b4d2c728c/dall-e-person-explaining-to-coworkers.png?auto=webp&width=1280",
    },
    {
      id: 9,
      content: "Message 10",
      username: "User 10",
      fileAttachments: [],
      createdAt: "7/16/2024",
      userimg:
        "https://www.zdnet.com/a/img/resize/65c53a5f93470920c32524d907dc0f1554f9e767/2022/10/01/4dbd71ee-49de-4b9b-a78e-261b4d2c728c/dall-e-person-explaining-to-coworkers.png?auto=webp&width=1280",
    },
    {
      id: 10,
      content: "Message 11",
      username: "User 11",
      fileAttachments: [],
      createdAt: "7/16/2024",
      userimg:
        "https://www.zdnet.com/a/img/resize/65c53a5f93470920c32524d907dc0f1554f9e767/2022/10/01/4dbd71ee-49de-4b9b-a78e-261b4d2c728c/dall-e-person-explaining-to-coworkers.png?auto=webp&width=1280",
    },
    {
      id: 11,
      content: "Message 12",
      username: "User 12",
      fileAttachments: [],
      createdAt: "7/16/2024",
      userimg:
        "https://www.zdnet.com/a/img/resize/65c53a5f93470920c32524d907dc0f1554f9e767/2022/10/01/4dbd71ee-49de-4b9b-a78e-261b4d2c728c/dall-e-person-explaining-to-coworkers.png?auto=webp&width=1280",
    },
    {
      id: 12,
      content: "Message 13",
      username: "User 13",
      fileAttachments: [],
      createdAt: "7/16/2024",
      userimg:
        "https://www.zdnet.com/a/img/resize/65c53a5f93470920c32524d907dc0f1554f9e767/2022/10/01/4dbd71ee-49de-4b9b-a78e-261b4d2c728c/dall-e-person-explaining-to-coworkers.png?auto=webp&width=1280",
    },
    {
      id: 13,
      content: "Message 14",
      username: "User 14",
      fileAttachments: [],
      createdAt: "7/16/2024",
      userimg:
        "https://www.zdnet.com/a/img/resize/65c53a5f93470920c32524d907dc0f1554f9e767/2022/10/01/4dbd71ee-49de-4b9b-a78e-261b4d2c728c/dall-e-person-explaining-to-coworkers.png?auto=webp&width=1280",
    },
    {
      id: 14,
      content: "Message 15",
      username: "User 15",
      fileAttachments: [],
      createdAt: "7/16/2024",
      userimg:
        "https://www.zdnet.com/a/img/resize/65c53a5f93470920c32524d907dc0f1554f9e767/2022/10/01/4dbd71ee-49de-4b9b-a78e-261b4d2c728c/dall-e-person-explaining-to-coworkers.png?auto=webp&width=1280",
    },
    {
      id: 15,
      content: "Message 16",
      username: "User 16",
      fileAttachments: [],
      createdAt: "7/16/2024",
      userimg:
        "https://www.zdnet.com/a/img/resize/65c53a5f93470920c32524d907dc0f1554f9e767/2022/10/01/4dbd71ee-49de-4b9b-a78e-261b4d2c728c/dall-e-person-explaining-to-coworkers.png?auto=webp&width=1280",
    },
    {
      id: 16,
      content: "Message 17",
      username: "User 17",
      fileAttachments: [],
      createdAt: "7/16/2024",
      userimg:
        "https://www.zdnet.com/a/img/resize/65c53a5f93470920c32524d907dc0f1554f9e767/2022/10/01/4dbd71ee-49de-4b9b-a78e-261b4d2c728c/dall-e-person-explaining-to-coworkers.png?auto=webp&width=1280",
    },
    {
      id: 17,
      content: "Message 18",
      username: "User 18",
      fileAttachments: [],
      createdAt: "7/16/2024",
      userimg:
        "https://www.zdnet.com/a/img/resize/65c53a5f93470920c32524d907dc0f1554f9e767/2022/10/01/4dbd71ee-49de-4b9b-a78e-261b4d2c728c/dall-e-person-explaining-to-coworkers.png?auto=webp&width=1280",
    },
    {
      id: 18,
      content: "Message 19",
      username: "adamjtroup",
      fileAttachments: [],
      createdAt: "7/16/2024",
      userimg:
        "https://www.shutterstock.com/image-photo/teamwork-meeting-tablet-business-people-600nw-2251938325.jpg",
    },
    {
      id: 19,
      content: "Message 20",
      username: "User 20",
      fileAttachments: [],
      createdAt: "7/16/2024",
      userimg:
        "https://www.zdnet.com/a/img/resize/65c53a5f93470920c32524d907dc0f1554f9e767/2022/10/01/4dbd71ee-49de-4b9b-a78e-261b4d2c728c/dall-e-person-explaining-to-coworkers.png?auto=webp&width=1280",
    },
  ];

  return (
    <div className="channel-main">
      <div className="channel-content-main">
        <div className="inner-channel-content">
          {messages.map((message, index) => {
            console.log(message.username, username)
            return (
              <>
              {message.username == username ? (
                <div key={index} className="my-indiv-message-channel-container">
                    <div className="channel-upper-msg-container">
                      <p></p>
                      <div className="channel-msg-profile-img-container">
                        <img
                          src={message.userimg}
                          className="channel-msg-profile-img"
                        />
                      </div>
                      <div className="channel-msg-username-container">
                        <p className="channel-msg-username">{message.username}</p>
                      </div>
                      <p className="msg-createdat">Sent {message.createdAt}</p>
                    </div>
                    <div className="channel-msg-content">
                      <p>{message.content}</p>
                    </div>
                  </div>
              ) : (

                  <div key={index} className="indiv-message-channel-container">
                    <div className="channel-upper-msg-container">
                      <p></p>
                      <div className="channel-msg-profile-img-container">
                        <img
                          src={message.userimg}
                          className="channel-msg-profile-img"
                        />
                      </div>
                      <div className="channel-msg-username-container">
                        <p className="channel-msg-username">{message.username}</p>
                      </div>
                      <p className="msg-createdat">Sent {message.createdAt}</p>
                    </div>
                    <div className="channel-msg-content">
                      <p>{message.content}</p>
                    </div>
                  </div>
              )}
              </>
            );
          })}
        </div>
      </div>
      <div className="messaging-container">
        <img
          src="https://cdn-icons-png.freepik.com/512/3756/3756616.png"
          className="attachment-icon"
        />
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
