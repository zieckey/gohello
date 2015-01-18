package pkg

import (
	"fmt"
)

func TestMap() {
	type Kvmap map[string]string
	type SectionMap map[string]Kvmap

	m := make(Kvmap)
	m["a"] = "av"
	m["b"] = "bv"
	fmt.Printf("m: %v\n", m)

	// 查询是否存在某个key
	_, exists := m["a"]
	fmt.Printf("a in map exists : %v", exists)

	fmt.Printf("===> Assign m -> s\n")
	s := make(SectionMap)
	s[""] = m	// 是引用的拷贝，map本身的内容都同一份
	s[""]["s"] = "sv"
	m["c"] = "cv"
	fmt.Printf("m: %v\n", m)
	sd, _ := s[""]
	fmt.Printf("s[\"\"]: %v\n", sd)

	fmt.Printf("===> Assign another m -> s\n")
	m = make(Kvmap)
	s["m"] = m
	s["m"]["s"] = "nnnsv"
	m["c"] = "nnncv"
	fmt.Printf("m: %v\n", m)
	sd, _ = s["m"]
	fmt.Printf("s[\"m\"]: %v\n", sd)
	fmt.Printf("s: %v\n", s)
	
	v,ok := m["c"]
	fmt.Printf("the value of [c] in m is [%v] ok=%v", v,ok)
}
