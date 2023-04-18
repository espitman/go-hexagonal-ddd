package router

import (
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/app/handlers"
	appServices "git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/app/services"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/services"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/infrastructure/api"
)

func (router *Router) teamRouter() {
	teamRepository := api.NewFootballAPIClient()
	teamUseCase := services.NewTeamService(teamRepository)

	appTeamService := appServices.NewTeamService(teamUseCase)
	teamHandler := handlers.NewTeamHandler(*appTeamService)

	teamGroup := router.r.Group("/team")

	teamGroup.GET("/:id", teamHandler.GetById)
}
