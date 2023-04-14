package repositories

import "git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/models"

type ListRepository interface {
	Create(list *models.List) (*models.List, error)
	GetByID(id string) (*models.List, error)
	GetByName(name string) (*models.List, error)
	Update(id string, list *models.List) (*models.List, error)
	Delete(id string) error
	GetAll() ([]*models.List, error)
}
