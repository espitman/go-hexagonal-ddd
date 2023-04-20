package appServices

import (
	appModel "github.com/espitman/go-hexagonal-ddd/internal/app/models"
	"github.com/espitman/go-hexagonal-ddd/internal/domain/models"
	"strconv"
)

func (s *ListService) getTeamsByItems(items []*models.Item) []appModel.Team {
	teamsItemIds := make(map[string]string, len(items))
	teamIds := make([]int64, 0, len(items))

	for _, mItem := range items {
		teamIds = append(teamIds, mItem.ItemCode)
		teamsItemIds[strconv.FormatInt(mItem.ItemCode, 10)] = mItem.ID
	}
	teams, _ := s.teamUseCase.GetTeamsByIds(teamIds)
	modelTeams := make([]appModel.Team, 0, len(teams))
	for _, mTeam := range teams {
		modelTeams = append(modelTeams, appModel.Team{
			Id:     teamsItemIds[strconv.FormatInt(mTeam.Id, 10)],
			TeamId: mTeam.Id,
			Name:   mTeam.Name,
			Crest:  mTeam.Crest,
		})
	}
	return modelTeams
}
