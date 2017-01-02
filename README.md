# GoErr: Errors wrapper for Go

Standard error structure.

Errors are stored in the following structure
```go
type Err struct {
	code  string
	Props map[string]interface{}
}
```

The `code` tells you what the error is, and the `Props` are where additional information are stored (where it happened, structure property, etc.)


## Install

`go get -u github.com/hippoai/goerr.git`

## Manual


You can create a new error by using `New`

```go
func New(code string, props map[string]interface{}) *Err
```
