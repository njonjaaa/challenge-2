package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i =", i)
	}

	for j := 0; j <= 10; j++ {
		if j == 5 {
			var sentence = "C A M A P B O"
			for index, value := range sentence {
				if index%2 == 0 {
					{
						fmt.Printf("character %U '%c' starts at byte position %d\n", value, value, index)
					}
				}
			}
		} else {
			fmt.Println("Nilai j=", j)
		}
	}
}
