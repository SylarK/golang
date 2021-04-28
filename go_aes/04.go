package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
)

func main() {

	bytes := make([]byte, 32) // gerando um número randomico de 32 bytes para AES-256
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	key := hex.EncodeToString(bytes)
	plaintext := []byte(key)
	//plaintext := []byte(stringToEncrypt)

	f, _ := os.Create("myfile04.txt")

	defer f.Close()
	f.Write(plaintext)

	//Criando um novo Cipher Block a partir da chave que gerei
	block, err := aes.NewCipher(plaintext)
	if err != nil {
		fmt.Println(err.Error())
	}

	//Criando um novo GCM

	aesGCM, err := cipher.NewGCM(block)

	//Criando um nonce (número randômico usado uma vez)
	nonce := make

	fmt.Println(key)
}

/*
!important

key (32-bytes for AES-256 encryption)
nonce (a random no., only used only once)
data to encrypt - plaintext in this example

*/
