package cypher

import "fmt"

func shift13(r rune) rune {

	step := 13

	if r >= 'a' && r <= 'z' {
		return 'a' + (r-'a'+rune(step))%26
	}
	if r >= 'A' && r <= 'Z' {
		return 'A' + (r-'A'+rune(step))%26
	}
	return r
}

func encrypt_rot13() {

	s := getInput("Enter the message to encrypt (rot13): ", []string{})
	if s == "" {
		fmt.Println("No message entered!")
		return
	}

	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		runes[i] = shift13(runes[i])
	}
	fmt.Println("Encrypted message using rot13:")
	fmt.Println(string(runes))
	fmt.Println()
}

func decrypt_rot13() {

	s := getInput("Enter the message to decrypt (rot13): ", []string{})
	if s == "" {
		fmt.Println("No message entered!")
		return
	}

	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		runes[i] = shift13(runes[i])
	}
	fmt.Println("Decrypted message using rot13:")
	fmt.Println(string(runes))
	fmt.Println()
}
