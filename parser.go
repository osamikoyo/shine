package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Tokenize(code string) []string {
	tokens := []string{}
	var current strings.Builder
	inString := false
	escape := false

	for _, s := range code {
		if inString {
			if escape {
				current.WriteRune(s)
				escape = false
			} else if s == '\\' {
				escape = true
				current.WriteRune(s)
			} else if s == '"' {
				current.WriteRune(s)
				inString = false
				tokens = append(tokens, current.String())
				current.Reset()
			} else {
				current.WriteRune(s)
			}
		} else {
			if s == '(' || s == ')' {
				if current.Len() > 0 {
					tokens = append(tokens, current.String())
					current.Reset()
				}
				tokens = append(tokens, string(s))
			} else if unicode.IsSpace(s) {
				if current.Len() > 0 {
					tokens = append(tokens, current.String())
					current.Reset()
				}
			} else if s == '"' {
				inString = true
				current.WriteRune(s)
			} else {
				current.WriteRune(s)
			}
		}
	}

	if current.Len() > 0 {
		if inString {
			fmt.Println("Error: unclosed string literal")
			return nil
		}
		tokens = append(tokens, current.String())
	}

	outmanager.Println("high", "tokens: ", tokens)
	return tokens
}

func atom(token string) Value {
	outmanager.Println("high", token, " on atom")
	if num, err := strconv.ParseFloat(token, 64); err == nil {
		return Number(num)
	}
	if token == "true" {
		return Bool(true)
	}
	if token == "false" {
		return Bool(false)
	}
	if token == "nil" {
		return Nil
	}
	if token[0] == '"' && token[len(token)-1] == '"' && len(token) >= 2 {
		return String(token[1 : len(token)-1])
	}
	return Symbol(token)
}

func Read(tokens []string) (Value, []string, error) {
	if len(tokens) == 0 {
		return nil, tokens, errors.New("unexpected EOF")
	}
	token := tokens[0]
	tokens = tokens[1:]
	if token == "(" {
		list := List{}
		for len(tokens) > 0 && tokens[0] != ")" {
			val, newTokens, err := Read(tokens)
			if err != nil {
				return nil, tokens, err
			}
			list = append(list, val)
			tokens = newTokens
		}
		if len(tokens) == 0 {
			return nil, tokens, errors.New("missing )")
		}
		tokens = tokens[1:] // skip )
		return list, tokens, nil
	} else if token == ")" {
		return nil, tokens, errors.New("unexpected )")
	} else {
		return atom(token), tokens, nil
	}
}
