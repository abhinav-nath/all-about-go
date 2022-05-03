package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
	d := cards.newDeckFromFile("my_cards")
	d.print()
}
