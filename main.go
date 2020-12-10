package main

import "fmt"

import "github.com/fjukstad/headwind/yr"

func main() {
	fc, err := yr.GetForecast(0, 0)
	if err != nil {
		fmt.Println("Could not get forecast", err)
	}
	// bearing = 142.525

	direction := fc.Properties.TimeSeries[0].Data.Instant.Details.WindFromDirection
	speed := fc.Properties.TimeSeries[0].Data.Instant.Details.WindSpeed

	bearing := 142.5

	// to work
	headwind := false
	if direction < 90 && direction > 200 {
		headwind = true
	}

	tailwind := !headwind

	fmt.Println(direction - bearing)

	fmt.Println("Headwind", headwind)

	fmt.Println("tail", tailwind)

	speedDescription := ""
	if speed < 4 {
		speedDescription = "lett"
	} else if speed < 8 {
		speedDescription = "tung"
	} else {
		speedDescription = "kjør bil"
	}

	fmt.Println(speedDescription)
}
