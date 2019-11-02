package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	//create a map key:string value:string
	m := make(map[string]string)

	//prompt for name and store in map

	fmt.Println("Please enter name:\n")

	reader_name := bufio.NewReader(os.Stdin)
	receiver_name, _ := reader_name.ReadString('\n')

	m["name"] = strings.TrimSuffix(receiver_name, "\n")

	//prompt for address nd store in map

	fmt.Println("Please enter address:\n")

	reader_address := bufio.NewReader(os.Stdin)
	receiver_address, _ := reader_address.ReadString('\n')

	m["address"] = strings.TrimSuffix(receiver_address, "\n")

	//marshall the map into a byte array containing json representation of map

	bajson, _ := json.Marshal(m)

	fmt.Println(string(bajson))

}
