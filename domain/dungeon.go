package domain

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Dungeon struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name        string    `gorm:"unique"`
	Description string
	Members     []User `gorm:"many2many:user_dungeons"`
	UserID      uuid.UUID
}

func (dungeon *Dungeon) BeforeCreate(tx *gorm.DB) error {
	dungeon.ID = uuid.NewV4()

	return nil
}
