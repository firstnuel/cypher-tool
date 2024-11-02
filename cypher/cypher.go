package cypher

import "fmt"

func CypherTool() {

	fmt.Println()
	fmt.Println("Welcome to the Cypher Tool!")
	fmt.Println()

	choice := ""

	for {
		show_operations()

		choice = getInput("Enter choice: ", []string{"1", "2", "3"})

		if choice == "1" {
			show_cypher_options()
			cypherChoice := getInput("Enter cypher choice: ", []string{"1", "2", "3"})

			if cypherChoice == "1" {
				encrypt_rot13()
			}
			if cypherChoice == "2" {
				encrypt_reverse()
			}
			if cypherChoice == "3" {
				encrypt_vigenere()
			}
		}

		if choice == "2" {
			show_decrypt_options()
			decryptChoice := getInput("Enter decrypt choice: ", []string{"1", "2", "3"})

			if decryptChoice == "1" {
				decrypt_rot13()
			}
			if decryptChoice == "2" {
				decrypt_reverse()
			}
			if decryptChoice == "3" {
				decrypt_vigenere()
			}

		}

		if choice == "3" {
			break
		}
	}
	fmt.Println()
	fmt.Println("Thanks for using our groups cypher tool!")
	fmt.Println()
}
