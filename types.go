package main

type (
	Symbol string
	Value  interface{}
	List   []Value
	Number float64
	String string
	Bool   bool
	Func   func(args ...Value) (Value, error)

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
