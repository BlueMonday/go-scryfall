package scryfall_test

import (
	"context"
	"fmt"

	scryfall "github.com/BlueMonday/go-scryfall"
)

func ExampleClient_SearchCards() {
	ctx := context.Background()
	client, err := scryfall.NewClient()
	if err != nil {
		fmt.Println(err.Error())
	}

	so := scryfall.SearchCardsOptions{
		Unique:        scryfall.UniqueModePrints,
		Order:         scryfall.OrderSet,
		Dir:           scryfall.DirDesc,
		IncludeExtras: true,
	}
	result, err := client.SearchCards(ctx, "storm cro", so)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%s\n", result.Cards[0].Colors)
	// Output: [U]
}
