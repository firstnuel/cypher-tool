package cypher

import (
	"fmt"
	"strings"
)

func show_operations() {
	fmt.Println("Select operation (1/2/3):")
	fmt.Print("1. Encrypt.\n2. Decrypt.\n3. Exit.\n\n")
}

func show_cypher_options() {
	fmt.Println()
	fmt.Println("Select cypher (1/2/3/4):")
	fmt.Print("1. ROT13.\n2. Reverse.\n3. Vigenere. \n4. Back.\n\n")
}

func show_decrypt_options() {
	fmt.Println()
	fmt.Println("Select decrypt (1/2/3/4):")
	fmt.Print("1. Decrypt ROT13.\n2. Decrypt Reverse.\n3. Decrypt Vigenere. \n4. Back.\n\n")

}

func shouldContinue() bool {
	toContinue := getInput("Perform another operation? (y/n): ", []string{"y", "n", "Y", "N"})
	return strings.ToLower(toContinue) == "y"
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
		return true
	}
	return shouldContinue()
}
