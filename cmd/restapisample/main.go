package main

import (
	"github.com/datoga/restapisample"
)

func main() {
	repo := restapisample.NewInMemoryRepo()
	server := restapisample.NewServer(repo)

	if err := server.Start(); err != nil {
		panic(err)
	}
}
