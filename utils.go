package evolution

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

func StructToQueryString(data interface{}) (string, error) {
	queryParams := url.Values{}

	v := reflect.ValueOf(data)
	kind := v.Kind()
	if kind != reflect.Struct {
		return "", fmt.Errorf("input is not a struct")
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)

		tag := fieldType.Tag.Get("query")
		if tag == "" || tag == "-" {
			tag = fieldType.Name
		}

		switch field.Kind() {
		case reflect.String:
			if field.String() != "" {
				queryParams.Add(tag, field.String())
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if field.Int() != 0 {
				queryParams.Add(tag, strconv.FormatInt(field.Int(), 10))
			}
		case reflect.Bool:
			if field.Bool() {
				queryParams.Add(tag, strconv.FormatBool(field.Bool()))
			}
		case reflect.Slice:
			for j := 0; j < field.Len(); j++ {
				queryParams.Add(tag, fmt.Sprintf("%v", field.Index(j)))
			}
		}
	}

	return queryParams.Encode(), nil
}
