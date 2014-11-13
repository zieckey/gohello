package main

func f() int {
	return 0
}

func test_fallthrough() {
	i := 0
	
	//没有 fallthrough：
	switch i {
		case 0: // 空的 case 体
		case 1:
		f() // 当 i == 0 时，f 不会被调用！
	}
	
	//而这样：
	switch i {
		case 0: fallthrough
		case 1:
		f() // 当 i == 0 时，f 会被调用！
	}
}

func main() {

}




