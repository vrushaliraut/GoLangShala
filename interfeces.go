package main

import "fmt"

type ninjaWeapon interface {
	attack()
}

type ninjaStar struct{}

func (n ninjaStar) attack() {
	fmt.Println("Throwing ninjastar")
}

type ninjaSword struct{}

func (n ninjaSword) attack() {
	fmt.Println("Throwing ninjaSword")
}

func interfaces() {
	weapons := []ninjaWeapon{
		ninjaStar{},
		ninjaSword{},
	}

	for _, weapon := range weapons {
		weapon.attack()
	}
}
