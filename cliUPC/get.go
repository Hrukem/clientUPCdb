// Package cliUPC implements requests to the site
package cliUPC

import (
	"fmt"
	"golang_ninja/clientUPCdb/config"
	"golang_ninja/clientUPCdb/console"
	"io/ioutil"
	"net/http"
)

// Get function implements GET requests to the site with apikey
func Get() []byte {
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

	cli := http.Client{}
	resp, err := cli.Do(req)
	defer func() {
		if err = resp.Body.Close(); err != nil {
			fmt.Println("error close Body", err)
		}
	}()

	answer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Body)

	return answer
}
