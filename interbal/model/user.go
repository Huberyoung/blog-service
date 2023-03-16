package model

import "github.com/jinzhu/gorm"

type User struct {
	*Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetUser struct {
	Token string `json:"token"`
}

func (u User) TableName() string {
	return "blog_user"
}

func (u User) Get(db *gorm.DB) (User, error) {
	if err := db.Where("username = ? and password = ?", u.Username, u.Password).First(&u).Error; err != nil {
		return User{}, err
	}
	return u, nil
}
