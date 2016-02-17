package main

import (
	"github.com/huichen/wukong/engine"
	"github.com/huichen/wukong/types"
	"log"
)

var (
// searcher是协程安全的
	searcher = engine.Engine{}
)

func main() {
	// 初始化
	dict := "../../../../../github.com/huichen/wukong/data/dictionary.txt"
	//dict := "dictionary.txt"
	searcher.Init(types.EngineInitOptions{SegmenterDictionaries: dict})
	defer searcher.Close()

	// 将文档加入索引
	searcher.IndexDocument(0, types.DocumentIndexData{Content: "此次百度收购将成中国互联网最大并购"})
	searcher.IndexDocument(1, types.DocumentIndexData{Content: "百度宣布拟全资收购91无线业务"})
	searcher.IndexDocument(2, types.DocumentIndexData{Content: "百度是中国最大的搜索引擎"})
	searcher.IndexDocument(3, types.DocumentIndexData{Content: "中国人事考试网：中国人事考试中心网站由人力资源和社会保障部人事考试中心主办，提供中国人事考试信息网上服务。中国人事考试中心主要承担公务员录用考试、专业技术人员资格考试、公务员遴选考试、中央单位接收安置军转干部考试和事业单位公开招聘考试等五大类，50余项考试的命题、阅卷、考务组织、考试技术指导和考试服务等工作。"})

	// 等待索引刷新完毕
	searcher.FlushIndex()

	// 搜索输出格式见types.SearchResponse结构体
	res := searcher.Search(types.SearchRequest{Text:"中国人"})
	log.Printf("num=%d ", res.NumDocs)
	for _, d := range res.Docs {
		log.Printf("docId=%d", d.DocId)
		log.Print("\tscore:", d.Scores)
		log.Print("\tTokenLocations:", d.TokenLocations)
		log.Print("\tTokenSnippetLocations:", d.TokenSnippetLocations)
	}
}