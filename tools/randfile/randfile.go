package main 

import (
	"os"
	"io/ioutil"
	"fmt"
	"strconv"
	"math/rand"
	"time"
)

func usage() {
	fmt.Printf("Usage : %v <filename> <size>\n", os.Args[0])
}

func main() {
	if len(os.Args) != 3 {
		usage()
		os.Exit(-1)
	}
	
	filepath := os.Args[1]
	size, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("parameter <size> input error. [%v]\n", err.Error())
		usage()
		os.Exit(-2)
	}
	
	const std = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	
	rand.Seed(time.Now().UnixNano())
	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		buf[i] = std[rand.Intn(len(std))]
	}
	err = ioutil.WriteFile(filepath, buf, 0755)
	if err != nil {
		fmt.Printf("write data to file [%v] failed: %v\n", filepath, err.Error())
		os.Exit(-3)
	}
}

