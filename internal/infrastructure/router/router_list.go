package router

import (
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/app/handlers"
	appServices "git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/app/services"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/services"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/infrastructure/database/mongodb"
)

func (router *Router) listRouter() {
	listRepository := mongodb.NewListRepository(router.mongoClient, "jbm-wishes")
	listUseCase := services.NewListService(listRepository)
	appListService := appServices.NewListService(listUseCase)
	listHandler := handlers.NewListHandler(*appListService)

	router.r.GET("/list", listHandler.GetAll)
	router.r.GET("/list/:id", listHandler.GetById)
	router.r.POST("/list", listHandler.Create)
	router.r.PUT("/list/:id", listHandler.Update)
	router.r.DELETE("/list/:id", listHandler.Delete)
}
