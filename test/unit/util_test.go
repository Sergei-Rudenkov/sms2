package unit

import (
	"testing"
	"sms2/util"
)

func TestListOfObjectsToConcatString(t *testing.T){
	list := make([]interface{}, 0, 2)
	list = append(list, "t1")
	list = append(list, "t2")
	str := util.ListOfObjectsToConcatString(list)
	assertEqual(t, str, "t1,t2", "")
}

func TestStringToList(t *testing.T){
	list := util.StringToList("  [1,2,3] ")
	expected := []string{"1", "2", "3"}
	assertDeepEqual(t, list, expected, "")
}