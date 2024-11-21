package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Messaggio generico ricevuto dall'estensione
type Message struct {
	Text string `json:"text"`
}

// Risposta da inviare all'estensione
type Response struct {
	Response string `json:"response"`
	Data     string `json:"data"`
}

func readMessage() (Message, error) {
	var message Message

	// Leggi i primi 4 byte per la lunghezza del messaggio
	var length uint32
	if err := binary.Read(os.Stdin, binary.LittleEndian, &length); err != nil {
		return message, err
	}

	// Leggi il messaggio in base alla lunghezza
	buffer := make([]byte, length)
	if _, err := io.ReadFull(os.Stdin, buffer); err != nil {
		return message, err
	}

	// Decodifica il JSON
	if err := json.Unmarshal(buffer, &message); err != nil {
		return message, err
	}

	return message, nil
}

func sendMessage(response Response) error {
	// Codifica la risposta in JSON
	data, err := json.Marshal(response)
	if err != nil {
		return err
	}

	// Scrivi la lunghezza della risposta (4 byte, Little Endian)
	if err := binary.Write(os.Stdout, binary.LittleEndian, uint32(len(data))); err != nil {
		return err
	}

	// Scrivi il contenuto della risposta
	_, err = os.Stdout.Write(data)
	return err
}

func main() {
	for {
		// Leggi un messaggio dall'estensione
		message, err := readMessage()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading message: %v\n", err)
			break
		}

		// Gestisci il messaggio e crea una risposta
		response := Response{
			Response: "Message received",
			Data:     message.Text,
		}

		// Invia la risposta all'estensione
		if err := sendMessage(response); err != nil {
			fmt.Fprintf(os.Stderr, "Error sending message: %v\n", err)
			break
		}
	}
}
