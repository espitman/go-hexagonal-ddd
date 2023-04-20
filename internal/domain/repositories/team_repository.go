package repositories

// implemented by mongodb team_repository

import "github.com/espitman/go-hexagonal-ddd/internal/domain/models"

type TeamRepository interface {
	GetByID(id int64) (*models.Team, error)
}
