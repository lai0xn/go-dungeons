package service

import (
	"github.com/jn0x/reddigo/domain"
	"github.com/jn0x/reddigo/http/requests"
	"github.com/jn0x/reddigo/repositories"
	"github.com/jn0x/reddigo/storage"
	uuid "github.com/satori/go.uuid"
)

type DungeonService interface {
	CreateDungeon(dungeon requests.DungeonReq) error
	GetDungeonByID(id uuid.UUID) *domain.Dungeon
	SearchDungeon(name string) []domain.Dungeon
	DeleteDungeon(id uuid.UUID)
	JoinDungeon(dunID uuid.UUID, userID uuid.UUID) error
}

type dungeonService struct {
	repo    repositories.DungeonRepository
	userepo repositories.UserRepository
}

func (s *dungeonService) CreateDungeon(dungeon requests.DungeonReq) error {
	s.repo.CreateDungeon(&domain.Dungeon{
		Name:        dungeon.Name,
		Description: dungeon.Description,
	})
	return nil
}

func (s *dungeonService) GetDungeonByID(id uuid.UUID) *domain.Dungeon {
	dungeon, err := s.repo.GetDungeonByID(id)
	if err != nil {
		return nil
	}
	return &dungeon
}

func (s *dungeonService) SearchDungeon(name string) []domain.Dungeon {
	dungeons, err := s.repo.SearchDungeon(name)
	if err != nil {
		return nil
	}
	return dungeons
}

func (s *dungeonService) DeleteDungeon(id uuid.UUID) {
	s.repo.DeleteDungeon(id)
}

func (s *dungeonService) JoinDungeon(dunID uuid.UUID, userID uuid.UUID) error {
	dungeon, err := s.repo.GetDungeonByID(dunID)
	if err != nil {
		return err
	}
	user := s.userepo.GetUser(userID)
	storage.DB.Model(&user).Association("Dungeons").Append(dungeon)
	return nil
}

func NewDungeonService() DungeonService {
	return &dungeonService{
		repo:    repositories.NewDungeonRepository(),
		userepo: repositories.NewUserRepository(),
	}
}
