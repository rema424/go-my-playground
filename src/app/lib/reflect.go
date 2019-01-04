package lib

import (
	"fmt"
	"reflect"
)

// TypeAsertion ...
func TypeAsertion(data interface{}) {
	rt := reflect.TypeOf(data)  // reflect.Type型
	rv := reflect.ValueOf(data) // reflect.Value型

	rvt := rv.Type()      // reflect.Type型
	rvk := rv.Kind()      // 基礎型
	rvi := rv.Interface() // interface{}型

	fmt.Println("rt: ", rt)
	fmt.Println("rv: ", rv)
	fmt.Println("rvt: ", rvt)
	fmt.Println("rvk: ", rvk)
	fmt.Println("rvi: ", rvi)

	switch rvk {
	// ポインタの場合はポインタの指定先を取得して置き換える
	case reflect.Ptr:
		rv = rv.Elem()
		rvt = rv.Type()
		rvk = rv.Kind()
		rvi = rv.Interface()
	// inteface{}の場合は中身の型を取得して置き換える
	case reflect.Interface:
		rv = rv.Elem()
		rvt = rv.Type()
		rvk = rv.Kind()
		rvi = rv.Interface()
	}

	fmt.Println("Updated")
	fmt.Println("rv: ", rv)
	fmt.Println("rvt: ", rvt)
	fmt.Println("rvk: ", rvk)
	fmt.Println("rvi: ", rvi)

	switch rvk {
	case reflect.Map:
		fmt.Println("This is map.")
	case reflect.Slice:
		fmt.Println("This is slice.")
	case reflect.Array:
		fmt.Println("This is array.")
	case reflect.Struct:
		fmt.Println("This is struct.")
	default:
		fmt.Println("This is other type.")
	}
	fmt.Println("")
}
