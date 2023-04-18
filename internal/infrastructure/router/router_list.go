package router

import (
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/app/handlers"
	appServices "git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/app/services"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/services"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/infrastructure/api"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/infrastructure/repository"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/infrastructure/router/middlewares"
)

func (router *Router) listRouter() {
	listRepository := repository.NewListRepository(router.mongoClient, "jbm-wishes")
	listUseCase := services.NewListService(listRepository)

	itemRepository := repository.NewItemRepository(router.mongoClient, "jbm-wishes")
	itemUseCase := services.NewItemService(itemRepository)

	teamApiClient := api.NewAPIClient("http://varzesh3.boum.ir/")
	teamRepository := repository.NewTeamRepository(teamApiClient)
	teamUseCase := services.NewTeamService(teamRepository)

	appListService := appServices.NewListService(listUseCase, itemUseCase, teamUseCase)
	listHandler := handlers.NewListHandler(*appListService)

	listGroup := router.r.Group("/list")
	listGroup.Use(middlewares.Authorize)

	listGroup.GET("", listHandler.GetAll)
	listGroup.GET("/:id", listHandler.GetById)
	listGroup.POST("", listHandler.Create)
	listGroup.PUT("/:id", listHandler.Update)
	listGroup.DELETE("/:id", listHandler.Delete)
}
