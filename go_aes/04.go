package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func main() {
	//fmt.Println(encrypt("LucasAmado", "52fdfc072182654f163f5f0f9a621d729566c74d10037c4d7bbb0407d1e2c649"))
	fmt.Println(decrypt("0000000000000000000000002b0fc2d505ae6889c718a84ffbb9254c531ce87ce2761a9d3d84", "52fdfc072182654f163f5f0f9a621d729566c74d10037c4d7bbb0407d1e2c649"))
}

func encrypt(stringToEncrypt string, keyString string) (encryptedString string) {

	//key := hex.EncodeToString(keyString)
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)

	//Criando um novo Cipher Block a partir da chave que gerei
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
	}

	//Criando um novo GCM
	//man -> //https://golang.org/pkg/crypto/cipher/#NewGCM

	aesGCM, err := cipher.NewGCM(block)

	//Criando um nonce (número randômico usado uma vez)
	nonce := make([]byte, aesGCM.NonceSize())

	//Encriptando o dado usando o GCM

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)

	/* fmt.Println("Chave utilizada : ", key)
	fmt.Println("Cipher a partir da chave : ", block)
	fmt.Println("GCM novo criado : ", aesGCM)
	fmt.Println("Nonce criado: ", nonce)
	fmt.Println("Dado criado encryptografado: ", ciphertext) */

	return fmt.Sprintf("%x", ciphertext)

}

func decrypt(encryptedString string, keyString string) (decryptedString string) {

	//convertendo a chave e encriptografando os data para bytes
	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	//criando um novo cipher a partir da chave
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//criando um novo gcm
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//pegando o nonce size e extraindo do prefixo do dado encriptografado
	//este é um passo importante pois não podemos descriptografat o dado sem o nonce correto

	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Descriptografando o dado usando GCM Open

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}

/*
!important

key (32-bytes for AES-256 encryption)
nonce (a random no., only used only once)
data to encrypt - plaintext in this example

*/
