package main

import (
	"flag"
	"fmt"

	"github.com/huichen/sego"
)

var (
	text = flag.String("text", "中国互联网历史上最大的一笔并购案", "要分词的文本")
)

func main() {
	flag.Parse()

	var seg sego.Segmenter
	dict := "../../../../../github.com/huichen/wukong/data/dictionary.txt"
	seg.LoadDictionary(dict)

	segments := seg.Segment([]byte(*text))
	fmt.Println(sego.SegmentsToString(segments, true))
	tokens := sego.SegmentsToTokens(segments, true)
	for i, t := range tokens {
		fmt.Printf("i=%v t=%v\n", i, t.Text())
	}
}
