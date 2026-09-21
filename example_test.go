package gocard_test

import (
	"fmt"

	"github.com/asphaltbuffet/gocard"
)

func ExampleCard_String() {
	c := gocard.Card{Rank: gocard.Queen, Suit: gocard.Clubs}

	fmt.Println(c)
	// Output: Q♣
}
