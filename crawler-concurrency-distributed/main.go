package main

import (
	"flag"
	"function/crawler-concurrency-distributed/config"
	"function/crawler-concurrency-distributed/engine"
	itemServer "function/crawler-concurrency-distributed/persist/client"
	"function/crawler-concurrency-distributed/rpcsupport"
	"function/crawler-concurrency-distributed/scheduler"
	"function/crawler-concurrency-distributed/worker/client"
	"function/crawler-concurrency-distributed/zhenai/parser"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemSaverHost", "", "itemsaver host")
	workerHosts   = flag.String("workerHosts", "", "worker hosts(comma separated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemServer.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))
	processor := client.CreateProcessor(pool)

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

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("connected to %s", h)
		} else {
			log.Printf("error connecting to %s: %v", h, err)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
