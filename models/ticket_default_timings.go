package models
import (
 	"github.com/jinzhu/gorm"
 	"github.com/google/uuid"
 	"time"
)
type TicketDefaultTimings struct {
 	gorm.Model
 	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4():primaryKey:unique"`
 	TicketType string `gorm:"json:ticket_type"`
 	TicketPriority string `gorm:"json:ticket_priority"`
 	ExpiryDuration time.Time
 	Tat time.Time
 	Conditions string `gorm:"json:conditions"`
 	Status string `gorm:"json:status"`
 	CreatedAt time.Time 
 	UpdatedAt time.Time 
}