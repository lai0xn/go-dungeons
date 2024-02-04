package repositories

import (
	"errors"
	"fmt"

	"github.com/jn0x/reddigo/domain"
	"github.com/jn0x/reddigo/storage"
	uuid "github.com/satori/go.uuid"
)

type DungeonRepository interface {
	CreateDungeon(dungeon *domain.Dungeon) error
	DeleteDungeon(id uuid.UUID) error
	GetDungeonByID(id uuid.UUID) (domain.Dungeon, error)
	SearchDungeon(name string) ([]domain.Dungeon, error)
}

type dungeonRepository struct{}

func NewDungeonRepository() DungeonRepository {
	return DungeonRepository(&dungeonRepository{})
}

func (r *dungeonRepository) CreateDungeon(dungeon *domain.Dungeon) error {
	result := storage.DB.Create(&dungeon)
	if result.Error != nil {
		return errors.New("Dungeon with this name already exists")
	}
	return nil
}

func (r *dungeonRepository) DeleteDungeon(id uuid.UUID) error {
	storage.DB.Delete(&domain.Dungeon{}, id)
	return nil
}

func (r *dungeonRepository) GetDungeonByID(id uuid.UUID) (domain.Dungeon, error) {
	var dungeon domain.Dungeon
	storage.DB.First(&dungeon, "id = ?", id)
	return dungeon, nil
}

func (r *dungeonRepository) SearchDungeon(name string) ([]domain.Dungeon, error) {
	var dungeons []domain.Dungeon
	storage.DB.Find(&dungeons, "name LIKE ? ?", fmt.Sprintf("%%s%s", name))

	return dungeons, nil
}
