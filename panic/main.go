package main

import (
	"fmt"
	"os"
)

func ReadFile() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from error:", r)
		}
	}()

	_, err := os.Open("./settings.txt")
	if err != nil {
		panic("File not exists")
	}
}

func main() {

	ReadFile()

	fmt.Println("Fim")

}
