package repository

// implement team domain repository

import (
	"context"
	"encoding/json"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/models"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/repositories"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/infrastructure/api"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type teamRepository struct {
	apiClient   *api.Client
	redisClient *redis.Client
}

func NewTeamRepository(teamApiClient *api.Client, redisClient *redis.Client) repositories.TeamRepository {
	return &teamRepository{
		apiClient:   teamApiClient,
		redisClient: redisClient,
	}
}

func (r *teamRepository) GetByID(id int64) (*models.Team, error) {
	ctx := context.Background()
	idString := strconv.FormatInt(id, 10)
	key := "team:" + strconv.FormatInt(id, 10)
	var team models.Team

	cachedTeam, err := r.redisClient.Get(ctx, key).Result()
	if err == nil && cachedTeam != "" {
		err = json.Unmarshal([]byte(cachedTeam), &team)
		if err != nil {
			return nil, err
		}
		return &team, nil
	}

	_ = r.apiClient.Get("team.php?id="+idString, &team)
	jsonTeam, _ := json.Marshal(team)
	r.redisClient.Set(ctx, key, jsonTeam, 0)
	return &team, nil
}
