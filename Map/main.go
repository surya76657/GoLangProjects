package main

import "fmt"

func main() {

	// var colors map[string]string

	// colors := make(map[string]string)
	// colors["white"] = "#FFFFFF"
	// delete(colors, "white")

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
		"white": "#FFFFFF",
	}
	printMap(colors)

	// fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}
