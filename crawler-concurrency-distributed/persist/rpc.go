package persist

import (
	"context"
	"errors"
	"function/crawler-concurrency-distributed/engine"
	esV7 "github.com/olivere/elastic"
	"log"
)

type ItemSaverService struct {
	Client *esV7.Client
	Index  string
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

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := save(s.Client, s.Index, item)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v : %v", item, err)
	}
	return err
}
