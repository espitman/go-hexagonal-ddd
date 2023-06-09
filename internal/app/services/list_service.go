package appServices

import (
	"errors"
	"fmt"
	appModel "github.com/espitman/go-hexagonal-ddd/internal/app/models"
	"github.com/espitman/go-hexagonal-ddd/internal/domain/models"
	"github.com/espitman/go-hexagonal-ddd/internal/domain/useCases"
	"sync"
)

type ListService struct {
	listUseCases useCases.ListUseCase
	itemUseCase  useCases.ItemUseCase
	teamUseCase  useCases.TeamUseCase
}

func NewListService(
	listUseCases useCases.ListUseCase,
	itemUseCase useCases.ItemUseCase,
	teamUseCase useCases.TeamUseCase,
) *ListService {
	return &ListService{
		listUseCases: listUseCases,
		itemUseCase:  itemUseCase,
		teamUseCase:  teamUseCase,
	}
}

func (s *ListService) GetLists(userId int64) ([]*appModel.List, error) {
	lists, err := s.listUseCases.ListLists(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get lists: %w", err)
	}

	if len(lists) == 0 {
		return nil, errors.New("no lists found")
	}

	appLists := make([]*appModel.List, 0, len(lists))
	for _, list := range lists {
		appList := &appModel.List{
			ID:     list.ID,
			Name:   list.Name,
			UserId: list.UserId,
		}
		appLists = append(appLists, appList)
	}

	return appLists, nil
}

func (s *ListService) CreateList(newList *appModel.NewList, userId int64) (*appModel.List, error) {
	list := &models.List{
		Name:   newList.Name,
		UserId: userId,
	}

	list, err := s.listUseCases.AddList(list)
	if err != nil {
		return nil, fmt.Errorf("failed to create list: %w", err)
	}
	appList := &appModel.List{
		ID:        list.ID,
		Name:      list.Name,
		UserId:    list.UserId,
		CreatedAt: list.CreatedAt,
	}
	return appList, nil
}

func (s *ListService) GetListByID(id string) (*appModel.ListWithTeams, error) {
	var wg sync.WaitGroup
	wg.Add(2)

	var list *models.List
	var items []*models.Item
	var err error

	go func() {
		list, err = s.listUseCases.GetListByID(id)
		wg.Done()
	}()
	go func() {
		items, err = s.itemUseCase.GetItemsByListID(id)
		wg.Done()
	}()

	wg.Wait()

	if err != nil {
		return nil, fmt.Errorf("failed to get list by ID %s: %w", id, err)
	}
	if list == nil {
		return nil, nil
	}
	teams := s.getTeamsByItems(items)
	appList := &appModel.ListWithTeams{
		List: appModel.List{
			ID:        list.ID,
			Name:      list.Name,
			UserId:    list.UserId,
			CreatedAt: list.CreatedAt,
		},
		Teams: teams,
	}
	return appList, nil
}

func (s *ListService) UpdateList(id string, updatedList *appModel.NewList) (*appModel.List, error) {
	newList := models.List{
		Name: updatedList.Name,
	}
	list, err := s.listUseCases.UpdateList(id, &newList)
	if err != nil {
		return nil, fmt.Errorf("failed to update list with ID %s: %w", id, err)
	}

	appList := &appModel.List{
		ID:     list.ID,
		Name:   list.Name,
		UserId: list.UserId,
	}

	return appList, nil
}

func (s *ListService) DeleteList(id string) error {
	err := s.listUseCases.DeleteList(id)
	if err != nil {
		return fmt.Errorf("failed to delete list with ID %s: %w", id, err)
	}
	return nil
}
