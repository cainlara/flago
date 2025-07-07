package flago

import (
	"os"
	"strings"
)

func GetArgs() (map[string]string, error) {
	args := os.Args

	if len(args) == 0 {
		return map[string]string{}, ErrEmptyInput
	}

	return getPurgedArgs(args)
}

func getPurgedArgs(args []string) (map[string]string, error) {
	if len(args) == 0 {
		return map[string]string{}, ErrEmptyInput
	}

	purgedArgs := args[1:]

	if len(purgedArgs) == 0 {
		return map[string]string{}, ErrEmptyInput
	}

	if len(purgedArgs)%2 != 0 {
		purgedArgs = append(purgedArgs, "")
	}

	mappedArgs := make(map[string]string, len(purgedArgs)/2)

	for index := 0; index < len(purgedArgs)-1; index += 2 {
		argName := strings.TrimPrefix(purgedArgs[index], Dash)

		mappedArgs[argName] = purgedArgs[index+1]
	}

	return mappedArgs, nil
}
