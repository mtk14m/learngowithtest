package main

import (
	"fmt"

	bonjour "github.com/mtk14m/ttdgo/bonjour"
)

func main() {
	name := "Chris"
	fmt.Println(bonjour.Bonjour("", name))
}
