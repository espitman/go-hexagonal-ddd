package repository

// implement team domain repository

import (
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/models"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/repositories"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/infrastructure/api"
	"strconv"
)

type teamRepository struct {
	apiClient *api.Client
}

func NewTeamRepository(teamApiClient *api.Client) repositories.TeamRepository {
	return &teamRepository{
		apiClient: teamApiClient,
	}
}

func (r *teamRepository) GetByID(id int64) (*models.Team, error) {
	idString := strconv.FormatInt(id, 10)
	var team models.Team
	_ = r.apiClient.Get("team.php?id="+idString, &team)
	return &team, nil
}
