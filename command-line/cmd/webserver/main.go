package main

import (
	"log"
	"net/http"

	poker "github.com/bugg123/learn-go-with-tests/command-line"
)

const dbFileName = "game.db.json"

func main() {
	store, closeFunc, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("Unable to create file system player store from file %s %v", dbFileName, err)
	}

	defer closeFunc()

	server := poker.NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
