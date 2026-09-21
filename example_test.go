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

func ExampleDeck_Draw() {
	d := gocard.NewDeck()

	c, err := d.Draw()
	if err != nil {
		panic(err)
	}

	fmt.Println(c, d.Len())
	// Output: A♣ 51
}
