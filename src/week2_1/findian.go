package main

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
	"strings"
)

func main() {

	fmt.Println("Please, enter a string...\n")
	reader := bufio.NewReader(os.Stdin)
	receiver, _ := reader.ReadString('\n')
	//fmt.Println(receiver)
	receiver = strings.TrimSpace(strings.ToLower(receiver))
	/*fmt.Println(receiver)
	fmt.Println(strconv.FormatBool(strings.HasPrefix(receiver, "i")))
	fmt.Println(receiver)
	fmt.Println(strconv.FormatBool(strings.HasSuffix(receiver, "n")))
	fmt.Println(strconv.FormatBool(strings.Contains(receiver, "a")))*/
	if strings.HasPrefix(receiver, "i") && strings.HasSuffix(receiver, "n") && strings.Contains(receiver, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
