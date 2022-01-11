package main

import "log"

func main() {
	//var  colors map[string]string

	//colors := make(map[int]string)
	//colors[10] = "#ff00ff"

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4fb450",
		"white": "#ffffff",
	}

	printColors(colors)
}

func printColors(c map[string]string) {
	for color, hex := range c {
		log.Println("Hex code for", color, "is", hex)
	}
}
