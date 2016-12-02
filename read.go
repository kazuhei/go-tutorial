package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	message := make([]byte, 12)
	_, err = file.Read(message)
	if err != nil { // エラー処理
		log.Fatal(err)
	}

	// 文字列にして表示
	fmt.Print(string(message))
}
