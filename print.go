package main

import (
	"fmt"
	"strings"
)

func Print(v Value) string {
	switch val := v.(type) {
	case Number:
		return fmt.Sprintf("%g", val)
	case String:
		return fmt.Sprintf("%q", string(val))
	case Symbol:
		return string(val)
	case Bool:
		if val {
			return "true"
		}
		return "false"
	case List:
		parts := []string{}
		for _, item := range val {
			parts = append(parts, Print(item))
		}
		return "(" + strings.Join(parts, " ") + ")"
	case Func:
		return "<function>"
	case Lambda:
		return "<lambda>"
	default:
		return "nil"
	}
}
