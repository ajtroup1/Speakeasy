## Speakeasy
Speakeasy is a web app similar to Discord where users can join their friends in channels to send messages.

## API Documentation
The Speakeasy API uses Go Gorilla Mux to handle basic functions relating to users, messaging, and managing channels. The documentation is broken into these components accordingly.

The API runs on `http://localhost:8080/api/v1/`, which is defined in `api.go` via a subrouter.

# Users
User functionality contains the ability to:

- Register users 
    - Endpoint: `/register`
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
  
- Edit user information 
    - Endpoint: `/edit`
  
- Friend / unfriend a user 
    - Endpoint: `/friend`, `/unfriend`
  
- Block / unblock a user 
    - Endpoint: `/block`, `/unblock`
  
- Change a user's status depending on recent activity 
    - Endpoint: `/___`

# Channels
