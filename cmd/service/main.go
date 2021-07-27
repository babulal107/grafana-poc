package main

import (
	"fmt"
	"github.com/grafana-poc/internal/router"
	"log"
)

func main() {

	ginRouter := router.Init()
	if err := ginRouter.Run(fmt.Sprintf(":%d", 8080)); err != nil {
		log.Println("Unable to start application", err)
	}
}
