package params

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// Unpack fills the struct pointed by the ptr param
// with values from HTTP query parameters
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // reflect.StructField
		tag := fieldInfo.Tag           // reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
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
		return nil
	}
	return nil
}

// Pack packs a struct into url
func Pack(ptr interface{}) string {
	values := url.Values{}
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // reflect.StructField
		tag := fieldInfo.Tag           // reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		f := v.Field(i)
		if f.Kind() == reflect.Slice {
			for j := 0; j < f.Len(); j++ {
				setValue(&values, name, f.Index(j))
			}
		} else {
			setValue(&values, name, v.Field(i))
		}
	}
	return values.Encode()
}

func setValue(values *url.Values, param string, v reflect.Value) {
	switch v.Kind() {
	case reflect.String:
		values.Add(param, v.Interface().(string))
	case reflect.Int:
		values.Add(param, strconv.Itoa(v.Interface().(int)))
	case reflect.Bool:
		b := v.Interface().(bool)
		if b {
			values.Add(param, "true")
		} else {
			values.Add(param, "false")
		}
	}
}
