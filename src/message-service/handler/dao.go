package handler

import (
	"gomicrodemo1/common/model"
	"time"
)

func GetUserSessionByToken(token string) (model.User, error) {

	var (
		interval = 86400
		user     model.User
	)

	sql := `
select
  id,
  username,
  nickname
from user
where id = (
  select user_id
  from user_login_history
  where token = ?
        and logout = 0
        and date_add(create_time, interval ? second) > ?
  limit 1
)

`

	err := DB.Raw(sql, token, interval, time.Now()).Scan(&user).Error
	return user, err

}

func Login(user model.User) uint32 {

	var UserId struct {
		UserId uint32
	}

	sql := `
select user.id as user_id
from user
where user.username = ?
      and user.password = ?
`

	DB.Raw(sql, user.Username, user.Password).Scan(&UserId)
	return UserId.UserId

}

func GetUsernamesByIds(ids []uint32) ([]model.User, error) {

	users := make([]model.User, 0)

	err := DB.Select("id, nickname").Where(" id in (?) ", ids).Find(&users).Error

	return users, err

}
