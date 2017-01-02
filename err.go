package goerr

// Error implements our custom errors
type Err struct {
	code  string
	Props map[string]interface{}
}

// Error to implement the error interface
func (e *Err) Error() string {
	return e.code
}

// NewErr instanciates
func New(code string, props map[string]interface{}) *Err {
	return &Err{
		code:  code,
		Props: props,
	}
}
