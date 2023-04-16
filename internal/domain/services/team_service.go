package services

import (
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/models"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/repositories"
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
