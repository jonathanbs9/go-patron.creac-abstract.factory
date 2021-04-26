package main

import "fmt"

func main() {
	adidasFactory, _ := GetSportsFactory("adidas")
	//nikeFactory, _ := GetSportsFactory("nike")

	//nikeShoe := nikeFactory.MakeShoe()
	//nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	//printShoeDetails(nikeShoe)
	//printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)

	fmt.Printf("asd")
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
