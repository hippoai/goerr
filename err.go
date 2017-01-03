package goerr

// Error implements our custom errors
type Err struct {
	Code  string
	Props map[string]interface{}
}

// Error to implement the error interface
func (e *Err) Error() string {
	return e.Code
}

// NewErr instanciates
func New(code string, props map[string]interface{}) *Err {
	return &Err{
		Code:  code,
		Props: props,
	}
}
