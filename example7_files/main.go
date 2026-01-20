package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func example7_files() {
	fmt.Println("\n=== Example 7: File Handling ===")

	data := "Genesis Block Data"
	err := ioutil.WriteFile("block.txt", []byte(data), 0644)
	if err != nil {
		fmt.Println("Write error: ", err)
		return
	}
	fmt.Println("Wrote to block.txt")

	content, err := ioutil.ReadFile("block.txt")
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}
	fmt.Printf("Read from file: %s\n", string(content))

	os.Remove("block.txt")
}

func main() {
	example7_files()
}
