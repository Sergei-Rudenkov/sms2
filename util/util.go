package util

import (
	"strings"
)

// ListOfObjectsToConcatString - convert list []interface{} to concatenated
// string like: "1,2,3"
func ListOfObjectsToConcatString(list []interface{}) string{
	stringList := make([]string, len(list))
	for i := range list {
		stringList[i] = list[i].(string)
	}
	return strings.Join(stringList,",")
}

// StringToList - parse string of type "[1,2,3]" to list
func StringToList(str string) []string  {
	//trim
	str = strings.TrimSpace(str)
	str = strings.TrimPrefix(str, "[")
	str = strings.TrimSuffix(str, "]")
	//split
	list := strings.Split(str, ",")
	return list
}