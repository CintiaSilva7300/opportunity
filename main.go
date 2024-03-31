package main

import (
	"github.com/CintiaSilva7300/go-opportunity/config"
	"github.com/CintiaSilva7300/go-opportunity/router"
)

var (
	logger *config.Logger
)

func main() {
	//inicializa o logger
	logger = config.GetLogger("main")

	//inicializa a config
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err) //logger de erro
		return
	}

	//inicializa a aplicação
	router.Initialize()
}
