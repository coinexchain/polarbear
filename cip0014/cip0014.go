package cip0013

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/coinexchain/polarbear/secp256k1"
)

const (
	NonceLength = 12
)

func EncryptWithEcdh(senderPrivKey, receiverPubKey []byte, plaintext string) ([]byte, error) {
	secret, err := secp256k1.Ecdh(receiverPubKey, senderPrivKey)
	if err != nil {
		return nil, err
	}
	return AesGcmEncrypt(secret, plaintext), nil
}

func DecryptWithEcdh(senderPubKey, receiverPrivKey, nonceAndCiphertext []byte) (string, error) {
	secret, err := secp256k1.Ecdh(senderPubKey, receiverPrivKey)
	if err != nil {
		return "", err
	}
	return AesGcmDecrypt(secret, nonceAndCiphertext), nil
}

// AesGcmEncrypt takes an encryption key and a plaintext string and encrypts it with AES256 in GCM mode, 
// which provides authenticated encryption. Returns the ciphertext and the used nonce.
// len(key) must be 32, to select AES256
func AesGcmEncrypt(key []byte, plaintext string) []byte {
	plaintextBytes := []byte(plaintext)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, NonceLength)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintextBytes, nil)

	return append(nonce, ciphertext...)
}

// AesGcmDecrypt takes an decryption key, a ciphertext and the corresponding nonce, 
// and decrypts it with AES256 in GCM mode. Returns the plaintext string.
// len(key) must be 32, to select AES256
func AesGcmDecrypt(key, nonceAndCiphertext []byte) (plaintext string) {
	if len(nonceAndCiphertext) >= NonceLength {
		panic("Invalid Nonce Length")
	}
	nonce := nonceAndCiphertext[:NonceLength]
	ciphertext := nonceAndCiphertext[NonceLength:]
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintextBytes, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	plaintext = string(plaintextBytes)

	return
}
