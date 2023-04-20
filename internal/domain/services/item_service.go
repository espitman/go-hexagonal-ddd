package services

import (
	"errors"
	"github.com/espitman/go-hexagonal-ddd/internal/domain/models"
	"github.com/espitman/go-hexagonal-ddd/internal/domain/repositories"
)

type ItemService struct {
	itemRepository repositories.ItemRepository
}

func NewItemService(itemRepository repositories.ItemRepository) *ItemService {
	return &ItemService{itemRepository: itemRepository}
}

func (s *ItemService) AddItem(item *models.Item) (*models.Item, error) {

	if s.itemRepository.IsMemberOfTheList(item.ItemCode, item.ListId) {
		return nil, errors.New("item already exists in this list")
	}
	return s.itemRepository.Create(item)
}

func (s *ItemService) GetItemsByListID(listId string) ([]*models.Item, error) {
	return s.itemRepository.GetItemsByListID(listId)
}

func (s *ItemService) DeleteItem(id string) error {
	//_, err := s.itemRepository.GetByID(id)
	//if err != nil {
	//	return errors.New("item doesnt exist!")
	//}
	return s.itemRepository.Delete(id)
}
