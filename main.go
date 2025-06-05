package main

import (
	"log"

	"github.com/thiagovandieten/gator/internal/config"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

	s := state{
		cfg: &cfg,
	}

	cmds := Commands{
		Names: make(map[string]func(*state, Command)),
	}
	// err = config.SetUser("Thiago", cfg)
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }

	// cfg, err = config.GetConfig()

	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// fmt.Printf("Configuration: %+v\n", cfg)
}
