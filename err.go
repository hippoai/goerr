package goerr

// Error implements our custom errors
type Err struct {
	Code  string                 `json:"code"`
	Props map[string]interface{} `json:"props"`
}

// Error to implement the error interface
func (e *Err) Error() string {
	return e.Code
}

// New instanciates
func New(code string, props map[string]interface{}) *Err {
	return &Err{
		Code:  code,
		Props: props,
	}
}

// NewS instanciates an error without props
func NewS(code string) *Err {
	return New(code, map[string]interface{}{})
}
