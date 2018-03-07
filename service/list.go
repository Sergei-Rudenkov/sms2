package service

import (
	"errors"
	"regexp"
)

func LTelnetArgumentParser(commandName string, args ...string) (map[string]string, error) {
	var err error
	argMap := make(map[string]string)
	switch commandName {
	case `lget`, `lremove`:
		if len(args) > 0 &&  args[0] != ``{
			argMap[`key`] = args[0]
		}else{
			err = errors.New("absence of key argument in `lget` or `lremove` operation")
		}
		if len(args) > 1 &&  args[1] != ``{
			sliceArgument := args[1] // should look like e.g. `[1:5]` or `[:5]` or `[1:]`
			expression := "^\\[(?P<first>\\d*?):(?P<last>\\d*?)\\]$"
			r := regexp.MustCompile(expression)
			startEndArgumentList := r.FindStringSubmatch(sliceArgument)
			if len(startEndArgumentList)  >= 2 {
				argMap[`first`] = startEndArgumentList[1]
				argMap[`last`] = startEndArgumentList[2]
			}
		}
	}
	return argMap, err
}
