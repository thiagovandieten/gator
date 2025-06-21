package main

import (
	"text/tabwriter"

	"github.com/thiagovandieten/gator/internal/config"
	"github.com/thiagovandieten/gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
	tab *tabwriter.Writer
}
