package cypher

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func show_operations() {
	fmt.Println("Select operation (1/2/3):")
	fmt.Print("1. Encrypt.\n2. Decrypt.\n3. Exit.\n\n")
}

func show_cypher_options() {
	fmt.Println()
	fmt.Println("Select cypher (1/2/3/4):")
	fmt.Print("1. ROT13.\n2. Reverse.\n3. Vigenere. \n4. Back.\n\n\n")
}

func show_decrypt_options() {
	fmt.Println()
	fmt.Println("Select decrypt (1/2/3/4):")
	fmt.Print("1. Decrypt ROT13.\n2. Decrypt Reverse.\n3. Decrypt Vigenere. \n4. Back.\n\n\n")

}

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
	fmt.Println()
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
	fmt.Println()
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

func reverseAlphabet(r rune) rune {

	if r >= 'a' && r <= 'z' {
		return 'a' + ('z' - r)
	}
	if r >= 'A' && r <= 'Z' {
		return 'A' + ('Z' - r)
	}
	return r
}

func encrypt_reverse() {

	fmt.Println()
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

	fmt.Println()
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

func encrypt_vigenere() {
	fmt.Println()
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
