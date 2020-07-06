package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

// Needed because the cipher accepts only 16, 32 or 48 bytes key
func reduceTo16bytes(s string) ([]byte, error) {
	hasher := md5.New()
	_, err := hasher.Write([]byte(s))
	if err != nil {
		return nil, err
	}
	return hasher.Sum(nil), nil
}

func Encrypt(key, text string) (string, error) {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	cipherKey, err := reduceTo16bytes(key)
	if err != nil {
		return "", err
	}
	// cipherKey must be 16, 32 or 48 bytes long
	block, err := aes.NewCipher(cipherKey)
	if err != nil {
		return "", err
	}

	plaintext := []byte(text)
	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to be secure.
	return fmt.Sprintf("%x", ciphertext), nil
}

func Decrypt(key, ciphertext string) (string, error) {
	cipherKey, err := reduceTo16bytes(key)
	if err != nil {
		return "", err
	}
	text, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(cipherKey)
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(text) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(text, text)
	return string(text), nil
}
