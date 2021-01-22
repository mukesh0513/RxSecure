package utils

import (
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"reflect"
)

func Validate(input map[string]interface{}, assertRule map[string]reflect.Kind) error{

	for key, assertType := range assertRule {
		value, ok := input[key]
		if !ok {
			return errors.New(fmt.Sprintf("[%v] not present in map", key))
		}

		//If we donot want to assert for type then the function will skip asserting type for that key
		if assertType == reflect.Invalid {
			continue
		}

		if value == nil {
			return errors.New(cast.ToString(key) + cast.ToString(" is null"))
		}

		if reflect.TypeOf(value).Kind() != assertType {
			return errors.New(cast.ToString(key) + cast.ToString(" not of type ") + cast.ToString(assertType))
		}
	}

	return nil
}
