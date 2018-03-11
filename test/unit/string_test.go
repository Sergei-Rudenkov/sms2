package unit

import (
	"testing"
	"sms2/service"
)
/////////////////////////////////////////////////
//
// test arguments parsing for String operations
//
/////////////////////////////////////////////////


func TestGet_HappyPath(t *testing.T)  {
	const key = "kget"
	m, err := service.STelnetArgumentParser("get", key)
	assertEqual(t, err, nil, "")
	assertEqual(t, m[`key`], key, "")
}

func TestGet_WrongArg(t *testing.T)  {
	_, err := service.STelnetArgumentParser("get")
	assertNotNil(t, err, "")
}

func TestSet_HappyPath(t *testing.T)  {
	const key = "kset"
	const value = "str"
	m, err := service.STelnetArgumentParser("set", key, value)
	assertEqual(t, err, nil, "")
	assertEqual(t, m[`key`], key, "")
	assertEqual(t, m[`value`], value, "")
}

func TestSet_WrongArg(t *testing.T)  {
	_, err := service.STelnetArgumentParser("set")
	assertNotNil(t, err, "")
}