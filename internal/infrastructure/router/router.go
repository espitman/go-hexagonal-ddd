package router

import (
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/domain/repositories"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/infrastructure/api"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/infrastructure/database/mongodb"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/infrastructure/redis"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/infrastructure/repository"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
)

type Router struct {
	r               *gin.Engine
	mongoClient     *mongo.Client
	redisConnection *redis.Connection
	listRepository  *repositories.ListRepository
	itemRepository  *repositories.ItemRepository
	teamRepository  *repositories.TeamRepository
}

func NewRouter() *Router {
	r := gin.Default()
	mongoClient, _ := mongodb.NewMongoClient()
	redisConnection := redis.NewConnection("localhost:6379", "1234", 0)
	redisClient, _ := redisConnection.NewClient()

	listRepository := repository.NewListRepository(mongoClient, "jbm-wishes")

	itemRepository := repository.NewItemRepository(mongoClient, "jbm-wishes")

	teamApiClient := api.NewAPIClient("http://varzesh3.boum.ir/")
	teamRepository := repository.NewTeamRepository(teamApiClient, redisClient)

	return &Router{
		r:               r,
		mongoClient:     mongoClient,
		redisConnection: redisConnection,
		listRepository:  &listRepository,
		itemRepository:  &itemRepository,
		teamRepository:  &teamRepository,
	}
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
	router.teamRouter()
	router.r.Run() // listen and serve on 0.0.0.0:8080
}
