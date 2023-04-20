package appServices

import (
	"fmt"
	appModel "github.com/espitman/go-hexagonal-ddd/internal/app/models"
	"github.com/espitman/go-hexagonal-ddd/internal/domain/useCases"
)

type TeamService struct {
	teamUseCases useCases.TeamUseCase
	itemUseCase  useCases.ItemUseCase
}

func NewTeamService(teamUseCases useCases.TeamUseCase) *TeamService {
	return &TeamService{
		teamUseCases: teamUseCases,
	}
}

func (s *TeamService) GetTeamByID(id int64) (*appModel.Team, error) {
	team, err := s.teamUseCases.GetTeamByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get team by ID %s: %w", id, err)
	}
	if team == nil {
		return nil, nil
	}
	appTeam := &appModel.Team{
		TeamId: team.Id,
		Name:   team.Name,
		Crest:  team.Crest,
	}
	return appTeam, nil
}
