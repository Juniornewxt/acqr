package especificacao

import (
	"reflect"
)

func EstruturaTLV(data map[string]string, result interface{}) error {
	val := reflect.ValueOf(result).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		tag := typ.Field(i).Tag.Get("isoJUNIOR")
		if value, ok := data[tag]; ok {
			if field.CanSet() {
				field.SetString(value)
			}
		}
	}
	return nil
}
