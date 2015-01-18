package pkg

import (
	"hash/crc32"
	"fmt"
)

func TestCRC32IEEE() {
	d := "abc"
	crc := crc32.ChecksumIEEE([]byte(d))
	fmt.Printf("\n\ncrc32(%v)=%v\n", d, crc)
}