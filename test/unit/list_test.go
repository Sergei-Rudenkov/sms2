package unit

import (
	"testing"
	"sms2/service"
)

/////////////////////////////////////////////////
//
// test arguments parsing for List operations
//
/////////////////////////////////////////////////

func TestLget_HappyPath(t *testing.T)  {
	const key = "k"
	const selectArgument = "[4:7]"
	m, err := service.LTelnetArgumentParser("lget", key, selectArgument)
	assertEqual(t, err, nil, "")
	assertEqual(t, m[`key`], key, "")
	assertEqual(t, m[`first`], "4", "")
	assertEqual(t, m[`last`], "7", "")
}

func TestLget_WrongArg(t *testing.T)  {
	_, err := service.LTelnetArgumentParser("lget")
	assertNotNil(t, err, "")
}

func TestLadd_HappyPath(t *testing.T)  {
	const key = "kadd"
	const add = "5"
	m, err := service.LTelnetArgumentParser("ladd", key, add)
	assertEqual(t, err, nil, "")
	assertEqual(t, m[`key`], key, "")
	assertEqual(t, m[`item`], "5", "")
}

func TestLadd_WrongArg(t *testing.T)  {
	const key = "kadd"
	_, err := service.LTelnetArgumentParser("ladd", key)
	assertNotNil(t, err, "")
}