package main

import "fmt"
import "strings"
import "bufio"
import "os"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		inputFormatted := cleanInput(input)
		if len(inputFormatted) == 0 {
			continue
		}
		fmt.Printf("Your command was: %v\n", inputFormatted[0])
	}
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	words := strings.Fields(text) 
	return words
}
