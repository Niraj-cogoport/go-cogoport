package users

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteUser(id string) models.GoUser {
	db := config.GetDB()
	var user models.GoUser
	db.Where("id = ?", id).Delete(&user)
	return user
}