package main

import "fmt"

func useWeapon(ninja string, weapon string) string {
	return fmt.Sprintf(ninja + "is using" + weapon)
}

// multiple return value
func isValidLevel(level int) (int, bool) {
	if level > 10 {
		return level, true
	}
	return level, false
}

//variadic functions

func attack(evilNinjas ...string) {
	for _, evilNinja := range evilNinjas {
		fmt.Println("attacking evilNinjas in a loop ", evilNinja)
	}
}

func functionsDemo() {
	usage := useWeapon("Tommy", "Ninja Star")
	level, valid := isValidLevel(11)

	fmt.Println(usage, level, valid)
	attack("Tommy", "Johnny")
	attack("Tommy", "Johnny", "Andy", "Bobby")

	//if you already have multiple arguments in slice apply them using function func(slice...)
	evilNinjas := []string{"Tommy", "Johnny", "Andy"}
	attack(evilNinjas...)

	//closures
	attackToo := attack
	attackToo(evilNinjas...)
}
