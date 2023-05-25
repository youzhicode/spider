package parser

import (
	"regexp"
	"spider/engine"
)

const careerRe = `<a href="(/index/search/zhiye/\S+)">(\S+)</a>`

func Career(content []byte) engine.ParseResult {
	re := regexp.MustCompile(careerRe)
	matchs := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}

	for _, m := range matchs {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        "http://www.hongniang.com" + string(m[1]),
			ParserFunc: ParseCareer,
		})
	}
	return result
}
