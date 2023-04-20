package useCases

import "github.com/espitman/go-hexagonal-ddd/internal/domain/models"

type ItemUseCase interface {
	AddItem(item *models.Item) (*models.Item, error)
	GetItemsByListID(listId string) ([]*models.Item, error)
	DeleteItem(id string) error
}
