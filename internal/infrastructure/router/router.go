package router

import (
	"github.com/espitman/go-hexagonal-ddd/internal/config"
	"github.com/espitman/go-hexagonal-ddd/internal/domain/repositories"
	"github.com/espitman/go-hexagonal-ddd/internal/infrastructure/api"
	"github.com/espitman/go-hexagonal-ddd/internal/infrastructure/database/mongodb"
	"github.com/espitman/go-hexagonal-ddd/internal/infrastructure/redis"
	"github.com/espitman/go-hexagonal-ddd/internal/infrastructure/repository"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
)

type Router struct {
	r               *gin.Engine
	conf            *config.Config
	mongoClient     *mongo.Client
	redisConnection *redis.Connection
	listRepository  *repositories.ListRepository
	itemRepository  *repositories.ItemRepository
	teamRepository  *repositories.TeamRepository
}

func NewRouter(conf *config.Config) *Router {
	r := gin.Default()
	mongoClient, _ := mongodb.NewMongoClient(conf)
	redisConnection := redis.NewConnection(conf.RedisHost+":"+conf.RedisPort, conf.RedisPassword, conf.RedisDb)
	redisClient, _ := redisConnection.NewClient()

	listRepository := repository.NewListRepository(mongoClient, conf.DBDatabase)

	itemRepository := repository.NewItemRepository(mongoClient, conf.DBDatabase)

	teamApiClient := api.NewAPIClient(conf.APIBaseUrl)
	teamRepository := repository.NewTeamRepository(teamApiClient, redisClient)

	return &Router{
		r:               r,
		conf:            conf,
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
