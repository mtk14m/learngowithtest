package main

import "fmt"

// const anglais = "Anglais"
const francais = "Français"
const espagnol = "Espagnol"

const prefixeSalutAnglais = "Hello, "
const prefixeSalutFrancais = "Bonjour, "
const prefixeSalutEspagnol = "Hola, "

func Bonjour(langue, name string) string {
	if langue == francais {
		if name == "" {
			name = "monde"
		}
		return prefixeSalutFrancais + name
	}

	if langue == espagnol {
		if name == "" {
			name = "mundo"
		}
		return prefixeSalutEspagnol + name
	}

	if name == "" {
		name = "world"
	}
	return prefixeSalutAnglais + name
}

func main() {
	name := "Chris"
	fmt.Println(Bonjour(espagnol, name))
}
