# go-scryfall

[![Build Status](https://travis-ci.org/BlueMonday/go-scryfall.svg?branch=master)](https://travis-ci.org/BlueMonday/go-scryfall) [![GoDoc](https://godoc.org/github.com/BlueMonday/go-scryfall?status.svg)](https://godoc.org/github.com/BlueMonday/go-scryfall) [![Coverage Status](https://coveralls.io/repos/github/BlueMonday/go-scryfall/badge.svg?branch=master)](https://coveralls.io/github/BlueMonday/go-scryfall?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/BlueMonday/go-scryfall)](https://goreportcard.com/report/github.com/BlueMonday/go-scryfall)

Golang client for the [Scryfall](https://scryfall.com/) API.


```golang
package main

import (
	"context"
	"log"

	scryfall "github.com/BlueMonday/go-scryfall"
)

func main() {
	ctx := context.Background()
	client, err := scryfall.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	so := scryfall.SearchCardsOptions{
		Unique:        scryfall.UniqueModePrint,
		Order:         scryfall.OrderSet,
		Dir:           scryfall.DirDesc,
		IncludeExtras: true,
	}
	result, err := client.SearchCards(ctx, "storm cro", so)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s", result.Cards[0].Colors)
}
```
