package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/bugg123/learn-go-with-tests/command-line"
)

const dbFileName = "game.db.json"

func main() {
	store, closeFunc, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("Unable to create file system player store from file %s %v", dbFileName, err)
	}

	defer closeFunc()

	fmt.Println("Let's play Poker")
	fmt.Println("Type {Name} wins to record a win")
	poker.NewCLI(store, os.Stdin).PlayPoker()
}
