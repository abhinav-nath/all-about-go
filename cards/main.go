package main

func main() {
	cards := deck{"Ace of Spades", "Ace of Diamonds"}
	cards = append(cards, "Queen of Hearts")
	cards.print()
}
