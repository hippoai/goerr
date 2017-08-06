package goerr

import (
	"strings"
	"testing"
)

// TestNew tests instanciates
func TestNew(t *testing.T) {

	err := New("ERROR_CODE", map[string]interface{}{
		"key": "abc",
		"prop": map[string]interface{}{
			"value1": true,
			"value2": 1.2,
		},
	})

	if err == nil {
		t.Fatal("Error is nil?")
	}

	var e error
	e = err

	if !strings.HasPrefix(e.Error(), "[ERROR_CODE]") {
		t.Fatal("Wrong error code")
	}

	eErr, ok := e.(*Err)
	if !ok {
		t.Fatal("Can't parse back")
	}

	if eErr.Props["key"].(string) != "abc" {
		t.Fatal("Wrong key")
	}

	err2 := NewS("ERR")
	err2_ := New("ERR", map[string]interface{}{})

	if err2.Code != err2_.Code {
		t.Fatal("Whaaat")
	}

}
