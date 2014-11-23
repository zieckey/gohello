package main

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
)

func main() {
	hans := "中国人"
	// 默认输出 [[zhong] [guo] [ren]]
	fmt.Println(pinyin.Pinyin(hans, pinyin.Args{}))
	// 包含声调 [[zhōng] [guó] [rén]]
	fmt.Println(pinyin.Pinyin(hans, pinyin.Args{Style: pinyin.TONE}))
	// 声调用数字表示 [[zho1ng] [guo2] [re2n]]
	fmt.Println(pinyin.Pinyin(hans, pinyin.Args{Style: pinyin.TONE2}))
	// 开启多音字模式 [[zhong zhong] [guo] [ren]]
	fmt.Println(pinyin.Pinyin(hans, pinyin.Args{Heteronym: true}))
	// [[zho1ng zho4ng] [guo2] [re2n]]
	fmt.Println(pinyin.Pinyin(hans, pinyin.Args{Style: pinyin.TONE2, Heteronym: true}))
}
