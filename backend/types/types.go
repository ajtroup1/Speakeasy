package types

import "time"

type User struct {
    ID          uint      `json:"id"`
    Username    string    `json:"username"`
    Firstname   string    `json:"firstname"`
    Lastname    string    `json:"lastname"`
    Email       string    `json:"email"`
    PhoneNumber string    `json:"phoneNumber"`
    Password    string    `json:"password"`
    Status      int       `json:"status"`
    Imglink     string    `json:"imglink"`
    CreatedAt   time.Time `json:"createdAt"`
}

type FriendList struct {
    ID        uint   `json:"id"`
    CreatedBy User   `json:"createdBy"`
    Members   []User `json:"members"`
}

type Channel struct {
    ID          int       `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Imglink     string    `json:"imglink"`
    CreatedAt   time.Time `json:"createdAt"`
    CreatedBy   User      `json:"createdBy"`
    Members     []User    `json:"members"`
}

type Message struct {
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"createdAt"`
    CreatedBy uint      `json:"createdBy"`
}
