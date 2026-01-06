package main

import (
	"database/sql"
	"fmt"
	"go/gator/internal/cli"
	"go/gator/internal/config"
	"go/gator/internal/database"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments")
		return
	}
	state := cli.State{}
	configStruct, err := config.Read()
	if err != nil {
		log.Fatal(fmt.Errorf("error reading config file:%w", err))
		return
	}
	db, err := sql.Open("postgres", configStruct.Db_url)

	db_queries := database.New(db)
	state.Db = db_queries

	state.Config = &configStruct
	commands := cli.Commands{
		ComMap: make(map[string]func(*cli.State, cli.Command) error, 0),
	}
	commands.Register("login", cli.HandlerLogin)
	commands.Register("register", cli.Register)
	commands.Register("reset", cli.Reset)
	var command cli.Command
	command.Name = os.Args[1]
	command.Args = append(command.Args, os.Args[2:]...)
	if err := commands.Run(&state, command); err != nil {
		log.Fatal(err)
	}

}
