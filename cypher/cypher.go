package cypher

import "fmt"

func CypherTool() {

	fmt.Println()
	fmt.Println("Welcome to the Cypher Tool!")

	choice := ""

	show_operations()

	choice = getInput("Enter choice: ", []string{"1", "2"})
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
		fmt.Println("Pussy")
	}

}
