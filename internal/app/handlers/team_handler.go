package handlers

import (
	_ "github.com/espitman/go-hexagonal-ddd/internal/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TeamHandler struct {
	service appServices.TeamService
}

func NewTeamHandler(service appServices.TeamService) *TeamHandler {
	return &TeamHandler{service}
}

// GetById Get team by ID.
// @Summary Get team by ID
// @Description Get a team using the provided ID
// @Tags Team
// @Param id path string true "Team ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} appModel.Team
// @Failure 401 {object} commonModels.ErrorResponse
// @Failure 404 {object} commonModels.ErrorResponse
// @Router /team/{id} [get]
func (app *TeamHandler) GetById(c *gin.Context) {
	teamIdParam := c.Param("id")
	teamId, _ := strconv.ParseInt(teamIdParam, 10, 64)
	team, err := app.service.GetTeamByID(teamId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "team not found"})
		return
	}
	c.JSON(http.StatusOK, team)
}
