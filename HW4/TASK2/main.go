package main

import (
	"fmt"
)

// Завдання про студентські оцінки. Створити slice float з оцінками студентів з певного предмету.
// Написати функцію, яка приймає на вхід slice оцінок та повертає середню оцінку з цього предмету.
// Використовуючи цю функцію, обчислити середній бал з предмету та вивести його на екран.

func CalculateAverage(inputFloatSlice []float64) float64 {
	sumOfMarks := 0.0
	substractFromLen := 0
	for _, mark := range inputFloatSlice {
		if mark >= 0 && mark <= 12 {
			sumOfMarks += mark
		} else {
			fmt.Println("---")
			fmt.Println("Mark is invalid --> ", mark)
			fmt.Println("Therefore this mark will not be used in calculations.")
			fmt.Println("---")
			substractFromLen++
		}
	}

	return sumOfMarks / float64(len(inputFloatSlice)-substractFromLen)

}

func main() {
	sliceFloatMath := []float64{9.34, 10.12, 11, 2, 6.1, 6.1, 4.1, -3.1}

	fmt.Println("*****************")
	fmt.Println("The average mark for Math Subject:")
	fmt.Println("*****************")
	fmt.Print(CalculateAverage(sliceFloatMath))

}
