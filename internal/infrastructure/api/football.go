package api

import (
	"encoding/json"
	"fmt"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/models"
	"net/http"
	"strconv"
)

type FootballAPIClient struct {
	baseURL string
}

func NewFootballAPIClient() *FootballAPIClient {
	return &FootballAPIClient{baseURL: "https://api.football-data.org/v4/"}
}

func (c *FootballAPIClient) GetByID(id int64) (*models.Team, error) {
	idString := strconv.FormatInt(id, 10)
	req, err := http.NewRequest("GET", c.baseURL+"teams/"+idString, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Auth-Token", "b5862c05790e48f3881e32cf3faa3edc")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get team by name, status code: %d", resp.StatusCode)
	}
	var team *models.Team
	if err := json.NewDecoder(resp.Body).Decode(&team); err != nil {
		return nil, err
	}
	return team, nil
}
