package main

import "fmt"

import "html/template"
import "net/http"

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

	description := ""
	if speed < 4 {
		description = "lett"
	} else if speed < 8 {
		description = "tung"
	} else {
		description = "kjør bil"
	}

	fmt.Println(description)
	tmpl := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := WindInfo{
			Headwind:    headwind,
			Description: description,
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8080", nil)
}

type WindInfo struct {
	Headwind    bool
	Description string
}
