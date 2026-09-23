package main

import "fmt"

func Bonjour(name string) string {
	return "Bonjour " + name
}

func main() {
	name := "Chris"
	fmt.Println(Bonjour(name))
}
