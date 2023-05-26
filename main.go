package main

import (
	"spider/engine"
	"spider/hongniang/parser"
)

func main() {
	url := "http://www.hongniang.com/index/search?sort=0&wh=1&sex=" +
		"0&starage=0&province=&city=&marriage=0&edu=0&income=0&height=0&" +
		"pro=0&house=0&child=0&xz=0&sx=0&mz=0&hometownprovince=0&zhiye="
	engine.Run(engine.Request{Url: url, ParserFunc: parser.Career})

	//content, err := fetcher.Fetcher("http://www.hongniang.com/index/search/zhiye/bailing")
	//if err != nil {
	//	fmt.Printf("fetch error:%v", err)
	//	os.Exit(1)
	//}
	//parser.ParseCareer(content)
}
