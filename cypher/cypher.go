package cypher

import "fmt"

func CypherTool() {

	fmt.Println("\nWelcome to the Cypher Tool!")
	fmt.Println()

	choice := ""
	show_operations()

	for {
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

func handleCypherChoice(choice string) bool {
	switch choice {
	case "1":
		encrypt_rot13()
	case "2":
		encrypt_reverse()
	case "3":
		encrypt_vigenere()
	case "4":
		show_operations()
	default:
		return true
	}
	return shouldContinue()
}

func handleDecryptChoice(choice string) bool {
	switch choice {
	case "1":
		decrypt_rot13()
	case "2":
		decrypt_reverse()
	case "3":
		decrypt_vigenere()
	case "4":
		show_operations()
	default:
		return true
	}
	return shouldContinue()
}
