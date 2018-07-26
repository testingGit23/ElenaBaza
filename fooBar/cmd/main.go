package main

import (
	"ElenaBaza/fooBar/pkg/makeOrder"
	"ElenaBaza/fooBar/pkg/menuViewer"
	"fmt"
)

func main() {
	menu := menuViewer.MenuView()
	fmt.Println(menu)
	order := makeOrder.Ordering()
	fmt.Println(order)
}
