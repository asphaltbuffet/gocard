package gocard_test

import (
	"fmt"
	"math/rand/v2"

	"github.com/asphaltbuffet/gocard"
)

func ExampleCard_String() {
	c := gocard.Card{Rank: gocard.Queen, Suit: gocard.Clubs}

	fmt.Println(c)
	// Output: Q♣
}

func ExampleCard_Layout() {
	l := gocard.Card{Rank: gocard.Queen, Suit: gocard.Hearts}.Layout()

	fmt.Println(l.Width, l.Height, l.Border, l.Accent)
	// Output: 7 5 BorderRounded AccentRed
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

func ExampleDeck_Deal() {
	d := gocard.NewDeck()

	ways, err := d.Deal(4, 5)
	if err != nil {
		panic(err)
	}

	// A way is just a slice of cards — gocard does not know who holds one.
	fmt.Println(len(ways), len(ways[0]), d.Len())
	// Output: 4 5 32
}

func ExampleDeck_Shuffle() {
	// A seeded source makes the deal reproducible.
	d := gocard.NewDeck()
	d.Shuffle(rand.New(rand.NewPCG(1, 1)))

	c, err := d.Draw()
	if err != nil {
		panic(err)
	}

	fmt.Println(c.Valid(), d.Len())
	// Output: true 51
}
