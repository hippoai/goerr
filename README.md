# GoErr: Errors wrapper for Go

Standard error structure.

Errors are stored in the following structure
```go
type Err struct {
	Code  string
	Props map[string]interface{}
}
```

The `code` tells you what the error is, and the `Props` are where additional information are stored (where it happened, structure property, etc.)


## Install

`go get -u github.com/hippoai/goerr.git`

## Manual


You can create a new error by using `New`

```go
code := "MY_ERROR_CODE"
props := map[string]interface{}{
  "info about": "my error",
}
err := New(code, props)
```

and if you have no interesting property to return

```go
err := NewS(code)
```
