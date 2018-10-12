package params

import (
	"fmt"
	"strings"
)

type Validator func(string) error

func GetValidator(name string) Validator {
    if v, ok := validators[name]; ok {
        return v
    }
    return noValidator
}

func emailValidator(value string) error {
	if !strings.Contains(value, "@") {
		return fmt.Errorf("email has to contain @")
	}
	return nil
}

func noValidator(value string) error {
	return nil
}

var validators = make(map[string]Validator)

func init() {
    validators["email"] = emailValidator
}