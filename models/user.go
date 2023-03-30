package models

type User struct {
	Id       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	FullName string `json:"full_name" gorm:"type: varchar(100)"`
	Username string `json:"username" gorm:"unique; type: varchar(100)"`
	Email    string `json:"email" gorm:"unique; type: varchar(100)"`
	Password string `json:"password" gorm:"type: varchar(255)"`
	Role     string `json:"role" gorm:"type: varchar(20)"`
	Gender   string `json:"gender" gorm:"type: varchar(10)"`
	Phone    string `json:"phone" gorm:"unique; type: varchar(20)"`
	Address  string `json:"address" gorm:"type: text"`
	Photo    string `json:"photo" gorm:"type:varchar(255)"`
}

type CheckAuthResponse struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
}

func (UserResponse) TableName() string {
	return "users"
}
