package app

import (
	"sellease-ai/app/router"
	"sellease-ai/config"
	"sellease-ai/internal/handler"
	"sellease-ai/internal/repository"
	"sellease-ai/internal/usecase"
	"sellease-ai/logger"
)

func Start() {
	config := config.GetConfig()

	// db, err := database.PrepareDatabase()
	// if err != nil {
	// 	panic(err)
	// }

	_, err := logger.InitLogger(config.Environment)
	if err != nil {
		logger.Errorf("error initializing logger", err)
	}

	repo := repository.InitRepository()
	uc := usecase.Init(repo)
	hndlr := handler.Init(uc)

	router := router.PrepareRouter(&router.RouterContext{
		// DB:      db,
		Repo:    repo,
		Usecase: uc,
		Handler: hndlr,
	})

	logger.Infof("server running at port %s", config.ServerPort)
	err = router.Run(":" + config.ServerPort)
	if err != nil {
		logger.Fatalf("error running server - %s", err.Error())
	}
}
