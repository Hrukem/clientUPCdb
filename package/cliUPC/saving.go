// Package cliUPC implements requests to the site
package cliUPC

import (
	"bytes"
	"fmt"
	"golang_ninja/clientUPCdb/package/config"
	"golang_ninja/clientUPCdb/package/console"
	"golang_ninja/clientUPCdb/package/dataEntry"
	"golang_ninja/clientUPCdb/package/sendRequest"
	"io"
	"mime/multipart"
	"net/http"
)

// Saving function implements POST requests to the site with apikey
func Saving() string {
	fmt.Println("enter UPC code")
	upcCode := console.InputDateFromConsole()

	dataUPC := dataEntry.EnterData()

	mp, body := createMultipartData(dataUPC)

	req := createRequest(mp, body, upcCode)

	return sendRequest.SendRequest(req)
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
