package useCases

import "github.com/espitman/go-hexagonal-ddd/internal/domain/models"

type TeamUseCase interface {
	GetTeamByID(teamId int64) (*models.Team, error)
	GetTeamsByIds(teamIds []int64) ([]*models.Team, error)
}
