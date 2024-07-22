## Speakeasy
Speakeasy is a web app similar to Discord where users can join their friends in channels to send messages.

## API Documentation
*****ABOUT MAKEFILE
The Speakeasy API uses Go Gorilla Mux to handle basic functions relating to users, messaging, and managing channels. The documentation is broken into these components accordingly.

The API runs on `http://localhost:8080/api/v1/`, which is defined in `api.go` via a subrouter.

# Users
User functionality contains the ability to:

- Return all users
    - Endpoint: `/users`
    - Method: GET
    - Expects no payload
    - Returns a 200 on successful execution

- Return user by id
    - Endpoint: `/users/{id}`
    - Method: GET
    - Expects no payload (just query params)
    - Returns a 200 on successful execution

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
  
- Change a user's status depending on recent activity 
    - Endpoint: `/___`

` 

# Channels
