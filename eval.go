package main

import (
	"fmt"
	"os"
)

func Eval(value Value, env *Env) (Value, error) {
	fmt.Println("Evaluating:", Print(value))
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

		if sym, ok := first.(Symbol); ok && sym == "quote" {
			if len(rest) != 1 {
				return nil, fmt.Errorf("quote expects exactly one argument")
			}
			return rest[0], nil
		}

		if sym, ok := first.(Symbol); ok && sym == "import" {
			switch v := rest[0].(type) {
			case String:
				code, err := os.ReadFile(string(v))
				if err != nil {
					return nil, err
				}

				repl := NewRepl(env)
				if err = repl.StartRepl(string(code)); err != nil {
					return nil, err
				}

				return Nil, err
			case List:
				for _, arg := range v {
					path, ok := arg.(String)
					if !ok {
						return nil, fmt.Errorf("expects string, got: %v", arg)
					}

					code, err := os.ReadFile(string(path))
					if err != nil {
						return nil, err
					}

					repl := NewRepl(env)
					if err = repl.StartRepl(string(code)); err != nil {
						return nil, err
					}
				}
			}
		}
		if sym, ok := first.(Symbol); ok && sym == "define" {
			if len(rest) != 2 {
				return nil, fmt.Errorf("define expects exactly two arguments")
			}
			sym, ok := rest[0].(Symbol)
			if !ok {
				return nil, fmt.Errorf("define first argument must be a symbol")
			}
			val, err := Eval(rest[1], env)
			if err != nil {
				return nil, err
			}
			env.SetVariable(sym, val)
			return val, nil
		}

		if sym, ok := first.(Symbol); ok && sym == "let" {
			if len(rest) > 2 {
				return nil, fmt.Errorf("let expects exacly 2 arguments")
			}

			bindings, ok := rest[0].(List)
			if !ok {
				return nil, fmt.Errorf("let first argument must to be list")
			}

			newEnv := NewEnv(env)

			for _, binding := range bindings {
				bindList, ok := binding.(List)
				if !ok || len(bindList) != 2 {
					return nil, fmt.Errorf("binding must to be list")
				}

				s, ok := bindList[0].(Symbol)
				if !ok {
					return nil, fmt.Errorf("bindlist first argument must be symbol")
				}

				val, err := Eval(bindList[1], newEnv)
				if err != nil {
					return nil, err
				}

				newEnv.SetVariable(s, val)
			}

			return Eval(rest[1], newEnv)
		}

		if sym, ok := first.(Symbol); ok && sym == "lambda" {
			if len(rest) != 2 {
				return nil, fmt.Errorf("lambda expects exactly two arguments: parameters and body")
			}
			params, ok := rest[0].(List)
			if !ok {
				return nil, fmt.Errorf("lambda parameters must be a list")
			}
			for _, param := range params {
				if _, ok := param.(Symbol); !ok {
					return nil, fmt.Errorf("lambda parameters must be symbols")
				}
			}
			body := rest[1]
			return Lambda{params: params, body: body, env: NewEnv(env)}, nil
		}

		if sym, ok := first.(Symbol); ok && sym == "if" {
			if len(rest) < 2 || len(rest) > 3 {
				return nil, fmt.Errorf("if expects 2 or 3 arguments: condition, then, [else]")
			}
			condition, err := Eval(rest[0], env)
			if err != nil {
				return nil, err
			}
			fmt.Println("If condition result:", Print(condition))
			if condition != Nil && condition != False {
				return Eval(rest[1], env)
			} else if len(rest) == 3 {
				return Eval(rest[2], env)
			}
			return Nil, nil
		}
		proc, err := Eval(first, env)
		if err != nil {
			return nil, err
		}
		if lam, ok := proc.(Lambda); ok {
			if len(rest) != len(lam.params) {
				return nil, fmt.Errorf("lambda expects %d arguments, got %d", len(lam.params), len(rest))
			}
			newEnv := NewEnv(lam.env)
			for i, param := range lam.params {
				val, err := Eval(rest[i], env)
				if err != nil {
					return nil, err
				}
				fmt.Println("Binding", param, "to", Print(val))
				newEnv.SetVariable(param.(Symbol), val)
			}
			return Eval(lam.body, newEnv)
		}

		fun, ok := proc.(Func)
		if !ok {
			return nil, fmt.Errorf("first element must be a function, got %v", Print(proc))
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
