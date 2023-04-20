package router

import (
	"github.com/espitman/go-hexagonal-ddd/internal/app/handlers"
	appServices "github.com/espitman/go-hexagonal-ddd/internal/app/services"
	"github.com/espitman/go-hexagonal-ddd/internal/domain/services"
	"github.com/espitman/go-hexagonal-ddd/internal/infrastructure/router/middlewares"
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
