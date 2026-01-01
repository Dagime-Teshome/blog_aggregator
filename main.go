package main

import (
	"fmt"
	"go/gator/internal/cli"
	"go/gator/internal/config"
	"log"
	"os"
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

	state.Config = &configStruct
	commands := cli.Commands{
		ComMap: make(map[string]func(*cli.State, cli.Command) error, 0),
	}
	commands.Register("login", cli.HandlerLogin)
	var command cli.Command
	command.Name = os.Args[1]
	command.Args = append(command.Args, os.Args[2:]...)
	if err := commands.Run(&state, command); err != nil {
		log.Fatal(err)
	}
}
