package client

import (
	"fmt"
	"function/crawler-concurrency-distributed/config"
	"function/crawler-concurrency-distributed/engine"
	"function/crawler-concurrency-distributed/rpcsupport"
	"function/crawler-concurrency-distributed/worker"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort))
	if err != nil {
		return nil, err
	}
	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult
		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DesrializeResult(sResult), nil
	}, nil
}
