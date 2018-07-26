package makeOrder

import (
	"ElenaBaza/fooBar/pkg/readingCSV"
	"fmt"
)

func Ordering(menu []readingCSV.Drink) map[string]int {
	fmt.Println("NOTE:ENTER to finish your order")
	fmt.Println("What will you order?")
	order := make(map[string]int)
	count := 0
	var drink string
	flag := false
	for {
		if count == 0 && flag {
			break
		}
		if flag {
			order[drink] = count
		}
		count = 0
		fmt.Scanf("%s %d", &drink, &count)
		for i, elem := range menu {
		}
		flag = true
	}
	return order
}
