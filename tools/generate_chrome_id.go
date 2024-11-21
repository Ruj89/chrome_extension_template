package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func calculateExtID(path string) string {
	// Calcolo SHA-256
	hasher := sha256.New()
	hasher.Write([]byte(path))
	hashBytes := hasher.Sum(nil)
	hashHex := hex.EncodeToString(hashBytes)

	// Conversione in extid
	extid := ""
	for i := 0; i < 32; i++ {
		// Converte il carattere esadecimale in un valore intero
		value := 0
		if hashHex[i] >= '0' && hashHex[i] <= '9' {
			value = int(hashHex[i] - '0')
		} else if hashHex[i] >= 'a' && hashHex[i] <= 'f' {
			value = int(hashHex[i]-'a') + 10
		}
		// Somma 'a' per creare un carattere alfabetico
		extid += string('a' + value)
	}
	return extid
}

func main() {
	// Verifica degli argomenti
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <path>")
		os.Exit(1)
	}

	// Prendi il valore di 'path' dalla riga di comando
	path := os.Args[1]
	extid := calculateExtID(path)
	fmt.Println(extid)
}
