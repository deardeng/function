package main

import (
	"fmt"
	"function/crawler-concurrency-distributed/config"
	"function/crawler-concurrency-distributed/rpcsupport"
	"function/crawler-concurrency-distributed/worker"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", config.WorkerPort), worker.CrawlService{}))
}
