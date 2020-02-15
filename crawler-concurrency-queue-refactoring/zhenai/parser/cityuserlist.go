package parser

import (
	"function/crawler-concurrency-queue-refactoring/engine"
	"regexp"
)

var (
	cityUserListRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityMore       = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCityUserList(contents []byte) engine.ParseResult {
	matches := cityUserListRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}

	for _, m := range matches {
		name := string(m[2])
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url: string(m[1]),
				ParserFunc: func(bytes []byte) engine.ParseResult {
					return ParseProfile(bytes, name)
				},
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
