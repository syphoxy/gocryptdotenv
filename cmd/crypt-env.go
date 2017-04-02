package main

import (
	"flag"
	"log"

	gocryptdotenv ".."
)

func main() {
	var filename, operation, key string

	flag.StringVar(&filename, "f", "", "path to .env file (absolute or relative)")
	flag.StringVar(&operation, "t", "", "type of operation (encrypt, decrypt)")
	flag.StringVar(&key, "k", "", "key for decryption or encryption")

	flag.Parse()

	if operation == "decrypt" {
		if err := gocryptdotenv.DecryptFile(filename, key); err != nil {
			log.Fatalln(err)
		}
	} else if operation == "encrypt" {
		if err := gocryptdotenv.EncryptFile(filename, key); err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln("unknown operation:", operation)
	}
}
