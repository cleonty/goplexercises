package params

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type valueValidator struct {
	value     reflect.Value
	validator Validator
}

// Unpack fills the struct pointed by the ptr param
// with values from HTTP query parameters
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	fields := make(map[string]valueValidator)
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // reflect.StructField
		tag := fieldInfo.Tag           // reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		validatorTag := tag.Get("validator")
		validator := GetValidator(validatorTag)
		fields[name] = valueValidator{v.Field(i), validator}
	}
	for name, values := range req.Form {
		f := fields[name].value
		validator := fields[name].validator
		if !f.IsValid() {
			continue
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, validator, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, validator, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, validator Validator, value string) error {
	if err := (validator(value)); err != nil {
		return err
	}
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)
	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)
	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)
	default:
		return fmt.Errorf("неподдерживаемый тип %s", v.Type())
	}
	return nil
}
