# Cypher Tool

## What does the tool do?

The **Cypher Tool** is an interactive command-line application written in Go that allows users to encrypt and decrypt text using a variety of simple cipher techniques. This tool provides basic encryption options for encoding messages and secure communication. It currently supports three types of ciphers: **ROT13**, **Reverse**, and **Vigenère**. Users can choose to either encrypt or decrypt messages and can easily navigate back and forth within the menu.

## Tool Usage

To use the tool, simply run it from the command line:

```bash
go run .
```

Example Interaction
Here’s a sample interaction with the tool:

```bash
Welcome to the Cypher Tool!

Select operation (1/2/3):
1. Encrypt.
2. Decrypt.
3. Exit.

Enter choice: 1

Select cypher (1/2/3/4):
1. ROT13.
2. Reverse.
3. Vigenere. 
4. Back.

Enter cypher choice: 1
Enter the text to encrypt: Hello World
Encrypted text: Uryyb Jbeyq

Would you like to perform another operation? (y/n)
```
In this example, the user selects the encryption operation, chooses ROT13 as the cipher, and inputs the text to be encrypted.

## Navigation
- Main Menu:

    - Choose `1` to encrypt a message.
    - Choose `2` to decrypt a message.
    - Choose `3` to exit the tool.

- Cypher Selection (after choosing encrypt or decrypt):

    - `1` for ROT13
    - `2` for Reverse
    - `3` for Vigenère
    - `4` to return to the main menu

The tool then prompts the user to enter text for encryption or decryption based on the selected operation.

### Explanation of the Ciphers Used

1. ROT13: ROT13 (Rotate by 13 places) is a simple substitution cipher that replaces each letter in a message with the letter 13 places after it in the alphabet. It works as both an encryption and decryption method since applying ROT13 twice returns the original text.

    - Example: `"Hello"` becomes `"Uryyb"`.

2. Reverse: The Reverse cipher reverses the characters in the text to create an encoded message. This cipher is easy to decrypt, as reversing the text again restores the original message.

    - Example: `"Hello"` becomes `"olleH"`.

3. Vigenère Cipher: The Vigenère cipher is a method of encrypting alphabetic text by using a keyword. Each letter in the text is shifted based on the corresponding letter in the keyword, creating a more complex pattern of substitutions than ROT13.

    - Example: Using the keyword `"KEY"` to encrypt `"HELLO"`:
        - Plaintext: `H E L L O`
        - Key: `K E Y K E`
        - Ciphertext: `R I J V S`
        
The Cypher Tool provides a simple, accessible way to practice and explore basic encryption methods.