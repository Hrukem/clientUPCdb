// Package sendRequest implements sending a request
// to the site and returns a response
package sendRequest

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendRequest(req *http.Request) string {
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

	return string(answer)
}
