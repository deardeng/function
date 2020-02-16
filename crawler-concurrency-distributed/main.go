package main

import (
	"fmt"
	"function/crawler-concurrency-distributed/config"
	"function/crawler-concurrency-distributed/engine"
	itemServer "function/crawler-concurrency-distributed/persist/client"
	"function/crawler-concurrency-distributed/scheduler"
	"function/crawler-concurrency-distributed/worker/client"
	"function/crawler-concurrency-distributed/zhenai/parser"
)

func main() {
	itemChan, err := itemServer.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	processor, err := client.CreateProcessor()
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemSaver:        itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
	//e.Run(engine.Request{
	//	Url:    "http://www.zhenai.com/zhenghun/beijing",
	//	Parser: engine.NewFuncParser(parser.ParseCityUserList, config.ParseCityUserList),
	//})
}
