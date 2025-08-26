package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/chzyer/readline"
)

func RouteInput(repl *Repl, history string) error {
	os.Mkdir("shine", 0o520)
	os.Create(history)
	config := readline.Config{
		Prompt:       "shine~~> ",
		HistoryFile:  history,
		HistoryLimit: 100,
	}

	rl, err := readline.NewEx(&config)
	if err != nil {
		return err
	}
	defer rl.Close()

	fmt.Println("welcome to Shine Lisp REPL. Type 'exit' to quit.")

	for {
		line, err := rl.Readline()
		if err != nil {
			return err
		}

		line = strings.TrimSpace(line)

		if line == "exit" {
			fmt.Println("shine on your crazy diamond")
			return nil
		}

		if err := repl.StartRepl(line); err != nil {
			fmt.Println(err)
		}

		rl.SaveHistory(line)
	}
}
