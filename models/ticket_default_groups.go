package models
import (
 	"github.com/jinzhu/gorm"
 	"github.com/google/uuid"
 	"time"
)
type TicketDefaultGroups struct {
 	gorm.Model
 	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4():primaryKey:unique"`
 	TicketType string `gorm:"json:ticket_type"`
 	GroupId Group `gorm:"type:uuid:group_id"`
 	Status string `gorm:"json:status"`
 	CreatedAt time.Time 
 	UpdatedAt time.Time 
}