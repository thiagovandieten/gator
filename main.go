package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/thiagovandieten/gator/internal/config"
)

func fatal(e error) {
	fmt.Println(e.Error())
	os.Exit(1)
}

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		fatal(err)
	}

	s := state{
		cfg: &cfg,
	}

	cmds := Commands{
		CommandHandlers: make(map[string]func(*state, Command) error),
	}

	err = cmds.register("login", handlerLogin)
	if err != nil {
		fatal(err)
	}

	if len(os.Args) < 2 {
		err := errors.New("no command given")
		fatal(err)
	}

	args := os.Args[1:]
	if len(args) < 2 {
		args = append(args, "")
	}

	c := Command{
		Name: args[0],
		Args: args[1:],
	}

	err = cmds.run(&s, c)
	if err != nil {
		fatal(err)
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
