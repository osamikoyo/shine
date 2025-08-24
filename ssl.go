package main

import (
	"fmt"
	"strings"
)

func SetStandartLibrary(env *Env) {
	env.SetFunc(">", func(args ...Value) (Value, error) {
		first := args[0]
		second := args[1]

		f, ok := first.(Number)
		if !ok {
			return nil, fmt.Errorf("expected number, got: %v", first)
		}

		s, ok := second.(Number)
		if !ok {
			return nil, fmt.Errorf("expected number, got %v", second)
		}

		if f > s {
			return True, nil
		} else {
			return False, nil
		}
	})

	env.SetFunc("==", func(args ...Value) (Value, error) {
		first := args[0]
		second := args[1]

		switch arg := first.(type) {
		case String:
			s, ok := second.(String)
			if !ok {
				return nil, fmt.Errorf("expected string, got: %v", Print(second))
			}

			if s == arg {
				return True, nil
			}

			return False, nil
		case Number:
			s, ok := second.(Number)
			if !ok {
				return nil, fmt.Errorf("expected number, got: %v", Print(second))
			}

			if s == arg {
				return True, nil
			}

			return False, nil
		default:
			return nil, fmt.Errorf("unknown type: %v", Print(second))
		}
	})

	env.SetFunc("/", func(args ...Value) (Value, error) {
		count, ok := args[0].(Number)
		if !ok {
			return nil, fmt.Errorf("expected number, got %v", Print(count))
		}

		args = args[1:]

		for _, arg := range args {
			num, ok := arg.(Number)
			if !ok {
				return nil, fmt.Errorf("expected number, got %v", Print(num))
			}

			count /= num
		}

		return count, nil
	})

	env.SetFunc("+", func(args ...Value) (Value, error) {
		sum := Number(0)
		for _, arg := range args {
			num, ok := arg.(Number)
			if !ok {
				return nil, fmt.Errorf("expected number, got %v", Print(arg))
			}
			sum += num
		}
		return sum, nil
	})

	env.SetFunc("*", func(args ...Value) (Value, error) {
		exp := Number(1)
		for _, arg := range args {
			num, ok := arg.(Number)
			if !ok {
				return nil, fmt.Errorf("expected number, got %v", Print(arg))
			}

			exp *= num
		}

		return exp, nil
	})

	env.SetFunc("-", func(args ...Value) (Value, error) {
		count, ok := args[0].(Number)
		if !ok {
			return nil, fmt.Errorf("expected number, got %v", Print(count))
		}
		args = args[1:]
		for _, arg := range args {
			num, ok := arg.(Number)
			if !ok {
				return nil, fmt.Errorf("expected number, got %v", Print(arg))
			}
			count -= num
		}

		return count, nil
	})

	env.SetFunc("concat", func(args ...Value) (Value, error) {
		var result strings.Builder
		for _, arg := range args {
			str, ok := arg.(String)
			if !ok {
				return nil, fmt.Errorf("expected string, got %v", Print(arg))
			}
			result.WriteString(string(str))
		}
		return String(result.String()), nil
	})

	env.SetVariable("x", Number(5))
}
