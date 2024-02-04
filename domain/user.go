package domain

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             uuid.UUID `gorm:"type:uuid;primary_key;"`
	Username       string    `gorm:"unique"`
	Password       string
	DungeonsJoined []Dungeon `gorm:"many2many:user_dungeons"`
	DungeonsOwned  []Dungeon
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.NewV4()

	return nil
}
