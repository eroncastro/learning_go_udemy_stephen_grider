package main

import "fmt"

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Printf("%s: %s\n", k, v)
	}
}

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	printMap(colors)

	x := make(map[string]string)

	x["lala"] = "popo"
	x["haha"] = "hehe"

	fmt.Println(x)
	fmt.Println(x["lala"])

	delete(x, "lala")
	fmt.Println(x)

	printMap(x)
}
