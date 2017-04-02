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

func DecryptFile(filename string, key string) (err error) {
	vars, err := godotenv.Read(filename)
	if err != nil {
		return
	}

	for i, x := range vars {
		var nonce [24]byte
		var derivedKey [32]byte

		parts := strings.Split(x, " ")

		if len(parts) != 2 {
			log.Fatalln(errors.New(i + " has invalid stored in it"))
		}

		ciphertext, _ := hex.DecodeString(parts[1]) // TODO check errors here

		copy(nonce[:], ciphertext[:24])
		salt, _ := hex.DecodeString(parts[0]) // TODO check errors here
		derivedKeyBytes := pbkdf2.Key([]byte(key), salt, 4096, 32, sha256.New)
		copy(derivedKey[:], derivedKeyBytes)

		plaintext, _ := secretbox.Open([]byte{}, ciphertext[24:], &nonce, &derivedKey)

		fmt.Print(i, "='", string(plaintext), "'\n")
	}

	return nil
}

func EncryptFile(filename string, key string) (err error) {
	vars, err := godotenv.Read(filename)
	if err != nil {
		return
	}

	for i, value := range vars {
		var nonce [24]byte
		var derivedKey [32]byte

		_, _ = io.ReadFull(rand.Reader, nonce[:]) // TODO check for errors here
		salt, _ := makeSalt(32)                   // TODO check for errors here
		derivedKeyBytes := pbkdf2.Key([]byte(key), salt, 4096, 32, sha256.New)
		copy(derivedKey[:], derivedKeyBytes)

		ciphertext := secretbox.Seal(nonce[:], []byte(value), &nonce, &derivedKey) // TODO check for errors here

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
