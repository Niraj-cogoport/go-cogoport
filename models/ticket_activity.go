package models
import (
 	"github.com/jinzhu/gorm"
 	"github.com/google/uuid"
 	"time"
)
type TicketActivity struct {
 	gorm.Model
 	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4():primaryKey:unique"`
 	TicketId Ticket
 	Type string 
 	TicketUserId TicketUser 
 	UserType string 
 	Description string 
 	Data string
 	IsRead bool 
 	Status string 
 	CreatedAt time.Time 
 	UpdatedAt time.Time 
}