//02 decrypt

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
)

func main() {

	fmt.Println("Decriptar Programa")

	key := []byte("passwordsecret11")
	ciphertext, err := ioutil.ReadFile("myfile01.data")
	if err != nil {
		fmt.Println(err.Error())
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(plaintext))

}
