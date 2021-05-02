package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	
	var dataToEncrypt string
	var encryptedData string
	var decryptedData string

	hash := verifyFile("myHash.txt")
	
	in := bufio.NewReader(os.Stdin)

	fmt.Print("Dado a ser encriptografado: ")
	line, _ := in.ReadString('\n')
	dataToEncrypt = line[0 : len(line)-2]

	encryptedData = encrypt(dataToEncrypt, hash)
	fmt.Println("Dado encriptografado:", encryptedData)

	decryptedData = decrypt(encryptedData, hash)
	fmt.Println("Dado descriptografado: ", decryptedData)
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

func verifyFile(filename string) (hash string){
	file, err := os.Open(filename)
	if err != nil {
		f, err := os.Create(filename)
		if err != nil {
			fmt.Println(err.Error())
		}
		f.Close()
		
		return createHash()
	}
	file.Close()

	content, err := ioutil.ReadFile(filename);
	if err != nil{
		fmt.Println(err)
	}
	
	return string(content)
}

func createHash() (hash string){
	
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil{
		panic(err.Error())
	}

	hasher := hex.EncodeToString(bytes)

	f, err := os.OpenFile("myHash.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer f.Close()
	f.WriteString(hasher + "\n")

	fmt.Println("Hash Criado com sucesso.")

	return hasher
}

/*
!important

key (32-bytes for AES-256 encryption)
nonce (a random no., only used only once)
data to encrypt - plaintext in this example

*/
