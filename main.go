package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	repl := NewRepl()

	if len(os.Args) > 2 && os.Args[1] == "-f" {
		filename := os.Args[2]
		code, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", filename, err)
			os.Exit(1)
		}
		fmt.Println("Input from file:", string(code))
		err = repl.StartRepl(string(code))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Shine Lisp REPL! Type 'exit' to quit.")

	for {
		fmt.Print("shine~~> ")
		code, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			break
		}

		code = strings.TrimSpace(code)
		if code == "" {
			continue
		}
		if code == "exit" {
			fmt.Println("Shine on Your Crazy Diamond!")
			break
		}

		fmt.Println("Input:", code)
		err = repl.StartRepl(code)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
