package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	response, err := http.Get("https://www.theninjaproject.bbva/login")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		//fmt.Println(response)

		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {

			fmt.Println(string(body))
		}

	}

}
