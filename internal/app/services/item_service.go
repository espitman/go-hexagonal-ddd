package appServices

import (
	"fmt"
	appModel "github.com/espitman/go-hexagonal-ddd/internal/app/models"
	"github.com/espitman/go-hexagonal-ddd/internal/domain/models"
	"github.com/espitman/go-hexagonal-ddd/internal/domain/useCases"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ItemService struct {
	itemUseCases useCases.ItemUseCase
	listUseCase  useCases.ListUseCase
	teamUseCase  useCases.TeamUseCase
}

func NewItemService(
	itemUseCases useCases.ItemUseCase,
	listUseCase useCases.ListUseCase,
	teamUseCase useCases.TeamUseCase,
) *ItemService {
	return &ItemService{
		itemUseCases: itemUseCases,
		listUseCase:  listUseCase,
		teamUseCase:  teamUseCase,
	}
}

func (s *ItemService) CreateItem(newItem *appModel.NewItem) (*appModel.Item, error) {

	_, err := s.listUseCase.GetListByID(newItem.ListId)
	if err != nil {
		return nil, fmt.Errorf("failed to add item: %w", err)
	}

	docListId, _ := primitive.ObjectIDFromHex(newItem.ListId)

	_, err = s.teamUseCase.GetTeamByID(newItem.ItemCode)
	if err != nil {
		return nil, fmt.Errorf("failed to add item: %w", err)
	}

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
		ListId:    item.ListId,
		ItemCode:  newItem.ItemCode,
		CreatedAt: item.CreatedAt,
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
