package main
import (
    "os"
    "strconv"
    "time"
    "flag"
    "log"
)

func main() {
    flag.Parse()
    fp, _ := os.OpenFile("test_truncate." + strconv.FormatInt(time.Now().UnixNano(), 10) + ".txt", os.O_CREATE | os.O_RDWR, 0755)
    fp.WriteString("1234567890\n")
    n, _ := fp.Seek(0, os.SEEK_CUR)
    log.Printf("seek return n=%v", n)

    fp.WriteString("abcdefghij\n")
    n, _ = fp.Seek(0, os.SEEK_CUR)
    log.Printf("seek return n=%v", n)

    fp.Close()
}