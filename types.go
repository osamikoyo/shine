package main

type (
	Symbol     string
	Value      interface{}
	Error      string
	List       []Value
	Number     float64
	String     string
	Bool       bool
	Dictionary map[string]Value
	Func       func(args ...Value) (Value, error)

	Lambda struct {
		params List
		body   Value
		env    *Env
	}
)

var (
	Nil   Value = nil
	True        = Bool(true)
	False       = Bool(false)
)
