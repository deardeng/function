package parser

import (
	"function/crawler-concurrency-queue/engine"
	"regexp"
)

var cityUserListRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

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

	return result
}
