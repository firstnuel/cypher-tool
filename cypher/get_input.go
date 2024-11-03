package cypher

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, expected []string) string {
	if len(expected) == 0 {

		fmt.Printf("%v", prompt)
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			return ""
		}
		return strings.TrimSpace(text)
	}

	text := ""
	var err error

	for {
		fmt.Printf("%v", prompt)
		reader := bufio.NewReader(os.Stdin)
		text, err = reader.ReadString('\n')
		if err != nil {
			return ""
		}
		text = strings.TrimSpace(text)

		if includesString(expected, text) {
			break
		}
		fmt.Println("Invalid choice!")
	}

	return text
}

func includesString(list []string, value string) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}
