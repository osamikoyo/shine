package main

import (
	"errors"
	"fmt"
	"strings"
)

func SetStandartLibrary(env *Env) {
	env.SetFunc("get", func(args ...Value) (Value, error) {
		dict, ok := args[0].(Symbol)
		if !ok {
			return nil, fmt.Errorf("expects symbol, got: %v", args[0])
		}

		key, ok := args[1].(String)
		if !ok {
			return nil, fmt.Errorf("expects string, got: %v", args[1])
		}

		res, ok := env.GetVariable(dict)
		if !ok {
			return nil, fmt.Errorf("undefined")
		}

		resDict, ok := res.(Dictionary)
		if !ok {
			return nil, fmt.Errorf("%v is not dictionary", res)
		}

		return resDict[string(key)], nil
	})
	env.SetFunc("set", func(args ...Value) (Value, error) {
		dict, ok := args[0].(Symbol)
		if !ok {
			return nil, fmt.Errorf("excepts symbol, got: %v", args[0])
		}

		key, ok := args[1].(String)
		if !ok {
			return nil, fmt.Errorf("excepts string, got: %v", args[1])
		}

		res, ok := env.GetVariable(dict)
		if !ok {
			return nil, fmt.Errorf("unknown variable: %v", dict)
		}

		resMap, ok := res.(Dictionary)
		if !ok {
			return nil, fmt.Errorf("expects dictionary, got: %v", res)
		}

		resMap[string(key)] = args[2]

		env.SetVariable(dict, resMap)

		return args[2], nil
	})

	env.SetFunc("error", func(args ...Value) (Value, error) {
		msg, ok := args[0].(String)
		if !ok {
			return nil, fmt.Errorf("excepts strings, got: %v", args[0])
		}

		return nil, errors.New(string(msg))
	})
	env.SetFunc("$", func(args ...Value) (Value, error) {
		var list []Value

		list = append(list, args...)

		return List(list), nil
	})
	env.SetFunc("abs", func(args ...Value) (Value, error) {
		x, ok := args[0].(Number)
		if !ok {
			return nil, fmt.Errorf("expected number, got %v", Print(args[0]))
		}
		if x < 0 {
			return -x, nil
		}
		return x, nil
	})

	env.SetFunc("++", func(args ...Value) (Value, error) {
		arg := args[0]

		count, ok := arg.(Number)
		if !ok {
			return nil, fmt.Errorf("excepts number, got: %v", arg)
		}

		count++

		return count, nil
	})

	env.SetFunc("--", func(args ...Value) (Value, error) {
		arg := args[0]

		count, ok := arg.(Number)
		if !ok {
			return nil, fmt.Errorf("excepts number, got: %v", arg)
		}

		count--

		return count, nil
	})

	env.SetFunc("<=", func(args ...Value) (Value, error) {
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

		if f <= s {
			return True, nil
		}
		return False, nil
	})
	env.SetFunc(">=", func(args ...Value) (Value, error) {
		first := args[0]
		second := args[1]

		f, ok := first.(Number)
		if !ok {
			return nil, fmt.Errorf("excepts number, got: %v", first)
		}

		s, ok := second.(Number)
		if !ok {
			return nil, fmt.Errorf("excepts number, got: %v", second)
		}

		if f >= s {
			return True, nil
		} else {
			return False, nil
		}
	})

	env.SetFunc("||", func(args ...Value) (Value, error) {
		first := args[0]
		second := args[1]

		f, ok := first.(Bool)
		if !ok {
			return nil, fmt.Errorf("excepts bool: %v", first)
		}

		s, ok := second.(Bool)
		if !ok {
			return nil, fmt.Errorf("excepts bool: %v", second)
		}

		if s || f {
			return True, nil
		} else {
			return False, nil
		}
	})
	env.SetFunc("&&", func(args ...Value) (Value, error) {
		first := args[0]
		second := args[1]

		f, ok := first.(Bool)
		if !ok {
			return nil, fmt.Errorf("excepts bool: %v", first)
		}

		s, ok := second.(Bool)
		if !ok {
			return nil, fmt.Errorf("excepts bool: %v", second)
		}

		if s && f {
			return True, nil
		} else {
			return False, nil
		}
	})

	env.SetFunc("<", func(args ...Value) (Value, error) {
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

		if f < s {
			return True, nil
		} else {
			return False, nil
		}
	})
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
