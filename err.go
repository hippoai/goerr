package goerr

import (
	"encoding/json"
	"fmt"
)

// pretty makes a thing pretty for print a JSON
func stringify(x interface{}) string {
	b, err := json.Marshal(x)
	if err != nil {
		return ""
	}
	return string(b)
}

// Error implements our custom errors
type Err struct {
	Code  string                 `json:"code"`
	Props map[string]interface{} `json:"props"`
}

// Error to implement the error interface
func (e *Err) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, stringify(e.Props))
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
