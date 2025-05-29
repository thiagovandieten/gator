package main

import (
	"fmt"
	"log"

	"github.com/thiagovandieten/gator/internal/config"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = config.SetUser("Thiago", cfg)
	if err != nil {
		log.Fatalf(err.Error())
	}

	cfg, err = config.GetConfig()

	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Configuration: %+v\n", cfg)
}
