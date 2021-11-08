package main

import (
	"fmt"
	"math/rand"
	"github.com/thatisuday/commando"
)

func main() {

	// configure commando
	commando.
		SetExecutableName("quoteCli").
		SetVersion("1.0.0").
		SetDescription("This tool prints you a quote")

	// configure the root command
	commando.
		Register("get-a-quote").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			quotes := [9]string{"“The purpose of our lives is to be happy.” — Dalai Lama",
				"“Life is what happens when you’re busy making other plans.” — John Lennon",
				"“Get busy living or get busy dying.” — Stephen King",
				"“You only live once, but if you do it right, once is enough.” — Mae West",
				"“Many of life’s failures are people who did not realize how close they were to success when they gave up.”– Thomas A. Edison",
				"“If you want to live a happy life, tie it to a goal, not to people or things.”– Albert Einstein",
				"“Never let the fear of striking out keep you from playing the game.”– Babe Ruth",
				"“Money and success don’t change people; they merely amplify what is already there.” — Will Smith",
				"“Your time is limited, so don’t waste it living someone else’s life. Don’t be trapped by dogma – which is living with the results of other people’s thinking.” – Steve Jobs"}
			fmt.Printf("%s",quotes[rand.Intn(9)])
		})
	// parse command-line arguments
	commando.Parse(nil)

}
