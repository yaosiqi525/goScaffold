package services

import (
	"goScaffold/database"
)

type user struct{}

func (imp *user) GetUserInfo(id int64) (user database.User) {
	database.DBUtils.ID(id).GetFirst(&user)
	return user
}

var User user
