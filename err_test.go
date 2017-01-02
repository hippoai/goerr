package goerr

import "testing"

// TestNew tests instanciates
func TestNew(t *testing.T) {

	err := NewErr("ERROR_CODE", map[string]interface{}{
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

	if e.Error() != "ERROR_CODE" {
		t.Fatal("Wrong error code")
	}

	eErr, ok := e.(*Err)
	if !ok {
		t.Fatal("Can't parse back")
	}

	if eErr.Props["key"].(string) != "abc" {
		t.Fatal("Wrong key")
	}

}
