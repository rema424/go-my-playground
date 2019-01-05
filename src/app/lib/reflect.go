package lib

import (
	"fmt"
	"reflect"
	"strings"
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
	fmt.Println("DebugPrint: ", DebugPrint(data))
	fmt.Println("")
}

// DebugPrint ...
func DebugPrint(data interface{}) (res string) {
	// relectの基礎型を取得する
	rv := reflect.ValueOf(data)
	rvk := rv.Kind()

	// ポインタ型やインタフェース型の場合は中身の型を取得する
	switch rvk {
	case reflect.Ptr, reflect.Interface:
		rv = rv.Elem()
		rvk = rv.Kind()
	}

	// 型で場合分けをして出力文字列を生成する
	switch rvk {
	case reflect.Struct:
		res = fmt.Sprintf("T:%v, V:%v", rv.Type(), DebugPrintReflectStruct(rv))
	case reflect.Array, reflect.Slice:
		// 配列・スライスの中身（0番目の要素）の型がStructの場合は専用のロジックで文字列を生成する
		if rv.Index(0).Kind() == reflect.Struct {
			var elements []string
			for i := 0; i < rv.Len(); i++ {
				// i番目の要素のreflect.Valueを取得する
				rv := rv.Index(i)

				// スライスに要素を詰める
				elements = append(elements, fmt.Sprintf("%d: %s", i, DebugPrintReflectStruct(rv)))
			}
			// スライスの中身をセパレーターを用いて結合する
			str := strings.Join(elements, ", ")

			str = "[" + str + "]"
			rvt := rv.Type()
			res = fmt.Sprintf("T:%v, V:%v", rvt, str)
		} else {
			rvt := rv.Type()
			rvi := rv.Interface()
			res = fmt.Sprintf("T:%v, V:%v", rvt, rvi)
		}
	default:
		rvt := rv.Type()
		rvi := rv.Interface()
		res = fmt.Sprintf("T:%v, V:%v", rvt, rvi)
	}
	return
}

// DebugPrintReflectStruct ...
func DebugPrintReflectStruct(rv reflect.Value) string {
	var fields []string

	// reflect.Type型を取得する
	rvt := rv.Type()

	// structのフィールド数でループ処理をする
	for i := 0; i < rv.NumField(); i++ {
		fk := rvt.Field(i) // reflect.StructField型
		fv := rv.Field(i)  // reflect.Value型

		// structのフィールド名を取得する
		fkn := fk.Name

		// structが外部公開されている場合はフィールドの値を取得する
		var fvi interface{}
		if fk.PkgPath == "" {
			fvi = fv.Interface()
		}

		// スライスに要素を詰める
		fields = append(fields, fmt.Sprintf("%v: %v", fkn, fvi))
	}
	// スライスの中身をセパレーターを用いて結合する
	str := strings.Join(fields, ", ")
	return "{" + str + "}"
}
