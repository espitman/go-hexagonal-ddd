package useCases

import "git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/models"

type ItemUseCase interface {
	AddItem(item *models.Item) (*models.Item, error)
	GetItemsByListID(listId string) (*models.Item, error)
	DeleteItem(id string) error
}
