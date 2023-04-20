package repositories

// implemented by mongodb list_repository

import "github.com/espitman/go-hexagonal-ddd/internal/domain/models"

type ListRepository interface {
	Create(list *models.List) (*models.List, error)
	GetByID(id string) (*models.List, error)
	GetByName(name string) (*models.List, error)
	Update(id string, list *models.List) (*models.List, error)
	Delete(id string) error
	GetAll() ([]*models.List, error)
	GetAllByUserId(userId int64) ([]*models.List, error)
}
