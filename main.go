package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	repl := NewRepl()

	for {
		fmt.Print("shine~~>")
		code, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)

			return
		}

		if code == "" {
			continue
		}
		if strings.HasPrefix(code, "exit") {
			fmt.Println("Shine on your crazy diamond")
			break
		}

		if err := repl.StartRepl(code); err != nil {
			fmt.Println(err)
		}
	}
}
