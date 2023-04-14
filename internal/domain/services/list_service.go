package services

import (
	"errors"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/models"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/repositories"
)

type ListService struct {
	listRepository repositories.ListRepository
}

func NewListService(listRepository repositories.ListRepository) *ListService {
	return &ListService{listRepository: listRepository}
}

func (s *ListService) AddList(list *models.List) (*models.List, error) {
	_, err := s.listRepository.GetByName(list.Name)
	if err == nil {
		return nil, errors.New("list name already exists")
	}
	return s.listRepository.Create(list)
}

func (s *ListService) GetListByID(id string) (*models.List, error) {
	return s.listRepository.GetByID(id)
}

func (s *ListService) ListLists() ([]*models.List, error) {
	return s.listRepository.GetAll()
}

func (s *ListService) UpdateList(id string, list *models.List) (*models.List, error) {
	_, err := s.listRepository.GetByID(id)
	if err != nil {
		return nil, errors.New("list doesnt exist!")
	}
	return s.listRepository.Update(id, list)
}

func (s *ListService) DeleteList(id string) error {
	_, err := s.listRepository.GetByID(id)
	if err != nil {
		return errors.New("list doesnt exist!")
	}
	return s.listRepository.Delete(id)
}

func (s *ListService) GetListByName(name string) (*models.List, error) {
	return s.listRepository.GetByName(name)
}
