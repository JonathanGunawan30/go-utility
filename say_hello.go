package go_utility

import "errors"

func SayHello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name cannot be empty")
	}
	return "Hello " + name, nil
}
