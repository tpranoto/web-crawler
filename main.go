package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tpranoto/gochallenge/common/config"
	presenterPDetail "github.com/tpranoto/gochallenge/page_detail/presenter"
	repoPDetail "github.com/tpranoto/gochallenge/page_detail/repo"
	usecasePDetail "github.com/tpranoto/gochallenge/page_detail/usecase"
	"github.com/tpranoto/gochallenge/rest"
)

func main() {
	//init configurations
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("fail to get configurations")
	}

	//PageDetail deps
	pageRepo := repoPDetail.NewPageContentRepo()
	pageUsecase := usecasePDetail.NewPageDetailUsecase(pageRepo)
	presenterPDetail := presenterPDetail.NewRestHandler(pageUsecase)

	//Assign rests endpoints
	router := http.NewServeMux()
	rest.RouterPresenter{
		Router:  router,
		PDetail: presenterPDetail,
	}.AssignPaths()

	log.Printf("serving on port: %d", cfg.Port.Main)
	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port.Main), router)
	if err != nil {
		log.Fatalf("failed to serve to %d", cfg.Port.Main)
	}
}
