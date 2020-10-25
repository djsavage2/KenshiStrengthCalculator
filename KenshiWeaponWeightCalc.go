package main

import (
	"fmt"
	"math"
)

func main() {
	var blunt float64
	var weight int
	var strength int

	fmt.Println("Kenshi Weapon Weight Calculator by DJSavage2")

	fmt.Println("Enter Weapon Blunt Damage:")

	_, err := fmt.Scanf("%f", &blunt)

	if err != nil {
		fmt.Println(err)
	}

	if blunt != 0 {
		strength = int(math.Ceil(blunt * 40))
	} else {
		fmt.Println("Enter Weapon Weight:")
		_, err := fmt.Scanf("%d", &weight)

		if err != nil {
			fmt.Println(err)
		}

		strength = weight * 2
	}

	fmt.Printf("Level %d strength is required to wield this weapon.\n", strength)

}
