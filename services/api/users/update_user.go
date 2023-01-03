package users

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"log"
)

func UpdateUser(id string, name string) models.GoUser{
	db := config.GetDB()
	
	var user models.GoUser
	log.Print(id)
	db.Where("id = ?", id).Find(&user)
	
	user.Name = name
	db.Save(&user)

	return user
}