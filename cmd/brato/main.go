package main

import (
	"log"
	"os"

	"github.com/ejholmes/brato"
)

func main() {
	f, err := os.Open("brato.json")
	if err != nil {
		log.Fatal(err)
	}

	c, err := brato.Read(f)
	if err != nil {
		log.Fatal(err)
	}

	b := brato.New()
	if err := b.Sync(c); err != nil {
		log.Fatal(err)
	}
}
