package main

import (
	"errors"
	"strings"
)

func Parse(message string)([]string, *string, error){
	if len(message) < 1 {
		return nil, nil, errors.New("message is empty")
	}

	if message[0:1] == "?" {
		split := strings.Split(message[1:], " ")
		lastWard := split[len(split) - 1]
		if lastWard[0:1] == "-" {
			last := lastWard[1:]
			return split, &last, nil
		}
		return split, nil, nil
	} else {
		return nil, nil, errors.New("no es un command")
	}
}