package util

import (
	"strings"
)

func ListOfObjectsToConcatString(list []interface{}) string{
	stringList := make([]string, len(list))
	for i := range list {
		stringList[i] = list[i].(string)
	}
	return strings.Join(stringList,",")
}

func StringToList(str string) []string  {
	//trim
	str = strings.TrimSpace(str)
	str = strings.TrimPrefix(str, "[")
	str = strings.TrimSuffix(str, "]")
	//split
	list := strings.Split(str, ",")
	return list
}