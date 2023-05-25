package parser

import (
	"fmt"
	"spider/fetcher"
	"testing"
)

func TestCareer(t *testing.T) {
	contents, err := fetcher.Fetcher("http://www.hongniang.com/index/search?sort=0&wh=1&sex=" +
		"0&starage=0&province=&city=&marriage=0&edu=0&income=0&height=0&" +
		"pro=0&house=0&child=0&xz=0&sx=0&mz=0&hometownprovince=0&zhiye=")

	if err != nil {
		panic(err)
	}

	//Career(contents)
	fmt.Printf("%s\n", contents)
}
