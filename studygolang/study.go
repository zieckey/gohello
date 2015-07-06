package main

import "github.com/zieckey/gohello/studygolang/pkg"
import "github.com/zieckey/gohello/studygolang/util"
import "log"

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	pkg.StudyForLoop()
	pkg.FizzBuzzConstIfElse()
	pkg.StringFunc()
	pkg.TestPanicRecover()
	pkg.TestInterface()
	pkg.TestAnimalInterface()
	util.TestCallerName()
	pkg.TestHttpGet()
	pkg.TestGoPointer()
	pkg.TestArray()
	pkg.TestGobBinary()
	pkg.TestCustomEncodeDecode()
	pkg.TestInterfaceSerialize()
	pkg.TestEmptyInterface2()
	pkg.TestFileReadAndWrite()
	pkg.TestWriteFile()
	pkg.TestReflect()
	pkg.TestJSON()
	pkg.TestExec_date()
	pkg.TestMap()
	pkg.TestVariableNumberArguments()
	pkg.TestMd5()
	pkg.TestGoSimpleJSONUnicode()
	pkg.TestJSONUnicode()
	pkg.TestEmptyInterface3()
	pkg.TestUnicodeRune()
	pkg.TestHMAC()
	pkg.TestMD5HMAC()
	pkg.TestDerive()
	pkg.TestBytes()
	pkg.TestDefer()
	pkg.TestStructCopy()
	pkg.StringRune()
	pkg.TestGoSimpleJSONArraySet()
	pkg.TestGoSimpleJSON()
	pkg.TestCRC32IEEE()
	pkg.BinaryEncoding()
	pkg.TestInterface1()
	pkg.Testappend()
	pkg.TestGoSimpleJSONWrongType()
	pkg.TestRegexp()
	pkg.TestGBK2UTF8()
	pkg.TestPrintf()
}

/*
package main

import (
  "fmt"
  "testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
  if a == b {
    return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestSimple(t *testing.T) {
  a := 42
  assertEqual(t, a, 42, "")
  assertEqual(t, a, 43, "This message is displayed in place of the default one")
}

*/
