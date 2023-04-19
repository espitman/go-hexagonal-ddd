package router

import (
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/app/handlers"
	appServices "git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/app/services"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/services"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/infrastructure/router/middlewares"
)

func (router *Router) itemRouter() {
	listUseCase := services.NewListService(*router.listRepository)
	itemUseCase := services.NewItemService(*router.itemRepository)
	teamUseCase := services.NewTeamService(*router.teamRepository)

	appItemService := appServices.NewItemService(itemUseCase, listUseCase, teamUseCase)
	itemHandler := handlers.NewItemHandler(*appItemService)

	itemGroup := router.r.Group("/item")
	itemGroup.Use(middlewares.Authorize)

	itemGroup.POST("", itemHandler.Create)
	itemGroup.DELETE("/:id", itemHandler.Delete)
}
