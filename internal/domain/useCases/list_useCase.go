package useCases

import "git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/models"

type ListUseCase interface {
	AddList(list *models.List) (*models.List, error)
	GetListByID(id string) (*models.List, error)
	ListLists() ([]*models.List, error)
	UpdateList(id string, list *models.List) (*models.List, error)
	DeleteList(id string) error
	GetListByName(name string) (*models.List, error)
}