package pkg

import "fmt"
import "strconv"

func TestUnicodeRune() {
    r := rune('啊')
    fmt.Println(RuneToAscii(r))
    fmt.Println(RuneToAscii('a'))
}

func RuneToAscii(r rune) string {
    if r < 128 {
        return string(r)
    } else {
        return "\\u" + strconv.FormatInt(int64(r), 16)
    }
}