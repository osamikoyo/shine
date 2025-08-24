package main

import (
	"fmt"
	"strings"
)

func Print(v Value) string {
	switch val := v.(type) {
	case String:
		return fmt.Sprintf("\"%s\"", val) // Печатаем строки с кавычками
	case List:
		strs := []string{}
		for _, item := range val {
			strs = append(strs, Print(item))
		}
		return "(" + strings.Join(strs, " ") + ")"
	case Symbol:
		return string(val)
	case Number:
		return fmt.Sprintf("%d", val)
	case Bool:
		if val {
			return "true"
		}
		return "false"
	case Func:
		return "<function>"
	default:
		return "nil"
	}
}
