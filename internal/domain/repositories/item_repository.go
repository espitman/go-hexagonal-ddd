package repositories

// implemented by mongodb item_repository

import (
	"github.com/espitman/go-hexagonal-ddd/internal/domain/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ItemRepository interface {
	Create(item *models.Item) (*models.Item, error)
	GetItemsByListID(id string) ([]*models.Item, error)
	Delete(id string) error
	IsMemberOfTheList(itemCode int64, listId primitive.ObjectID) bool
}
