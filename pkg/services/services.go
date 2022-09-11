package services

import (
	"fmt"
	"reflect"
	"strings"
)

var AvailableFields map[string][]string = map[string][]string{
	"create": {"name", "age", "email"},
	"update": {"name", "age", "email"},
}

func AcceptableFields(incoming map[string]interface{}, acceptability []string) error {
	for key := range incoming {
		if isAcceptible := isAcceptable(key, acceptability); !isAcceptible {
			return fmt.Errorf("%s is an not a valid field", key)
		}
	}
	return nil
}

func isAcceptable(key string, acceptability []string) bool {
	for _, value := range acceptability {
		if value == key {
			return true
		}
	}
	return false
}

func GetStructFieldByTag(tag string, s interface{}) (string, error) {
	rt := reflect.TypeOf(s)
	if rt.Kind() != reflect.Struct {
		return "", fmt.Errorf("bad type")
	}
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		v := strings.Split(f.Tag.Get("json"), ",")[0] // use split to ignore tag "options"
		if v == tag {
			return f.Name, nil
		}
	}
	return "", nil
}
