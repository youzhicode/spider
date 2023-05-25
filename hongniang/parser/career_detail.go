package parser

import (
	"fmt"
	"regexp"
	"spider/engine"
)

const itemRe = `<a target="_blank" href="(/user/member/\S+)">(\S+)</a>`

func ParseCareer(content []byte) engine.ParseResult {

	re := regexp.MustCompile(itemRe)
	matchs := re.FindAllSubmatch(content, -1)

	parseResult := engine.ParseResult{}
	for _, m := range matchs {
		fmt.Printf("%s\n", m[1:])
		parseResult.Items = append(parseResult.Items, string(m[2]))
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url:        "http://www.hongniang.com/" + string(m[1]),
			ParserFunc: ParseProfile,
		})
	}
	return parseResult
}
