package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	message := []byte("hello world\n")

	_, err = file.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}
