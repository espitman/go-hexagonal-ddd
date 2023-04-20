package repository

// implement team domain repository

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/espitman/go-hexagonal-ddd/internal/domain/models"
	"github.com/espitman/go-hexagonal-ddd/internal/domain/repositories"
	"github.com/espitman/go-hexagonal-ddd/internal/infrastructure/api"
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

	if team.Id == 0 {
		return nil, errors.New("team not found!")
	}

	jsonTeam, _ := json.Marshal(team)
	r.redisClient.Set(ctx, key, jsonTeam, 0)
	return &team, nil
}
