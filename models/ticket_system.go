package models

import (

	"github.com/jinzhu/gorm"
	// "github.com/tejas-cogo/go-cogoport/config"
	"time"
	"github.com/google/uuid"
)

var db *gorm.DB

//Model is sample of common table structure
type Model struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4():primaryKey:unique"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

func init() {

	// config.GetDB()
	//Printing query
	// db.LogMode(true)

	//Automatically create migration as per model
	db.Debug().AutoMigrate(
		&GroupMember{},
	)
}

//GetDB function return the instance of db
func GetDB() *gorm.DB {
	return db
}