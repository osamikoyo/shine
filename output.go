package main

import "fmt"

type OutputManager struct {
	regime string
}

func NewOutputManager(regime string) *OutputManager {
	return &OutputManager{
		regime: regime,
	}
}

func (o *OutputManager) Println(regime string, values ...any) {
	if regime == o.regime {
		fmt.Println(values...)
	}
}
