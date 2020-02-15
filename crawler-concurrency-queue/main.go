package main

import (
	"function/crawler-concurrency-queue/engine"
	"function/crawler-concurrency-queue/scheduler"
	"function/crawler-concurrency-queue/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
