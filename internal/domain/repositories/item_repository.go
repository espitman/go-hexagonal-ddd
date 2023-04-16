package repositories

// implemented by mongodb item_repository

import (
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ItemRepository interface {
	Create(item *models.Item) (*models.Item, error)
	GetItemsByListID(id string) ([]*models.Item, error)
	Delete(id string) error
	IsMemberOfTheList(itemCode int64, listId primitive.ObjectID) bool
}
