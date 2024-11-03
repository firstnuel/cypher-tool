package cypher

import "fmt"

func encrypt_reverse() {

	s := getInput("Enter the message to encrypt (reverse): ", []string{})
	if s == "" {
		fmt.Println("No message entered!")
		return
	}

	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		runes[i] = reverseAlphabet(runes[i])
	}
	fmt.Println("Encrypted message using reverse:")
	fmt.Println(string(runes))
	fmt.Println()

}

func decrypt_reverse() {

	s := getInput("Enter the message to encrypt (reverse): ", []string{})
	if s == "" {
		fmt.Println("No message entered!")
		return
	}

	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		runes[i] = reverseAlphabet(runes[i])
	}
	fmt.Println("Decrypted message using reverse:")
	fmt.Println(string(runes))
	fmt.Println()

}
