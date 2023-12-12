package models

import (
	"github.com/jinzhu/gorm"
	"github.com/google/uuid"
	"time"
	
)

type GroupMember struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4():primaryKey:unique"`
	GroupId Group
	TicketUserId TicketUser
	HierarchyLevel uint
	Status string
	CreatedAt time.Time 
	UpdatedAt time.Time 
}

func (gm *GroupMember) TableName() string {
	return "group_member"
}