package worker

import (
	"function/crawler-concurrency-distributed/engine"
	"log"
)

type CrawlService struct {
}

func (CrawlService) Process(req Request, result *ParseResult) error {
	log.Printf("worker get task %v : %v", req.Url, req.Parser)
	engineReq, err := DesrializeRequest(req)
	if err != nil {
		return err
	}

	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}

	*result = SerializeResult(engineResult)
	return nil
}
