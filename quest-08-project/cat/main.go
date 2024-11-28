package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		files, err := os.ReadDir(".")
		if err != nil {
			fmt.Println("Error reading directory:", err)
			return
		}
		for _, file := range files {
			if !file.IsDir() && len(file.Name()) > 3 && file.Name()[len(file.Name())-3:] != ".go" {
				args = append(args, file.Name())
			}
		}
	}

	if len(args) == 0 {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			fmt.Println(s.Text())
		}
		if err := s.Err(); err != nil {
			fmt.Println("Error reading standard input:", err)
		}
		return
	}
	for _, filename := range args {
		func() {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			defer file.Close()
			s := bufio.NewScanner(file)
			for s.Scan() {
				fmt.Println(s.Text())
			}
			if err := s.Err(); err != nil {
				fmt.Println("Error reading file:", err)
			}
		}()
	}
} 
