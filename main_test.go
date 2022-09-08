package main

import (
	"testing"
)

func TestGetQuote(t *testing.T) {

	sut := getQuote()
	if len(sut.Author) <= 0 {
		t.Error()
	}

}
