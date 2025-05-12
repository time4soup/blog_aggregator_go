package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/time4soup/blog_aggregator_go/internal/config"
	"github.com/time4soup/blog_aggregator_go/internal/database"

	_ "github.com/lib/pq"
)

//goose postgres "postgres://zakirk:@localhost:5432/gator" up

func main() {
	cfg := config.Read()
	cmds := NewCommands()
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", handlerBrowse)

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	s := NewState(dbQueries, &cfg)

	if len(os.Args) <= 1 {
		log.Fatal("not enough arguments provided.")
	}

	args := os.Args[1:]
	cmd := command{args[0], args[1:]}
	fmt.Println()
	err = cmds.run(&s, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
