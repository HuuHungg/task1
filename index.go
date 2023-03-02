package main

import (
	"fmt"
)

func main()  {
	total := 0
	result := 0
	fmt.Scan(&total)

	for i := 1; i < total; i++ {
		if total%i == 0  {
			result++
		}
	}

	fmt.Println(result)
}
	
