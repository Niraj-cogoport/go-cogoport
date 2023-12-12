package models
import (
 	"github.com/jinzhu/gorm"
 	"github.com/google/uuid"
 	"time"
)
type TicketDefaultTypes struct {
 	gorm.Model
 	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4():primaryKey:unique"`
 	TicketType string `gorm:"json:ticket_type"`
 	AdditionalOptions string `gorm:"json:additional_optionals"`
 	Status string `gorm:"json:status"`
 	CreatedAt time.Time 
 	UpdatedAt time.Time 
}