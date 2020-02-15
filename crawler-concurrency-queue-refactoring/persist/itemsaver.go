package persist

import (
	"context"
	"errors"
	"function/crawler-concurrency-queue-refactoring/engine"
	esV7 "github.com/olivere/elastic"
	"log"
)

func ItemSave(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	// must turn off sniff in docker
	client, err := esV7.NewClient(esV7.SetSniff(false))
	if err != nil {
		return nil, err
	}

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			err := save(client, index, item)
			if err != nil {
				log.Printf("Item saver: error saving item %v: %v", item, err)
				continue
			}
		}
	}()
	return out, nil
}

func save(client *esV7.Client, index string, item engine.Item) (err error) {

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().Index(index).Type(item.Type).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err = indexService.Do(context.Background())

	if err != nil {
		return err
	}
	return nil
}
