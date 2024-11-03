package cypher

import "fmt"

func CypherTool() {

	fmt.Println("\nWelcome to the Cypher Tool!")
	fmt.Println()

	choice := ""

	for {
		show_operations()
		choice = getInput("Enter choice: ", []string{"1", "2", "3"})

		switch choice {
		case "1":
			show_cypher_options()
			cypherChoice := getInput("Enter cypher choice: ", []string{"1", "2", "3", "4"})

			if !handleCypherChoice(cypherChoice) {
				fmt.Println("\nThanks for using our group's cypher tool!")
				return
			}

		case "2":
			show_decrypt_options()
			decryptChoice := getInput("Enter decrypt choice: ", []string{"1", "2", "3", "4"})

			if !handleDecryptChoice(decryptChoice) {
				fmt.Println("\nThanks for using our group's cypher tool!")
				return
			}

		case "3":
			fmt.Println("\nThanks for using our group's cypher tool!")
			return
		}
	}
}
