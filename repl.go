package main

import (
	"fmt"
)

type Repl struct {
	env *Env
}

func NewRepl() *Repl {
	env := NewEnv(nil)
	SetStandartLibrary(env)

	return &Repl{
		env: env,
	}
}

func (r *Repl) StartRepl(code string) error {
	tokens := Tokenize(code)
	remaining := tokens
	for len(remaining) > 0 {
		value, newRemaining, err := Read(remaining)
		if err != nil {
			return fmt.Errorf("read error: %v", err)
		}
		remaining = newRemaining
		fmt.Println("Read result:", Print(value))

		res, err := Eval(value, r.env)
		if err != nil {
			return fmt.Errorf("eval error: %v", err)
		}
		fmt.Println("Result:", Print(res))
	}
	return nil
}
