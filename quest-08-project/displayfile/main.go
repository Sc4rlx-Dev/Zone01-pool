package main

import (
	"fmt"
	"os"
)
func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}
	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}
	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	content := make([]byte, stat.Size())
	_, err = file.Read(content)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Print(string(content))
}
