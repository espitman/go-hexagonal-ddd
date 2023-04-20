package useCases

import "github.com/espitman/go-hexagonal-ddd/internal/domain/models"

type ListUseCase interface {
	AddList(list *models.List) (*models.List, error)
	GetListByID(id string) (*models.List, error)
	ListLists(userId int64) ([]*models.List, error)
	UpdateList(id string, list *models.List) (*models.List, error)
	DeleteList(id string) error
	GetListByName(name string) (*models.List, error)
}
