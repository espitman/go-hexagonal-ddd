package main

import (
	_ "git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/docs"
	"git.alibaba.ir/saeedheidari-go-prototypes/jbm-wishes/internal/infrastructure/router"
)

// @title           Jabama Wishes
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
	r := router.NewRouter()
	r.InitRouter()
}
