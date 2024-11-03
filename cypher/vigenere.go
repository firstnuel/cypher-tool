package cypher

import "fmt"

func encrypt_vigenere() {

	text := getInput("Enter the message to encrypt (vigenere): ", []string{})
	if text == "" {
		fmt.Println("No message entered!")
		return
	}
	keyword := getInput("Enter key: ", []string{})
	if keyword == "" {
		fmt.Println("No key entered!")
		return
	}

	encrypted := make([]byte, len(text))
	keywordLength := len(keyword)

	for i := 0; i < len(text); i++ {
		char := text[i]
		shift := (keyword[i%keywordLength] - 'a') % 26

		if char >= 'a' && char <= 'z' {
			encrypted[i] = 'a' + (char-'a'+shift)%26
		} else if char >= 'A' && char <= 'Z' {
			encrypted[i] = 'A' + (char-'A'+shift)%26
		} else {
			encrypted[i] = char
		}
	}

	fmt.Println("Encrypted message using Vigenere and key", keyword+":")
	fmt.Println(string(encrypted))
	fmt.Println()

}

func decrypt_vigenere() {
	fmt.Println()
	text := getInput("Enter the message to decrypt (vigenere): ", []string{})
	if text == "" {
		fmt.Println("No message entered!")
		return
	}
	keyword := getInput("Enter key: ", []string{})
	if keyword == "" {
		fmt.Println("No key entered!")
		return
	}

	decrypted := make([]byte, len(text))
	keywordLength := len(keyword)

	for i := 0; i < len(text); i++ {
		char := text[i]
		shift := (keyword[i%keywordLength] - 'a') % 26

		if char >= 'a' && char <= 'z' {
			decrypted[i] = 'a' + (char-'a'-shift+26)%26
		} else if char >= 'A' && char <= 'Z' {
			decrypted[i] = 'A' + (char-'A'-shift+26)%26
		} else {
			decrypted[i] = char
		}
	}

	fmt.Println("Decrypted message using Vigenere and key", keyword+":")
	fmt.Println(string(decrypted))
	fmt.Println()
}
