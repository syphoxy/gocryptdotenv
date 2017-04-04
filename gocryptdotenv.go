package gocryptdotenv

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/nacl/secretbox"
	"golang.org/x/crypto/pbkdf2"
)

func DecryptFile(filename string, key string) error {
	vars, err := godotenv.Read(filename)
	if err != nil {
		return err
	}

	for i, x := range vars {
		var nonce [24]byte
		var derivedKey [32]byte

		parts := strings.Split(x, " ")

		if len(parts) != 2 {
			log.Fatalln(errors.New(i + " has invalid stored in it"))
		}

		ciphertext, err := hex.DecodeString(parts[1])
		if err != nil {
			return err
		}

		copy(nonce[:], ciphertext[:24])
		salt, err := hex.DecodeString(parts[0])
		if err != nil {
			return err
		}
		derivedKeyBytes := pbkdf2.Key([]byte(key), salt, 4096, 32, sha256.New)
		copy(derivedKey[:], derivedKeyBytes)

		plaintext, success := secretbox.Open([]byte{}, ciphertext[24:], &nonce, &derivedKey)
		if !success {
			return err
		}

		fmt.Print(i, "='", string(plaintext), "'\n")
	}

	return nil
}

func EncryptFile(filename string, key string) error {
	vars, err := godotenv.Read(filename)
	if err != nil {
		return err
	}

	for i, value := range vars {
		var nonce [24]byte
		var derivedKey [32]byte

		_, err := io.ReadFull(rand.Reader, nonce[:])
		if err != nil {
			return err
		}
		salt, err := makeSalt(32)
		if err != nil {
			return err
		}
		derivedKeyBytes := pbkdf2.Key([]byte(key), salt, 4096, 32, sha256.New)
		copy(derivedKey[:], derivedKeyBytes)

		ciphertext := secretbox.Seal(nonce[:], []byte(value), &nonce, &derivedKey)

		fmt.Print(i, "='", hex.EncodeToString(salt), " ", hex.EncodeToString(ciphertext), "'\n")
	}

	return nil
}

func makeSalt(length int) (salt []byte, err error) {
	salt = make([]byte, length)
	if _, err := rand.Read(salt); err != nil {
		return salt, err
	}
	return salt, nil
}
