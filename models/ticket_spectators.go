package models
import (
    "github.com/jinzhu/gorm"
    "github.com/google/uuid"
    "time"
)
type TicketSpectator struct {
    gorm.Model
    ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4():primaryKey:unique"`
    TicketId Ticket 
    TicketUserId TicketUser 
    Status string `gorm:"json:source"`
    CreatedAt time.Time 
    UpdatedAt time.Time 
}