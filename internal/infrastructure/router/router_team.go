package router

import (
	"github.com/espitman/go-hexagonal-ddd/internal/app/handlers"
	appServices "github.com/espitman/go-hexagonal-ddd/internal/app/services"
	"github.com/espitman/go-hexagonal-ddd/internal/domain/services"
)

func (router *Router) teamRouter() {
	teamUseCase := services.NewTeamService(*router.teamRepository)

	appTeamService := appServices.NewTeamService(teamUseCase)
	teamHandler := handlers.NewTeamHandler(*appTeamService)

	teamGroup := router.r.Group("/team")

	teamGroup.GET("/:id", teamHandler.GetById)
}
