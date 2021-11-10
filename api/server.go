package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AlexandreLira/goApi/api/router"
	"github.com/AlexandreLira/goApi/config"
)

func Run() {
	config.Load()
	fmt.Printf("\n\tListening... localhost:%d\n", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
