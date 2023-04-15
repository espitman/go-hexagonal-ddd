package appServices

import (
	"fmt"
	appModel "git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/app/models"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/models"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/useCases"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ItemService struct {
	itemUseCases useCases.ItemUseCase
	listUseCase  useCases.ListUseCase
}

func NewItemService(
	itemUseCases useCases.ItemUseCase, listUseCase useCases.ListUseCase) *ItemService {
	return &ItemService{
		itemUseCases: itemUseCases,
		listUseCase:  listUseCase,
	}
}

func (s *ItemService) CreateItem(newItem *appModel.NewItem) (*appModel.Item, error) {

	_, err := s.listUseCase.GetListByID(newItem.ListId)
	if err != nil {
		return nil, fmt.Errorf("failed to add item: %w", err)
	}

	docListId, _ := primitive.ObjectIDFromHex(newItem.ListId)
	item := &models.Item{
		ListId:   docListId,
		ItemCode: newItem.ItemCode,
	}

	item, err = s.itemUseCases.AddItem(item)
	if err != nil {
		return nil, fmt.Errorf("failed to create item: %w", err)
	}
	appItem := &appModel.Item{
		ID:        item.ID,
		ListId:    newItem.ListId,
		ItemCode:  newItem.ItemCode,
		CreatedAt: item.CreatedAt,
	}
	return appItem, nil
}

func (s *ItemService) GetItemsByListID(listId string) (*appModel.Item, error) {
	item, err := s.itemUseCases.GetItemsByListID(listId)
	if err != nil {
		return nil, fmt.Errorf("failed to get item by ID %s: %w", listId, err)
	}

	if item == nil {
		return nil, nil
	}

	appItem := &appModel.Item{
		ID:     item.ID,
		ListId: listId,
	}

	return appItem, nil
}

func (s *ItemService) DeleteItem(id string) error {
	err := s.itemUseCases.DeleteItem(id)
	if err != nil {
		return fmt.Errorf("failed to delete item with ID %s: %w", id, err)
	}
	return nil
}
