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
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type RegisterUserPayload struct {
	Username    string `json:"username" validate:"required,min=4,max=25"`
	Password    string `json:"password" validate:"required,password"`
	Firstname   string `json:"firstname" validate:"required,min=2,max=255"`
	Lastname    string `json:"lastname" validate:"required,min=2,max=255"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phonenum"`
	ImgLink     string `json:"imglink"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Message struct {
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy int       `json:"createdBy"` // user ID
	ChannelID int       `json:"channelId"` // channel ID
}

type Channel struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	CreatedBy   int       `json:"createdBy"` // user ID
	ImgLink     string    `json:"imgLink"`
}
