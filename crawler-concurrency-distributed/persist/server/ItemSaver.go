package main

import (
	"flag"
	"fmt"
	"function/crawler-concurrency-distributed/config"
	"function/crawler-concurrency-distributed/persist"
	"function/crawler-concurrency-distributed/rpcsupport"
	"github.com/olivere/elastic"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		log.Printf("must specify a port")
		*port = config.ItemSaverPort
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
