package main

import (
	"fmt"
)

type Repl struct {
	env *Env
}

func NewRepl() *Repl {
	env := NewEnv()
	SetStandartLibrary(env)

	return &Repl{
		env: env,
	}
}

func (r *Repl) StartRepl(code string) error {
	tokens := Tokenize(code)
	value, _, err := Read(tokens)
	if err != nil {
		return err
	}

	res, err := Eval(value, r.env)
	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}
