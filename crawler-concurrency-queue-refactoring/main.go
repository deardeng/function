package main

import (
	"function/crawler-concurrency-queue-refactoring/engine"
	"function/crawler-concurrency-queue-refactoring/scheduler"
	"function/crawler-concurrency-queue-refactoring/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/beijing",
		ParserFunc: parser.ParseCityUserList,
	})
}
