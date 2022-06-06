package main

import (
	"fmt"
	"strconv"
)

func main() {
	var arr [5]int
	v1 := 0
	v2 := 0
	v3 := 0
	v4 := 0
	v5 := 0
	v6 := 0
	z := 0
	fmt.Println("Давай сиграем в игру Greed. Введи 5 чисел от 1 до 6 (одно в каждую строку):")
	for i := 0; i < len(arr); i++ {
		fmt.Scanln(&arr[i])
		switch arr[i] {
		case 1:
			v1++
		case 2:
			v2++
		case 3:
			v3++
		case 4:
			v4++
		case 5:
			v5++
		case 6:
			v6++

		}
	}
	switch v1 {
	case 3:
		z += 1000
	case 2:
		z += 200
	case 1:
		z += 100
	case 4:
		z += 1100
	case 5:
		z += 1200
	}
	if v2 == 3 {
		z += 200
	}
	if v3 == 3 {
		z += 300
	}
	if v4 == 3 {
		z += 400
	}
	switch v5 {
	case 3:
		z += 500
	case 2:
		z += 100
	case 1:
		z += 50
	case 4:
		z += 550
	case 5:
		z += 600
	}
	if v6 == 3 {
		z += 600
	}
	fmt.Println("Твой счет: " + strconv.Itoa(z))
}
