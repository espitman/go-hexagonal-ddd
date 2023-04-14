package router

import (
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/app/handlers"
	appServices "git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/app/services"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/services"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/infrastructure/database/mongodb"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/pkg/middleware"
)

func (router *Router) listRouter() {
	listRepository := mongodb.NewListRepository(router.mongoClient, "jbm-wishes")
	listUseCase := services.NewListService(listRepository)
	appListService := appServices.NewListService(listUseCase)
	listHandler := handlers.NewListHandler(*appListService)

	listGroup := router.r.Group("/list")
	listGroup.Use(middleware.AuthMiddleware)

	listGroup.GET("", listHandler.GetAll)
	listGroup.GET("/:id", listHandler.GetById)
	listGroup.POST("", listHandler.Create)
	listGroup.PUT("/:id", listHandler.Update)
	listGroup.DELETE("/:id", listHandler.Delete)
}
