package router

import (
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/app/handlers"
	appServices "git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/app/services"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/services"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/infrastructure/database/mongodb"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/pkg/middleware"
)

func (router *Router) itemRouter() {
	itemRepository := mongodb.NewItemRepository(router.mongoClient, "jbm-wishes")
	itemUseCase := services.NewItemService(itemRepository)

	listRepository := mongodb.NewListRepository(router.mongoClient, "jbm-wishes")
	listUseCase := services.NewListService(listRepository)

	appItemService := appServices.NewItemService(itemUseCase, listUseCase)
	itemHandler := handlers.NewItemHandler(*appItemService)

	itemGroup := router.r.Group("/item")
	itemGroup.Use(middleware.AuthMiddleware)

	itemGroup.POST("", itemHandler.Create)
	itemGroup.DELETE("/:id", itemHandler.Delete)
}
