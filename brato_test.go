package brato_test

import (
	"os"
	"testing"

	"github.com/ejholmes/brato"
)

func TestRead(t *testing.T) {
	f, err := os.Open("brato.json")
	if err != nil {
		t.Fatal(err)
	}

	_, err = brato.Read(f)
	if err != nil {
		t.Fatal(err)
	}
}
