package app

import (
	"fmt"
	"net/http"
	"product-packaging/internal/config"
	"product-packaging/internal/controllers"
	"product-packaging/internal/repository"
	"product-packaging/internal/router"
	"product-packaging/internal/usecases"
)

var httpRouter router.Router = router.NewMuxRouter()

func Run() {
	cfg := config.NewConfig()

	dbHandler := repository.NewPostgresHandler(cfg.Repo)
	boxRepo := repository.NewBoxRepo(dbHandler)
	box := usecases.NewBox(boxRepo)
	boxController := controllers.NewBoxController(box)

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "App is up and running...")
	})
	httpRouter.POST("/boxes/add", boxController.Add)
	httpRouter.GET("/boxes/sgtins", boxController.GetBoxesBySgtin)
	httpRouter.GET("/boxes/gtins", boxController.GetBoxesAndPacksByGtin)
	httpRouter.SERVE(cfg.Port)
}
