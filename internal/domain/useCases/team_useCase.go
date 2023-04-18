package useCases

import "git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/models"

type TeamUseCase interface {
	GetTeamByID(teamId int64) (*models.Team, error)
}
