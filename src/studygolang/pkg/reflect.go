package pkg

import (
	"fmt"
	"reflect"
)

func TestReflect() {

	type T struct {
		A int
		B string
	}
	t := T{12, "hello"}
	reflectedType := reflect.ValueOf(&t).Elem()
	types := reflectedType.Type()
	for i := 0; i < reflectedType.NumField(); i++ {
		f := reflectedType.Field(i)
		fmt.Printf("Reflecting index=%d : %s %s = %v\r\n", i, types.Field(i).Name, f.Type(), f.Interface())
	}
}
