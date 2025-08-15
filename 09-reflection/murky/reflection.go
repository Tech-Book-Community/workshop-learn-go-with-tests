package murky

import (
	"reflect"
)

func walk(x any, fn func(input string)) {

	val := getValue(x)

	// 使用 walkValue 函式，讓 switch 裡的遞迴呼叫更乾淨、更少重複。
	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn) // interfacte() 取出實體
	}

	switch val.Kind() {
	case reflect.Struct:
		// Numfield 只有struct 有
		// Interface() 的作用是：
		// 把 reflect.Value 再轉回原本的型別的 interface{}
		for i := range val.NumField() {
			// field := val.Field(i)
			// walk(field.Interface(), fn)
			walkValue(val.Field(i))
		}

	case reflect.String:
		fn(val.String())
	// slice 和 array 是不同的 Kind，要明確處理
	case reflect.Slice, reflect.Array:
		for i := range val.Len() {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for {
			//Recv() 是 chan 專用
			// 一次戳一個出來
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		{
			valFnResult := val.Call(nil) // nil 代表沒有使用任何argument
			for _, res := range valFnResult {
				walkValue(res)
			}
		}
	}

}

func getValue(x any) reflect.Value {
	val := reflect.ValueOf(x)

	// pointer 沒有欄位會沒有NumField
	// 要使用 Elem() 來解引用(Dereference)
	if val.Kind() == reflect.Pointer {
		val = val.Elem() // 變成實體 Struct
	}

	return val
}
