package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Dagime-Teshome/blog_aggregator/internal/cli"
	"github.com/Dagime-Teshome/blog_aggregator/internal/config"
	"github.com/Dagime-Teshome/blog_aggregator/internal/database"
	"github.com/Dagime-Teshome/blog_aggregator/internal/middleware"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	state := &cli.State{
		Db:     database.New(db),
		Config: &cfg,
	}

	commands := cli.NewCommands()

	commands.Register("help", cli.HandlerHelp(commands),
		"Show available commands", "gator help")

	commands.Register("login", cli.HandlerLogin,
		"Log in an existing user", "gator login <username>")
	commands.Register("register", cli.Register,
		"Register a new user", "gator register <username>")
	commands.Register("users", cli.GetUsers,
		"List all users", "gator users")

	commands.Register("reset", cli.Reset,
		"Reset all database tables", "gator reset")

	commands.Register("agg", cli.Agg,
		"Start feed aggregation", "gator agg <interval>")
	commands.Register("addfeed", middleware.LoggedInMiddleWare(cli.AddFeed),
		"Add a new feed", "gator addfeed <name> <url>")
	commands.Register("feeds", cli.FeedsList,
		"List all feeds", "gator feeds")
	commands.Register("follow", middleware.LoggedInMiddleWare(cli.Follow),
		"Follow a feed", "gator follow <url>")
	commands.Register("following", middleware.LoggedInMiddleWare(cli.Following),
		"List feeds you follow", "gator following")
	commands.Register("unfollow", middleware.LoggedInMiddleWare(cli.UnfollowFeed),
		"Unfollow a feed", "gator unfollow <url>")
	commands.Register("browse", middleware.LoggedInMiddleWare(cli.Browse),
		"Browse posts from followed feeds", "gator browse [limit]")

	if len(os.Args) < 2 {
		fmt.Print(commands.Help())
		os.Exit(1)
	}

	command := cli.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	if err := commands.Run(state, command); err != nil {
		log.Fatal(err)
	}
}
