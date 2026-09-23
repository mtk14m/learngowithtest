package main

import "fmt"

const prefixSalutFrançais = "Bonjour, "

func Bonjour(name string) string {
	if name != "" {
		return prefixSalutFrançais + name
	}
	return prefixSalutFrançais + "tout le monde"
}

func main() {
	name := "Chris"
	fmt.Println(Bonjour(name))
}
