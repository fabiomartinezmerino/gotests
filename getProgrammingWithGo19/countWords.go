package main

import "fmt"
import "strings"
import "bufio"
import "log"
import "os"

func main() {
	/*text := "as far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendor od sunshine in which he walked and walked and walked and walked..."
	 */

	//  wl := strings.Fields(text)

	text, err := readFile("/Users/Alicia/Documents/gotestslice.txt")
	if err != nil {
		log.Fatalf("readFile: %s", err)
	}
	s := strings.Join(text, "")
	wl := strings.Fields(s)

	dict := make(map[string]int64, len(wl))

	for _, word := range wl {
		dict[strings.Trim(strings.ToLower(word), `.,-`)]++
	}

	for word, times := range dict {
		fmt.Printf("La palabra %v ocurre %v veces \n", word, times)

	}
}

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()

}
