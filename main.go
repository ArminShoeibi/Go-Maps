package main

import "fmt"

func main() {
	colors1 := map[string]string{
		"Red":    "#FF0000",
		"Blue":   "#0000FF",
		"Yellow": "#FFFF00",
	}

	fmt.Println(colors1)

	delete(colors1, "Blue")
	fmt.Println(colors1)

	var colors2 map[string]string
	fmt.Println(colors2) // Zero Value

	colors3 := make(map[string]string)
	colors3["Cyan"] = "#00FFFF"
	fmt.Println(colors3)

}
