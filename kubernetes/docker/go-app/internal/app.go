package internal

import "github.com/anynines/tmp-homework-sh/kubernetes/docker/go-app/internal/api"

func Run() {
	router := api.SetupAPI()
	router.Run(":8080")
}
