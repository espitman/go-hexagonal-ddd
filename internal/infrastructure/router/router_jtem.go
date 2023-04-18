package router

import (
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/app/handlers"
	appServices "git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/app/services"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/services"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/infrastructure/repository"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/infrastructure/router/middlewares"
)

func (router *Router) itemRouter() {
	itemRepository := repository.NewItemRepository(router.mongoClient, "jbm-wishes")
	itemUseCase := services.NewItemService(itemRepository)

	listRepository := repository.NewListRepository(router.mongoClient, "jbm-wishes")
	listUseCase := services.NewListService(listRepository)

	appItemService := appServices.NewItemService(itemUseCase, listUseCase)
	itemHandler := handlers.NewItemHandler(*appItemService)

	itemGroup := router.r.Group("/item")
	itemGroup.Use(middlewares.Authorize)

	itemGroup.POST("", itemHandler.Create)
	itemGroup.DELETE("/:id", itemHandler.Delete)
}
