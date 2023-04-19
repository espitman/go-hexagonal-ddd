package router

import (
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/app/handlers"
	appServices "git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/app/services"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/services"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/infrastructure/router/middlewares"
)

func (router *Router) listRouter() {
	listUseCase := services.NewListService(*router.listRepository)
	itemUseCase := services.NewItemService(*router.itemRepository)
	teamUseCase := services.NewTeamService(*router.teamRepository)

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
