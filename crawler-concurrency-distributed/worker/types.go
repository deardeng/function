package worker

import (
	"errors"
	"fmt"
	"function/crawler-concurrency-distributed/config"
	"function/crawler-concurrency-distributed/engine"
	"function/crawler-concurrency-distributed/zhenai/parser"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items   []engine.Item
	Request []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()

	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items:   r.Items,
		Request: nil,
	}
	for _, req := range r.Requests {
		result.Request = append(result.Request, SerializeRequest(req))
	}
	return result
}

func DesrializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}

	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func DesrializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Requests: nil,
		Items:    r.Items,
	}

	for _, req := range r.Request {
		engineReq, err := DesrializeRequest(req)
		if err != nil {
			log.Printf("error deserilizing request : %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCityUserList:
		return engine.NewFuncParser(parser.ParseCityUserList, config.ParseCityUserList), nil
	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParse(userName), nil
		} else {
			return nil, fmt.Errorf("invalid arg: %v", p.Args)
		}
	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("unkown parser name")
	}
}
