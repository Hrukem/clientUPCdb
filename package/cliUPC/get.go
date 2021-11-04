// Package cliUPC implements requests to the site
package cliUPC

import (
	"fmt"
	"golang_ninja/clientUPCdb/package/config"
	"golang_ninja/clientUPCdb/package/console"
	"golang_ninja/clientUPCdb/package/sendRequest"
	"net/http"
)

// Get function implements GET requests to the site with apikey
func Get() string {
	fmt.Println("receiving data")
	fmt.Println("enter UPC code (12 or 13 numbers)")
	upcCode := console.InputDateFromConsole()

	url := "https://api.upcdatabase.org/product/" +
		upcCode +
		"?apikey=" +
		config.Cfg.ApiKey

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	return sendRequest.SendRequest(req)
}
