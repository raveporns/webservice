package models

type User struct {
    UserID   uint   `json:"userID"`
	Password string `json:"password,omitempty"` 
    Email    string `json:"email"`
    Role     string `json:"role"`
}
