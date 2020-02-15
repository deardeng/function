package parser

import (
	"function/crawler-concurrency-queue-refactoring/engine"
	"regexp"
)

var (
	cityUserListRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityMore       = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCityUserList(contents []byte, _ string) engine.ParseResult {
	matches := cityUserListRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}

	for _, m := range matches {
		url := string(m[1])
		name := string(m[2])
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url:        url,
				ParserFunc: ProfileParser(name),
			})
	}

	submatch := cityMore.FindAllSubmatch(contents, -1)
	for _, m := range submatch {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCityUserList,
		})
	}

	return result
}
