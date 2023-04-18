package router

import (
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/app/handlers"
	appServices "git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/app/services"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/services"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/infrastructure/api"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/infrastructure/repository"
)

func (router *Router) teamRouter() {
	teamApiClient := api.NewAPIClient("http://varzesh3.boum.ir/")
	teamRepository := repository.NewTeamRepository(teamApiClient)
	teamUseCase := services.NewTeamService(teamRepository)

	appTeamService := appServices.NewTeamService(teamUseCase)
	teamHandler := handlers.NewTeamHandler(*appTeamService)

	teamGroup := router.r.Group("/team")

	teamGroup.GET("/:id", teamHandler.GetById)
}
