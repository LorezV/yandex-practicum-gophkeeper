package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const keyLength = 32

var errEmptyFile = errors.New("empty file has been given")

func getKeyFromPass(keyString string) []byte {
	key := []byte(keyString)

	if len(key) < keyLength {
		for {
			key = append(key, key[0])
			if len(key) == keyLength {
				break
			}
		}
	} else if len(key) > keyLength {
		key = key[:keyLength]
	}

	return key
}

func Encrypt(keyString, stringToEncrypt string) string {
	if stringToEncrypt == "" {
		return stringToEncrypt
	}
	cipherBlock, err := aes.NewCipher(getKeyFromPass(keyString))
	if err != nil {
		log.Fatalf("Encrypt - aes.NewCipher - %v", err)
	}

	aead, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		log.Fatalf("Encrypt - cipher.NewGCM - %v", err)
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("Encrypt - io.ReadFull(rand.Reader, nonce) - %v", err)
	}

	return base64.URLEncoding.EncodeToString(aead.Seal(nonce, nonce, []byte(stringToEncrypt), nil))
}

func Decrypt(keyString, encryptedString string) (decryptedString string) {
	if encryptedString == "" {
		return encryptedString
	}
	encryptData, err := base64.URLEncoding.DecodeString(encryptedString)
	if err != nil {
		log.Fatal(err)
	}

	cipherBlock, err := aes.NewCipher(getKeyFromPass(keyString))
	if err != nil {
		log.Fatalf("Decrypt - aes.NewCipher - %v", err)
	}

	aead, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		log.Fatalf("Decrypt - cipher.NewGCM - %v", err)
	}

	nonceSize := aead.NonceSize()
	if len(encryptData) < nonceSize {
		log.Fatalf("Decrypt - aead.NonceSize - %v", err)
	}

	nonce, cipherText := encryptData[:nonceSize], encryptData[nonceSize:]
	plainData, err := aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		log.Fatalf("Decrypt - aead.Open - %v", err)
	}

	return string(plainData)
}

func EncryptFile(keyString, inputFilePath, outputFilePath string) error {
	fi, err := os.Stat(inputFilePath)
	if err != nil {
		return fmt.Errorf("EncryptFile - os.Stat - %w", err)
	}
	if size := fi.Size(); size == 0 {
		return fmt.Errorf("EncryptFile - fi.Size - %w", errEmptyFile)
	}
	cipherBlock, err := aes.NewCipher(getKeyFromPass(keyString))
	if err != nil {
		return fmt.Errorf("EncryptFile - aes.NewCipher - %w", errEmptyFile)
	}

	aead, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return fmt.Errorf("EncryptFile - cipher.NewGCM - %w", err)
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return fmt.Errorf("EncryptFile - io.ReadFull(rand.Reader, nonce) - %w", err)
	}
	inputData, err := os.ReadFile(inputFilePath)
	if err != nil {
		return fmt.Errorf("EncryptFile - os.ReadFile - %w", err)
	}

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("EncryptFile - os.Create - %w", err)
	}

	defer outputFile.Close()

	_, err = outputFile.WriteString(
		base64.URLEncoding.EncodeToString(aead.Seal(nonce, nonce, inputData, nil)))
	if err != nil {
		return fmt.Errorf("EncryptFile - base64.URLEncoding.EncodeToString - %w", err)
	}

	return nil
}

func DecryptFile(keyString, encryptedPath, decryptedFilePath string) error {
	encryptedData, err := os.ReadFile(encryptedPath)
	if err != nil {
		return fmt.Errorf("DecryptFile - os.ReadFile - %w", err)
	}

	encryptData, err := base64.URLEncoding.DecodeString(string(encryptedData))
	if err != nil {
		return fmt.Errorf("DecryptFile - base64.URLEncoding.DecodeString - %w", err)
	}

	cipherBlock, err := aes.NewCipher(getKeyFromPass(keyString))
	if err != nil {
		return fmt.Errorf("DecryptFile - aes.NewCipher - %w", err)
	}

	aead, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return fmt.Errorf("DecryptFile - cipher.NewGCM - %w", err)
	}

	nonceSize := aead.NonceSize()
	if len(encryptData) < nonceSize {
		return fmt.Errorf("DecryptFile - aead.NonceSize - %w", err)
	}

	nonce, cipherText := encryptData[:nonceSize], encryptData[nonceSize:]
	plainData, err := aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return fmt.Errorf("DecryptFile - aead.Open - %w", err)
	}
	outputFile, err := os.Create(decryptedFilePath)
	if err != nil {
		return fmt.Errorf("EncryptFile - os.Create - %w", err)
	}

	defer outputFile.Close()
	_, err = outputFile.WriteString(
		string(plainData))
	if err != nil {
		return fmt.Errorf("EncryptFile - outputFile.WriteString - %w", err)
	}

	return nil
}
