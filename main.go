package main

import (
	"fmt"
	"go/gator/internal/config"
	"log"
)

func main() {

	configStruct, err := config.Read()
	if err != nil {
		log.Fatal(fmt.Errorf("error reading config file:%w", err))
	}
	configStruct.SetUser("Dagime")
	newConfig, err := config.Read()
	if err != nil {
		log.Fatal(fmt.Errorf("Error reading config file%w", err))
	}
	fmt.Println(newConfig)
}
