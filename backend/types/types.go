package types

import (
	"time"
)

type User struct {
	ID          int       `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phonenum"`
	ImgLink     string    `json:"imglink"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	TextNotifications bool `json:"textNotifications"`
	EmailNotifications bool `json:"emailNotifications"`
}

type UserStore interface {
	GetAllUsers() ([]*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
	EditUser(User) error
}

type FriendStore interface {
	FriendUser(sendID, receiveID uint) error
	UnfriendUser(sendID, receiveID uint) error
	Refriend(sendID, receiveID uint) error
	GetFriendshipByIDs(sendID, receiveID uint) (bool, error)
}

type BlockStore interface {
	BlockUser(sendID, receiveID uint) error
	UnblockUser(sendID, receiveID uint) error
	GetBlockByIDs(sendID, receiveID uint) (bool, error)
}

type MessageStore interface {
	CreateMessage(Message) error
}

type ChannelStore interface {
	// GetChannelByID(id int) error
	CreateChannel(Channel) error
}

type RegisterUserPayload struct {
	Username    string `json:"username" validate:"required,min=4,max=25"`
	Password    string `json:"password" validate:"required,password"`
	Firstname   string `json:"firstname" validate:"required,min=2,max=255"`
	Lastname    string `json:"lastname" validate:"required,min=2,max=255"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phonenum"`
	ImgLink     string `json:"imglink"`
	TextNotifications bool `json:"textNotifications"`
	EmailNotifications bool `json:"emailNotifications"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type EditUserPayload struct {
	ID int
	Username string
	Firstname string
	Lastname string
	Email string
	PhoneNumber string
	ImgLink string
}

type FriendPayload struct {
    SendID    uint `json:"sendID" validate:"required,gt=0"`
    ReceiveID uint `json:"receiveID" validate:"required,gt=0,nefield=SendID"`
}

type Message struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy int       `json:"createdBy"` // user ID
	ChannelD  int       `json:"channeId"`  // channel ID
}

type CreateMessagePayload struct {
	Content   string `json:"content" validate:"required,min=1"`
	CreatedBy int    `json:"createdBy" validate:"required"` // user ID
	ChannelD int    `json:"channelId" validate:"required"`  // channel ID
}

type Channel struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	CreatedBy   int       `json:"createdBy"` // user ID
	ImgLink     string    `json:"imgLink"`
}

type CreateChannelPayload struct {
	Name        string    `json:"name" validate:"required,min=1"`
	Description string    `json:"description"`
	CreatedBy   int       `json:"createdBy" validate:"required"` // user ID
	ImgLink     string    `json:"imgLink"`
}