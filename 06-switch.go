package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("write", i, " as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

	fmt.Println(time.Now())
	fmt.Println(time.April.String())

	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	fmt.Println(time.Date(2019, 12, 2, 2, 2, 2, 8, beijing))

	fmt.Println(time.Now().Date())
	fmt.Println(time.Now().Hour())
	fmt.Println(time.Now().Weekday())
	fmt.Println(int(time.Now().Month()))
}
