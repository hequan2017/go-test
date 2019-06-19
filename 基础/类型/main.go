package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	// strings     反射 检查类型
	s := "hello world! go language"
	s_split := strings.Split(s, " ")
	s_join := strings.Join(s_split, "")
	fmt.Println(reflect.TypeOf(s_split))
	fmt.Println(s_join)

}
