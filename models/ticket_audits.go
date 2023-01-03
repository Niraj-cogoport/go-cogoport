package models
import (
    "github.com/jinzhu/gorm"
    "github.com/google/uuid"
    "time"
)
type TicketAudits struct {
    gorm.Model
    ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4():primaryKey:unique"`
    Object string `gorm:"json:object"`
    ObjectId uuid.UUID `gorm:"type:uuid:object_id"`
	Action string `gorm:"json:action"`
	Data string `gorm:"json:data"`
    Status string `gorm:"json:source"`
    CreatedAt time.Time 
    UpdatedAt time.Time 
}