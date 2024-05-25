package api

import (
	"log"
	db_ "telnesstech/db"
)

type Service struct {
	storage db_.Storage
	Handler Handler
}

func NewService() (Service, error) {
	db, err := db_.NewDbClient()
	if err != nil {
		log.Fatalf("could not create db client: %s", err)
	}

	storage, err := db_.NewStorage(db)
	if err != nil {
		log.Fatalf("could not create storage: %s", err)
	}

	handler, err := NewHandler(*storage)
	if err != nil {
		log.Fatalf("could not create handler: %s", err)
	}
	return Service{
		storage: *storage,
		Handler: *handler,
	}, nil
}
