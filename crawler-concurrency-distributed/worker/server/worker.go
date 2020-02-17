package main

import (
	"flag"
	"fmt"
	"function/crawler-concurrency-distributed/rpcsupport"
	"function/crawler-concurrency-distributed/worker"
	"log"
)

var port = flag.Int("port", 0, "the port for worker to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		log.Fatal("must specify a port")
	}
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))
}
