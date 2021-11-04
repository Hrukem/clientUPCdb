package main

import (
	"fmt"
	"golang_ninja/clientUPCdb/package/cliUPC"
	"golang_ninja/clientUPCdb/package/config"
	"golang_ninja/clientUPCdb/package/console"
	"log"
)

func main() {
	fmt.Println("start Client")
	fmt.Println("For exit enter 'Q'")

	err := config.Config()
	if err != nil {
		log.Fatal("Error config! End program.")
	}

	for {
		fmt.Println("\nSelect task\n1 - saving data\n2 - get data\nQ - exit")
		task := console.InputDateFromConsole()
		switch task {
		case "1":
			res := cliUPC.Saving()
			fmt.Println(res)
		case "2":
			res := cliUPC.Get()
			fmt.Println(res)
		case "q", "Q":
			fmt.Println("End program")
			return
		default:
		}
	}
}
