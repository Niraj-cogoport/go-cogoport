package users

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func CreateUser(user models.GoUser) models.GoUser {
	db := config.GetDB()
	// result := map[string]interface{}{}
	db.Create(&user)
	return user
}