package services

import (
	"github.com/espitman/go-hexagonal-ddd/internal/domain/models"
	"github.com/espitman/go-hexagonal-ddd/internal/domain/repositories"
	"sync"
)

type TeamService struct {
	teamRepository repositories.TeamRepository
}

func NewTeamService(teamRepository repositories.TeamRepository) *TeamService {
	return &TeamService{teamRepository: teamRepository}
}

func (s *TeamService) GetTeamByID(id int64) (*models.Team, error) {
	return s.teamRepository.GetByID(id)
}

func (s *TeamService) GetTeamsByIds(ids []int64) ([]*models.Team, error) {
	lenIds := len(ids)
	var wg sync.WaitGroup
	wg.Add(lenIds)
	teams := make([]*models.Team, 0, lenIds)
	for _, id := range ids {
		go func(id int64) {
			teamInfo, _ := s.GetTeamByID(id)
			teams = append(teams, teamInfo)
			wg.Done()
		}(id)
	}
	wg.Wait()
	return teams, nil
}
