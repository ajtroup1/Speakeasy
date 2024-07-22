## Speakeasy
Speakeasy is a web app similar to Discord where users can join their friends in channels to send messages.

## API Documentation
When navigated into the Speakeasy folder, cd into backend to run commands from the Makefile. The Makefile contains standard functions like run, test, ..., so reference this for what you can do with it
```
build:
	@go build -o bin/speakeasy cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/speakeasy

format:
	@go fmt ./...

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down
```
The Speakeasy API uses GO's Gorilla Mux to handle basic functions relating to users, messaging, and managing channels. The documentation is broken into these components accordingly.

The API runs on `http://localhost:8080/api/v1/`, which is defined in `api.go` via a subrouter.

# Users
User functionality contains the ability to:

- Return all users
    - Endpoint: `/users`
    - Method: GET
    - Expects no payload
    - Returns a 200 and User objects on successful execution:
        ```go
        type User struct {
            ID                 int       `json:"id"`
            Username           string    `json:"username"`
            Password           string    `json:"password"`
            Firstname          string    `json:"firstname"`
            Lastname           string    `json:"lastname"`
            Email              string    `json:"email"`
            PhoneNumber        string    `json:"phonenum"`
            ImgLink            string    `json:"imglink"`
            Status             int       `json:"status"`
            CreatedAt          time.Time `json:"createdAt"`
            TextNotifications  bool      `json:"textNotifications"`
            EmailNotifications bool      `json:"emailNotifications"`
        }
        ```

- Return user by id
    - Endpoint: `/users/{id}`
    - Method: GET
    - Expects no payload (just query params)
    - Returns a 200 and a User object on successful execution:
        ```go
        type User struct {
            ID                 int       `json:"id"`
            Username           string    `json:"username"`
            Password           string    `json:"password"`
            Firstname          string    `json:"firstname"`
            Lastname           string    `json:"lastname"`
            Email              string    `json:"email"`
            PhoneNumber        string    `json:"phonenum"`
            ImgLink            string    `json:"imglink"`
            Status             int       `json:"status"`
            CreatedAt          time.Time `json:"createdAt"`
            TextNotifications  bool      `json:"textNotifications"`
            EmailNotifications bool      `json:"emailNotifications"`
        }
        ```

- Search for a user
    - Endpoint `users/search/{searchParam}`
    - Method: GET
    - Expects no payload (just query params) 
    - Returns a 200 and a User object on successful execution:
        ```go
        type User struct {
            ID                 int       `json:"id"`
            Username           string    `json:"username"`
            Password           string    `json:"password"`
            Firstname          string    `json:"firstname"`
            Lastname           string    `json:"lastname"`
            Email              string    `json:"email"`
            PhoneNumber        string    `json:"phonenum"`
            ImgLink            string    `json:"imglink"`
            Status             int       `json:"status"`
            CreatedAt          time.Time `json:"createdAt"`
            TextNotifications  bool      `json:"textNotifications"`
            EmailNotifications bool      `json:"emailNotifications"`
        }
        ```

- Register users 
    - Endpoint: `/register`
    - Method: POST
    - Expects a payload:
        ```go
        type RegisterUserPayload struct {
            Username    string `json:"username" validate:"required,min=4,max=25"`
            Password    string `json:"password" validate:"required,password"`
            Firstname   string `json:"firstname" validate:"required,min=2,max=255"`
            Lastname    string `json:"lastname" validate:"required,min=2,max=255"`
            Email       string `json:"email" validate:"required,email"`
            PhoneNumber string `json:"phonenum"`
            ImgLink     string `json:"imglink"`
        }
        ```
    - Returns a 201 on successful execution

- Log in 
    - Endpoint: `/login`
    - Method: POST
    - Expects a payload:
        ```go
        type LoginUserPayload struct {
            Email    string `json:"email" validate:"required,email"`
            Password string `json:"password" validate:"required"`
        }
        ```
    - Returns a 200 on successful execution
  
- Edit user information 
    - Endpoint: `/edit`
    - Method: PUT
    - Expects a payload:
        ```go
        type EditUserPayload struct {
            ID          int
            Username    string
            Firstname   string
            Lastname    string
            Email       string
            PhoneNumber string
            ImgLink     string
        }
        ```
    - Returns a 200 on successful execution

- Change user password
    - Endpoint: `/changepassword`
    - Method: PUT
    - Expects a payload:
        ```go
        type ChangePasswordPayload struct {
            UserID              uint   `json:"userID" validate:"required"`
            CurrentPassword     string `json:"currentPassword" validate:"required"`
            NewPassword         string `json:"newPassword" validate:"required,password"`
            ConfirmNewPassword  string `json:"confirmPassword" validate:"required"`
        }
        ```
    - Returns a 200 on successful execution

