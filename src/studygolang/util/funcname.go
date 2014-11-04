package util

import (
	"runtime"
	"fmt"
	"strconv"
)

func CallerName(skip int) (name string, file string, line int, ok bool) {
	var pc uintptr
	if pc, file, line, ok = runtime.Caller(skip+1); !ok {
		return    
	}    
	name = runtime.FuncForPC(pc).Name()    
	return	
}

func CallerFuncInfo() string {
	var name string
	var file string
	var line int
	var ok bool
	if name, file, line, ok = CallerName(1); ok {
		r := file + ":" + strconv.Itoa(line) + " " + name
		return r
	}
	
	return ""
}

func TestCallerName() {
	var name string
	var file string
	var line int
	var ok bool
	if name, file, line, ok = CallerName(0); ok {
		fmt.Printf("%v:%d %v", file, line, name)
	}
}
