package object

import (
	"errors"
	"time"
)

type User struct {
	Username   string    `xorm:"'username'" json:"username"`
	Password   string    `xorm:"'password'" json:"password"`
	UserGroup  string    `xorm:"'user_group'" json:"user_group"`
	CreateTime time.Time `xorm:"'create_time'" json:"create_time"`
}

func (u *User) Authenticate() error {
	has, err := Engine.Where("username = ? AND password = ?", u.Username, u.Password).Get(u)
	if err != nil {
		return err
	}
	if !has {
		return errors.New("invalid username or password")
	}
	return nil
}