- Friend / unfriend a user 
    - Endpoint: `/friend`, `/unfriend`
    - Method: POST
    - Expects a payload:
        ```go
        type FriendPayload struct {
            SendID    uint `json:"sendID" validate:"required,gt=0"`
            ReceiveID uint `json:"receiveID" validate:"required,gt=0,nefield=SendID"`
        }
        ```
    - Returns a 200 on successful execution

- Get a user's friends by user ID
    - Endpoint: `/friends/{userID}`
    - Method: GET
    - Expects no payload (just query params) 
    - Returns a 200 on successful execution

  
- Block / unblock a user 
    - Endpoint: `/block`, `/unblock`
    - Method: POST
    - Expects a payload:
        ```go
        type FriendPayload struct {
            SendID    uint `json:"sendID" validate:"required,gt=0"`
            ReceiveID uint `json:"receiveID" validate:"required,gt=0,nefield=SendID"`
        }
        ```
        - Block and friend payload are identical so I use FriendPayload for both
    - Returns a 200 on successful execution

- Get a user's blocks by user ID
    - Endpoint: `/blocks/{userID}`
    - Method: GET
    - Expects no payload (just query params) 
    - Returns a 200 on successful execution
  
- Change a user's status depending on recent activity 
    - Endpoint: `/___`

- Deactivate a user's account
    - Endpoint: `/deactivateaccount`
    - Method: DELETE
    - Expects a payload:
        ```go
        type DeactivateAccountPayload struct {
            UserID uint
            ConfirmPassword string
        }
        ```
    - Returns a 200 on successful execution

# Channels

- Return all channels
    - Endpoint: `/channels`
    - Method: GET
    - Expects no payload
    - Returns a 200 on successful execution

- Return user by id
    - Endpoint: `/channels/{id}`
    - Method: GET
    - Expects no payload (just query params)
    - Returns a 200 on successful execution

- Create channel 
    - Endpoint: `/createchannel`
    - Method: POST
    - Expects a payload:
        ```go
        type CreateChannelPayload struct {
            Name        string `json:"name" validate:"required,min=1"`
            Description string `json:"description"`
            CreatedBy   int    `json:"createdBy" validate:"required"` // user ID
            ImgLink     string `json:"imgLink"`
        }
        ```
    - Returns a 201 on successful execution

- Edit channel information 
    - Endpoint: `/editchannel`
    - Method: PUT
    - Expects a payload:
        ```go
        type EditChannelPayload struct {
            Name        string    `json:"name"`
            Description string    `json:"description"`
            ImgLink     string    `json:"imgLink"`
        }
        ```
    - Returns a 200 on successful execution

- Toggle channel private
    - Endpoint: `/privatechanneltoggle`
    - Method: POST
    - Expects a payload:
        ```go
        type ToggleChannelPrivatePayload struct {
            ChannelD uint `json:"channelD" validate:"required"`
            UserID uint `json:"userID" validate:"required"`
        }
        ```
    - Returns a 200 on successful execution

- Manage channel members
    - Endpoint: `/channelmembers/{channelID}`
    - Methods: `GET`, `POST`, `DELETE`
    - GET method:
        - Expects no payload (just query params)
        - Returns a 200 on successful execution
    - POST method:
        - Expects a payload
        ```go
        type AddChannelMemberPayload struct {
            UserID uint `json:"userID"`
            ChannelID uint `json:"channelID"`
            AddingUserID uint `json:"addingUserID"`
        }
        ```
        - Returns a 200 on successful execution
    - DELETE method:
        - Expects a payload:
            ```go
            type AddChannelMemberPayload struct {
                UserID uint `json:"userID"`
                ChannelID uint `json:"channelID"`
                AddingUserID uint `json:"addingUserID"`
            }
            ```
            - Block and friend payload are identical so I use FriendPayload for both
        - Returns a 200 on successful execution

- Deactivate a channel
    - Endpoint: `/deactivatechannel`
    - Method: DELETE
    - Expects a payload:
        ```go
        type DeactivateChannelPayload struct {
            UserID uint
        }
        ```
    - Returns a 200 on successful execution


# Messages

- Return messages by channel id
    - Endpoint: `/messages/{channelID}`
    - Method: GET
    - Expects no payload (just query params)
    - Returns a 200 on successful execution

- Create message 
    - Endpoint: `/createmessage/{channelID}`
    - Method: POST
    - Expects a payload:
        ```go
        type CreateMessagePayload struct {
            Content   string `json:"content" validate:"required,min=1"`
            CreatedBy int    `json:"createdBy" validate:"required"` // user ID
            ChannellD  int    `json:"channelId" validate:"required"` // channel ID
        }
        ```
    - Returns a 201 on successful execution

- Delete a message
    - Endpoint: `/deletemessage`
    - Method: DELETE
    - Expects a payload:
        ```go
        type DeleteMessagePayload struct {
            UserID uint
            MessageID uint
        }
        ```
    - Returns a 200 on successful execution