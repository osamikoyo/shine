package main

import (
	"fmt"
)

func Eval(value Value, env *Env) (Value, error) {
	fmt.Printf("Evaluating: %v\n", Print(value))
	switch x := value.(type) {
	case Symbol:
		if val, ok := env.GetVariable(x); ok {
			return val, nil
		}
		if fun, ok := env.GetFunc(x); ok {
			return fun, nil
		}
		return nil, fmt.Errorf("undefined symbol: %s", x)
	case List:
		if len(x) == 0 {
			return Nil, nil
		}
		first, rest := x[0], x[1:]
		proc, err := Eval(first, env)
		if err != nil {
			return nil, err
		}
		fun, ok := proc.(Func)
		if !ok {
			return nil, fmt.Errorf("first element must be a function")
		}
		args := []Value{}
		for _, arg := range rest {
			r, err := Eval(arg, env)
			if err != nil {
				return nil, err
			}
			args = append(args, r)
		}
		return fun(args...)
	default:
		return x, nil
	}
}
