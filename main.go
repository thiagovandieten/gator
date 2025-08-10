package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	_ "github.com/lib/pq"
	"github.com/thiagovandieten/gator/internal/config"
	"github.com/thiagovandieten/gator/internal/database"
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

	db, err := sql.Open("postgres", cfg.DBinURL)
	if err != nil {
		fatal(err)
	}
	dbQueries := database.New(db)

	tab := tabwriter.NewWriter(os.Stdout, 2, 1, 4, ' ', tabwriter.TabIndent)
	s := state{
		cfg: &cfg,
		db:  dbQueries,
		tab: tab,
	}

	cmds := Commands{
		CommandHandlers: make(map[string]func(*state, Command) error),
	}

	err = cmds.registerAll(map[string]func(*state, Command) error{
		"login":     handlerLogin,
		"register":  handlerRegister,
		"reset":     handlerReset,
		"users":     handlerGetAllUsers,
		"agg":       handlerAgg,
		"addfeed":   middlewareLoggedin(handlerAddFeed),
		"feeds":     handlerFeeds,
		"follow":    middlewareLoggedin(handlerFollow),
		"following": handlerFollowing,
		"unfollow":  middlewareLoggedin(handlerUnfollowing),
		"browse":    middlewareLoggedin(handlerBrowse),
	})
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
}
