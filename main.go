package main

import (
	"log"
	"os"

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
		CommandHandlers: make(map[string]func(*state, Command) error),
	}

	cmds.register("login", handlerLogin)
	c := Command{
		Name: os.Args[0],
		Args: os.Args[1:],
	}

	cmds.run(&s, c)
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
