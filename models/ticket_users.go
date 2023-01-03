package models

import (
	"github.com/jinzhu/gorm"
	"github.com/google/uuid"
	"time"
)

type TicketUser struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4():primaryKey:unique"`
	Name string `gorm:"json:name"`
	Email string `gorm:"json:email"`
	MobileNumber string `gorm:"json:mobile_number"`
	RoleIds Role
	Source string `gorm:"json:source"`
	Type string `gorm:"json:type"`
	Status string `gorm:"json:status"`
	CreatedAt time.Time `gorm:"json:created_at"`
	UpdatedAt time.Time `gorm:"json:updated_at"`
}