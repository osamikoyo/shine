package main

type Env struct {
	funcs     map[Symbol]Func
	variables map[Symbol]Value
	parent    *Env
}

func NewEnv(parent *Env) *Env {
	return &Env{
		funcs:     make(map[Symbol]Func),
		variables: make(map[Symbol]Value),
		parent:    parent,
	}
}

func (e *Env) SetFunc(key Symbol, fun Func) {
	e.funcs[key] = fun
}

func (e *Env) SetVariable(key Symbol, value Value) {
	e.variables[key] = value
}

func (e *Env) GetFunc(key Symbol) (Func, bool) {
	if value, ok := e.funcs[key]; ok {
		return value, ok
	}
	if e.parent != nil {
		return e.parent.GetFunc(key)
	}
	return nil, false
}

func (e *Env) GetVariable(key Symbol) (Value, bool) {
	if value, ok := e.variables[key]; ok {
		return value, ok
	}
	if e.parent != nil {
		return e.parent.GetVariable(key)
	}
	return nil, false
}
