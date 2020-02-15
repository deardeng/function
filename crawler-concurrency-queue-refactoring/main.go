package main

import (
	"function/crawler-concurrency-queue-refactoring/engine"
	"function/crawler-concurrency-queue-refactoring/persist"
	"function/crawler-concurrency-queue-refactoring/scheduler"
	"function/crawler-concurrency-queue-refactoring/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSave("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemSaver:   itemChan,
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
