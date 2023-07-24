package encoders

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
)

func Encrypto(plaintext, keystring []byte) ([]byte, error) {

	// Key
	key := sha256.Sum256(keystring)

	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, fmt.Errorf("NewCipher error:%v", err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	iv := ciphertext[:aes.BlockSize]

	// 16 random bytes
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, fmt.Errorf("make random error:%v", err)
	}

	// Return an encrypted stream
	stream := cipher.NewCFBEncrypter(block, iv)

	// Encrypt bytes from plaintext to ciphertext
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}
