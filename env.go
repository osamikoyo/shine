package main

type Env struct {
	funcs     map[Symbol]Func
	variables map[Symbol]Value
}

func NewEnv() *Env {
	return &Env{
		funcs:     make(map[Symbol]Func),
		variables: make(map[Symbol]Value),
	}
}

func (e *Env) SetFunc(key Symbol, fun Func) {
	e.funcs[key] = fun
}

func (e *Env) SetVariable(key Symbol, value Value) {
	e.variables[key] = value
}

func (e *Env) GetFunc(key Symbol) (Func, bool) {
	value, ok := e.funcs[key]
	return value, ok
}

func (e *Env) GetVariable(key Symbol) (Value, bool) {
	value, ok := e.variables[key]
	return value, ok
}
