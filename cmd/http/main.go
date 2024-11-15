package main

import (
	"log"
	"skeleton/internal/boot"

	_ "github.com/lib/pq"
)

func main() {
	if err := boot.HTTP(); err != nil {
		log.Println("[HTTP] failed to boot http server due to " + err.Error())
	}
}
