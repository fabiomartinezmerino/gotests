package main

import (
	"fmt"
	//"io/ioutil"
	"net/http"
	"os"
)

func main() {

	response, err := http.Get("http://www.google.com")

	if err != nil {

		fmt.Println(err)
		os.Exit(1)

	}

	buffSlice := make([]byte, 9999)

	response.Body.Read(buffSlice)
	defer response.Body.Close()
	fmt.Println(string(buffSlice))

}
