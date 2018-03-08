package service

import (
	"errors"
	"regexp"
	"strconv"
	"sms2/storage"
)

func LTelnetArgumentParser(commandName string, args ...string) (map[string]string, error) {
	var err error
	cacheProvider := storage.GetCacheProviderType()
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
	case `ladd`:
		if len(args) > 0 &&  args[0] != ``{
			argMap[`key`] = args[0]
		} else {
			err = errors.New("absence of key argument in `ladd` operation")
		}
		if len(args) > 1 &&  args[1] != ``{
			argMap[`item`] = args[1]
		} else {
			err = errors.New("absence of item argument in `ladd` operation")
		}
		if cacheProvider == `agile` {
			if len(args) > 2  {
				if intTtl, intConversionErr := strconv.Atoi(args[2]); intConversionErr == nil && intTtl > 0 {
					argMap[`ttl`] = args[2]
					return argMap, err
				}
			}
			err = errors.New("please, provide int ttl greater then 0 as third argument")
		}
	}
	return argMap, err
}
