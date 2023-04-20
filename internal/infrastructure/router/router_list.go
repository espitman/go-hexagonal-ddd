package router

import (
	"github.com/espitman/go-hexagonal-ddd/internal/app/handlers"
	appServices "github.com/espitman/go-hexagonal-ddd/internal/app/services"
	"github.com/espitman/go-hexagonal-ddd/internal/domain/services"
	"github.com/espitman/go-hexagonal-ddd/internal/infrastructure/router/middlewares"
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
