package models
import (
 	"github.com/jinzhu/gorm"
 	"github.com/google/uuid"
 	"time"
)
type TicketTask struct {
 	gorm.Model
 	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4():primaryKey:unique"`
 	TicketId Ticket 
 	Title string `gorm:"json:title"`
 	CreatedByUserId uuid.UUID `gorm:"type:uuid:created_by_user_id"`
 	Status string `gorm:"json:status"`
 	CreatedAt time.Time 
 	UpdatedAt time.Time 
}