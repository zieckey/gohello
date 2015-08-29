package main
import (
    "os"
    "strconv"
    "time"
    "flag"
    "log"
)

func main() {
    var trunclen *int = flag.Int("t", -1, "The length of truncating")
    flag.Parse()
    fp, _ := os.OpenFile("test_truncate." + strconv.FormatInt(time.Now().UnixNano(), 10) + ".txt", os.O_CREATE | os.O_RDWR, 0755)
    fp.WriteString("1234567890\n")
    fp.WriteString("abcdefghij\n")

    if *trunclen >= 0 {
        fp.Truncate(int64(*trunclen))

        stat, _ := fp.Stat()
        log.Printf("%v len=%v", stat.Name(), stat.Size())
    }


    fp.Close()
}