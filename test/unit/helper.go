package unit

import (
	"testing"
	"fmt"
	"reflect"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func assertDeepEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if reflect.DeepEqual(a, b) {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func assertNotNil(t *testing.T, a interface{}, message string){
	if a != nil {
		return
	}
	if len(message) == 0 {
		message = "argument is nil"
	}
	t.Fatal(message)
}
