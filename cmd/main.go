package main

import (
	_ "github.com/espitman/go-hexagonal-ddd/docs"
	"github.com/espitman/go-hexagonal-ddd/internal/config"
	"github.com/espitman/go-hexagonal-ddd/internal/infrastructure/router"
)

// @title           Go Hexagonal ddd
// @version         1.0
// @description     This is a sample hexagonal domain driven for golang.
// @contact.name   API Support
// @contact.email  s.heidar@jabama.com
// @host      localhost:8080
// @BasePath  /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	conf, _ := config.LoadConfig("./config/config.json")
	r := router.NewRouter(conf)
	r.InitRouter()
}
