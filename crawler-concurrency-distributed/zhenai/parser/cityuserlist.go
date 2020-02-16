package parser

import (
	"function/crawler-concurrency-distributed/config"
	"function/crawler-concurrency-distributed/engine"
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
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url:    string(m[1]),
				Parser: NewProfileParse(string(m[2])),
			})
	}

	submatch := cityMore.FindAllSubmatch(contents, -1)
	for _, m := range submatch {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCityUserList, config.ParseCityUserList),
		})
	}

	return result
}
