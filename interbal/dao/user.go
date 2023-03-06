package dao

import (
	"blog-service/interbal/model"
)

func (d *Dao) GetUser(username, password string) (model.User, error) {
	user := &model.User{Username: username, Password: password}
	return user.Get(d.engine)
}
