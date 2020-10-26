package main

import (
	"fmt"
	"math"
)

func main() {
	var blunt float64
	var weight int

	fmt.Println("Kenshi Weapon Weight Calculator by DJSavage2")

	fmt.Println("\nEnter Weapon's Blunt Damage:")
	_, err := fmt.Scanf("%f\n", &blunt)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Enter Weapon's Weight:")
	_, err2 := fmt.Scanf("%d\n", &weight)
	if err2 != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Level %d strength is required to wield this weapon.\n", CalculateReqStrength(blunt, weight))
}

func CalculateReqStrength(bluntDmg float64, weightVal int) int {
	var strength int

	if bluntDmg != 0 {
		strength = int(math.Ceil(bluntDmg * 40))
	} else {
		strength = weightVal * 2
	}
	return strength
}
