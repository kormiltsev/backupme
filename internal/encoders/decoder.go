package encoders

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"fmt"
)

func Decrypto(ciphertext, keystring []byte) ([]byte, error) {

	// Key
	key := sha256.Sum256(keystring)

	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, fmt.Errorf("AES error:%v", err)
	}

	// if the text is too small, then it is incorrect
	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("income []byte too short:%v", err)
	}

	iv := ciphertext[:aes.BlockSize]

	// Remove the IV from the ciphertext
	ciphertext = ciphertext[aes.BlockSize:]

	// Return a decrypted stream
	stream := cipher.NewCFBDecrypter(block, iv)

	// Decrypt bytes from ciphertext
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}
