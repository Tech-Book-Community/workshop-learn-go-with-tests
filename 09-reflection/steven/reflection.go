package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	// 使用 walkValue 函式，讓 switch 裡的遞迴呼叫更乾淨、更少重複。
	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn) // interface() 取出實體
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}

	// 如果是 map，遍歷所有 key，對每個 value 遞迴呼叫 walk。
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	// 用反射的 `Recv()` 不斷從 channel 取值，直到 channel 關閉
	// channel 有順序，可以直接比對結果
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	// 如果是 function（reflect.Func）
	// 用 val.Call(nil) 呼叫它（nil代表無參數）
	// 取得所有回傳值（每個都是 reflect.Value）。
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
