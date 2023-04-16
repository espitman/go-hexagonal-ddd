package router

import (
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/infrastructure/database/mongodb"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
)

type Router struct {
	r           *gin.Engine
	mongoClient *mongo.Client
}

func NewRouter() *Router {
	r := gin.Default()
	mongoClient, _ := mongodb.NewMongoClient()
	return &Router{r: r, mongoClient: mongoClient}
}

func (router *Router) InitRouter() {
	router.r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.listRouter()
	router.itemRouter()
	router.r.Run() // listen and serve on 0.0.0.0:8080
}
