package utils

import (
	"github.com/spf13/cast"
	"log"
	"reflect"
)

func Validate(input map[string]interface{}, assertRule map[string]reflect.Kind) {

	for key, assertType := range assertRule {
		value, ok := input[key]
		if !ok {
			log.Fatal(ok)
		}

		//If we donot want to assert for type then the function will skip asserting type for that key
		if assertType == reflect.Invalid {
			continue
		}

		if value == nil {
			log.Fatal(cast.ToString(key) + cast.ToString(" is null"))
		}

		if reflect.TypeOf(value).Kind() != assertType {
			log.Fatal(cast.ToString(key) + cast.ToString(" not of type ") + cast.ToString(assertType))
		}
	}
}