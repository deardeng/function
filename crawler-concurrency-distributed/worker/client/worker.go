package client

import (
	"function/crawler-concurrency-distributed/config"
	"function/crawler-concurrency-distributed/engine"
	"function/crawler-concurrency-distributed/worker"
	"net/rpc"
)

func CreateProcessor(client chan *rpc.Client) engine.Processor {
	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)

		var sResult worker.ParseResult

		c := <-client

		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}

		return worker.DesrializeResult(sResult), nil
	}
}
