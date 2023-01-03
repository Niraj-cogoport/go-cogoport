package models

import (
	"github.com/jinzhu/gorm"
	"github.com/google/uuid"
	"time"
)

type Group struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4():primaryKey:unique"`
	Name string 
	Tags string
	Status string 
	CreatedAt time.Time 
	UpdatedAt time.Time 
}