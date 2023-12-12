package models
import (
 "github.com/jinzhu/gorm"
 "github.com/google/uuid"
 "time"
)
type Ticket struct {
 gorm.Model
 ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4():primaryKey:unique"`
 TicketUserId TicketUser 
 Source string `gorm:"json:source"`
 Type string `gorm:"json:type"`
 Category string `gorm:"json:category"`
 Subcategory string `gorm:"json:subcategory"`
 Description string `gorm:"json:description"`
 Priority string `gorm:"json:priority"`
 Tags string `gorm:"json:tags"`
 Data string `gorm:"json:data"`
 NotificationPreferences string `gorm:"json:notification_preferences"`
 Tat string `gorm:"json:tat"`
 ExpiryDate string `gorm:"json:expiry_date"`
 Status string `gorm:"json:status"`
 CreatedAt time.Time `gorm:"json:created_at"`
 UpdatedAt time.Time `gorm:"json:updated_at"`
}