package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.saveToFile("my_cards")
	d := newDeckFromFile("my_cards")
	d.print()
}
