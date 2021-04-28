/*
	AES --> Advanced Encryption Standard
*/

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	fmt.Println("Encriptar Mensagem")

	text := []byte("My Super Secret")
	key := []byte("passwordsecret11")

	//generate a new aes cipher
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
	}

	//gcm -> um modo de operação para chaves simetricas em blocos ciphers
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err.Error())
	}

	//criando um novo array de byte no tamanho de nonce
	nonce := make([]byte, gcm.NonceSize())

	// populando nossa sequencia nonce com dados randômicos
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	//fmt.Println(gcm.Seal(nonce, nonce, text, nil))
	err = ioutil.WriteFile("myfile01.data", gcm.Seal(nonce, nonce, text, nil), 0777)
}
