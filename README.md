# go-scryfall

[![Build Status](https://travis-ci.com/BlueMonday/go-scryfall.svg?branch=master)](https://travis-ci.com/BlueMonday/go-scryfall) [![PkgGoDev](https://pkg.go.dev/badge/github.com/BlueMonday/go-scryfall)](https://pkg.go.dev/github.com/BlueMonday/go-scryfall) [![Coverage Status](https://img.shields.io/coveralls/github/BlueMonday/go-scryfall/master.svg)](https://coveralls.io/github/BlueMonday/go-scryfall?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/BlueMonday/go-scryfall)](https://goreportcard.com/report/github.com/BlueMonday/go-scryfall)

![go-scryfall](./go-scryfall.png)

`go-scryfall` is a Golang client library for accessing the [Scryfall](https://scryfall.com/) API.

The Scryfall logo is copyrighted by Scryfall, LLC. `go-scryfall` is not created
by, affiliated with, or supported by Scryfall.

`go-scryfall` art was provided by [@jouste](https://twitter.com/jouste) the fierce drawbarian!

## Example

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
		Unique:        scryfall.UniqueModePrints,
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

## Rate Limiting

`go-scryfall` will rate limit requests to Scryfall's API. The default limit is
10 requests per second as recommended in the [REST API
documentation](https://scryfall.com/docs/api#rate-limits-and-good-citizenship).

To disable rate limiting use the `WithLimiter` option with a `nil` limiter when
constructing the client.

```golang
package main

import (
	"log"

	scryfall "github.com/BlueMonday/go-scryfall"
)

func main() {
	client, err := scryfall.NewClient(scryfall.WithLimiter(nil))
	if err != nil {
		log.Fatal(err)
	}
}
```
