package web_programming

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func init() {
	fmt.Println("===== Downloading a Web Page from Internet =====")

	url := "http://golang.org"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	html, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println(string(html))

}
