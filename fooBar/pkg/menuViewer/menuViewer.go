package menuViewer

import (
	"ElenaBaza/fooBar/pkg/readingCSV"
	"fmt"
	"time"
)

func MenuView() []readingCSV.Drink {
	coffeeMenu := readingCSV.DrinkConverter("../csv/coffeeMenu.csv")
	nonAlcoholicMenu := readingCSV.DrinkConverter("../csv/nonAlcoholicMenu.csv")
	alcoholicMenu := readingCSV.DrinkConverter("../csv/alcoholicMenu.csv")
	beerMenu := readingCSV.DrinkConverter("../csv/beerMenu.csv")
	fmt.Printf("Welcome to fooBar!\n")
	fmt.Println("here is the MENU")
	time.Sleep(time.Second)
	fmt.Println("\tCoffee Menu")
	for _, elem := range coffeeMenu {
		fmt.Printf("%-32s%.2f\n", elem.Name, elem.Price)
	}
	fmt.Println("\tNon Alcoholic Menu")
	for _, elem := range nonAlcoholicMenu {
		fmt.Printf("%-32s%.2f\n", elem.Name, elem.Price)
	}
	fmt.Println("\tAlcoholic Menu")
	for _, elem := range alcoholicMenu {
		fmt.Printf("%-32s%.2f\n", elem.Name, elem.Price)
	}
	fmt.Println("\tBeer Menu")
	for _, elem := range beerMenu {
		fmt.Printf("%-32s%.2f\n", elem.Name, elem.Price)
	}
	var ret []readingCSV.Drink
	ret = append(ret, coffeeMenu...)
	ret = append(ret, nonAlcoholicMenu...)
	ret = append(ret, alcoholicMenu...)
	ret = append(ret, beerMenu...)
	return ret
}
