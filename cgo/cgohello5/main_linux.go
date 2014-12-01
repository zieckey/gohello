package main

/*
ERROR

$ go build
# github.com/zieckey/gohello/cgo/cgohello5
main(.text): undefined: myprint
*/

func main() {
    s := "hello string in golang";
    Prints(s)
}
