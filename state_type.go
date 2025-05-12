package main

import (
	"github.com/time4soup/blog_aggregator_go/internal/config"
	"github.com/time4soup/blog_aggregator_go/internal/database"
)

func NewState(dbQueries *database.Queries, cfg *config.Config) state {
	return state{dbQueries, cfg}
}

type state struct {
	db     *database.Queries
	config *config.Config
}
