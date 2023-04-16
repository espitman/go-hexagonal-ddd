package repositories

// implemented by mongodb team_repository

import "git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/models"

type TeamRepository interface {
	GetByID(id int64) (*models.Team, error)
}
