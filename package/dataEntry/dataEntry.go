package dataEntry

import (
	"fmt"
	"golang_ninja/clientUPCdb/package/console"
)

// EnterData function implements input from console
func EnterData() map[string]string {
	dataUPC := make(map[string]string)

	fmt.Println("enter title")
	dataUPC["title"] = console.InputDateFromConsole()

	fmt.Println("enter alias")
	dataUPC["alias"] = console.InputDateFromConsole()

	fmt.Println("enter description")
	dataUPC["description"] = console.InputDateFromConsole()

	fmt.Println("enter manufacturer")
	dataUPC["manufacturer"] = console.InputDateFromConsole()

	fmt.Println("enter msrp")
	dataUPC["msrp"] = console.InputDateFromConsole()

	fmt.Println("enter category")
	dataUPC["category"] = console.InputDateFromConsole()

	return dataUPC
}
