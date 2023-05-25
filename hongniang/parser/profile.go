package parser

import (
	"regexp"
	"spider/engine"
	"spider/model"
)

const infoRe = `<li><span>(\S+)：</span>(\S+)</li> <li><span>(\S+)：</span>(\S+)</li>`

var (
	ageRe       = regexp.MustCompile(`<li><span>年龄：</span>(\S+)岁</li>`)
	marryRe     = regexp.MustCompile(`<li><span>婚况：</span>(\S+)</li>`)
	heightRe    = regexp.MustCompile(`<li><span>身高：</span>(\S+)</li>`)
	educationRe = regexp.MustCompile(`<li><span>学历：</span>(\S+)</li>`)
	inCome      = regexp.MustCompile(`<li><span>年收入：</span>(\S+)</li>`)
	workAddr    = regexp.MustCompile(`<li><span>工作地：</span>(\S+)</li>`)
	loveIdRe    = regexp.MustCompile(`<div class="loveid">Loveid:(\S+)</div>`)
	sexRe       = regexp.MustCompile(`<li><span>性别：</span>(\S+)</li>`)
	weightRe    = regexp.MustCompile(`<li><span>体重：</span>(\S+)</li>`)
)

func ParseProfile(content []byte) engine.ParseResult {
	var person model.Profile

	person.LoveId = findInfo(content, loveIdRe)
	person.Age = findInfo(content, ageRe)
	person.Sex = findInfo(content, sexRe)
	person.Height = findInfo(content, heightRe)
	person.Weight = findInfo(content, weightRe)
	person.IsMarry = findInfo(content, marryRe)
	person.Education = findInfo(content, educationRe)
	person.WorkAddr = findInfo(content, workAddr)
	person.Income = findInfo(content, inCome)

	parseResult := engine.ParseResult{}
	parseResult.Items = append(parseResult.Items, person)
	return parseResult
}

func findInfo(content []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(content)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
