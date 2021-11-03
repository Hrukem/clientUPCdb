// Package cliUPC implements requests to the site
package cliUPC

import (
	"bytes"
	"fmt"
	"golang_ninja/clientUPCdb/config"
	"golang_ninja/clientUPCdb/console"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

// Saving function implements POST requests to the site with apikey
func Saving() []byte {
	fmt.Println("enter data")
	fmt.Println("enter UPC code")
	upcCode := console.InputDateFromConsole()

	dataUPC := enterData()

	mp, body := createMultipartData(dataUPC)

	req := createRequest(mp, body, upcCode)

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

	return answer
}

// enterData function implements input from console
func enterData() map[string]string {
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

func createMultipartData(dataUPC map[string]string) (*multipart.Writer, io.ReadWriter) {
	var body io.ReadWriter = bytes.NewBufferString("")
	mp := multipart.NewWriter(body)
	for k, v := range dataUPC {
		err := mp.WriteField(k, v)
		if err != nil {
			fmt.Println("Error 'range dataUPC'", err)
		}
	}
	return mp, body
}

func createRequest(mp *multipart.Writer, body io.ReadWriter, upcCode string) *http.Request {
	contentType := fmt.Sprintf("multipart/form-data; boundary=%v", mp.Boundary())
	if err := mp.Close(); err != nil {
		fmt.Println("Error close 'mp'", err)
	}

	url := "https://api.upcdatabase.org/product/" +
		upcCode +
		"?apikey=" +
		config.Cfg.ApiKey

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", contentType)

	return req
}
