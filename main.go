package main

import "fmt"

import "github.com/fjukstad/headwind/yr"

func main() {
	fc, err := yr.GetForecast(0, 0)
	if err != nil {
		fmt.Println("Could not get forecast", err)
	}

	fmt.Println(fc)
}
