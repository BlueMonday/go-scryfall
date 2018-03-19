# go-scryfall

[![Build Status](https://travis-ci.org/BlueMonday/go-scryfall.svg?branch=master)](https://travis-ci.org/BlueMonday/go-scryfall) [![GoDoc](https://godoc.org/github.com/BlueMonday/go-scryfall?status.svg)](https://godoc.org/github.com/BlueMonday/go-scryfall) [![Coverage Status](https://img.shields.io/coveralls/github/BlueMonday/go-scryfall/master.svg)](https://coveralls.io/github/BlueMonday/go-scryfall?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/BlueMonday/go-scryfall)](https://goreportcard.com/report/github.com/BlueMonday/go-scryfall)

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

	sco := scryfall.SearchCardsOptions{
		Unique:        scryfall.UniqueModePrint,
		Order:         scryfall.OrderSet,
		Dir:           scryfall.DirDesc,
		IncludeExtras: true,
	}
	result, err := client.SearchCards(ctx, "storm cro", sco)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s", result.Cards[0].Colors)
}
```
