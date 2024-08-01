package reflection

import "reflect"

func Walk(x any, fn func(input string)) {
	val := getValue(x)

	walkValue := func(v reflect.Value) {
		Walk(v.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Map:
		iter := val.MapRange()
		for iter.Next() {
			walkValue(iter.Value())
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Chan:
		for {
			v, ok := val.Recv()
			if !ok {
				break
			}
			walkValue(v)
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, v := range valFnResult {
			walkValue(v)
		}
	}
}

func getValue(x any) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		return val.Elem()
	}

	return val
}
