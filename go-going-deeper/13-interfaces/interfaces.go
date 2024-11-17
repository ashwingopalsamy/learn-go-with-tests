package main

import "fmt"

// Define a Game struct
type Game struct {
	Title string
	Price int
}

// Define a Book struct
type Book struct {
	Title string
	Price int
}

// Define a Printer interface with a Print method
type Printer interface {
	Print()
}

// Implement Print method for Game
func (g *Game) Print() {
	fmt.Printf("Game: %s, Price: $%d\n", g.Title, g.Price)
}

// Implement Print method for Book
func (b *Book) Print() {
	fmt.Printf("Book: %s, Price: $%d\n", b.Title, b.Price)
}

// Function to print details of any Printer
func PrintDetails(p Printer) {
	p.Print()
}

func main() {
	// Initialize instances of Game and Book
	minecraft := &Game{Title: "Minecraft", Price: 20}
	atomicHabits := &Book{Title: "Atomic Habits", Price: 5}

	// Use the PrintDetails function to print each item
	PrintDetails(minecraft)
	PrintDetails(atomicHabits)

	// You can also create a slice of Printer interfaces
	store := []Printer{minecraft, atomicHabits}

	// Iterate over the store and print each item
	fmt.Println("Items in the store:")
	for _, item := range store {
		item.Print() // Calls the Print method for each item
	}
}
