package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var OutputLevel string = "base"

func runWithFile(filename string, repl *Repl) error {
	code, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		return err
	}
	fmt.Println("Input from file:", string(code))

	err = repl.StartRepl(string(code))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}
	return nil
}

func runForLibs(dirs []string, repl *Repl) error {
	for _, dir := range dirs {
		files, err := os.ReadDir(dir)
		if err != nil {
			fmt.Println(err)

			return err
		}

		for _, file := range files {
			code, err := os.ReadFile(filepath.Join(dir, file.Name()))
			if err != nil {
				fmt.Println(err)
				return err
			}

			if err = repl.StartRepl(string(code)); err != nil {
				fmt.Println(err)
				return err
			}
		}
	}

	return nil
}

func main() {
	var (
		dirs []string
		file string
	)

	repl := NewRepl()

	for i, arg := range os.Args {
		switch arg {
		case "--output-level":
			OutputLevel = os.Args[i+1]

		case "--lib":
			dirs = append(dirs, os.Args[i+1])

		case "--file":
			file = os.Args[i+1]
		}
	}

	if len(dirs) != 0 {
		if err := runForLibs(dirs, repl); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	if file != "" {
		if err := runWithFile(file, repl); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if err := RouteInput(repl, "shine/history"); err != nil {
		fmt.Println(err)
	}
}
