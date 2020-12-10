package main

import "fmt"
import "time"
import "html/template"
import "net/http"

import "github.com/fjukstad/headwind/yr"

func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fc, err := yr.GetForecast(0, 0)
		if err != nil {
			fmt.Println("Could not get forecast", err)
		}
		// bearing = 142.525
		t := time.Now()
		fmtTime := t.Format("2006-01-02 15:04:05")
		fmt.Printf("%s %s %s %s\n", fmtTime, r.RemoteAddr, r.Method, r.URL)

		direction := fc.Properties.TimeSeries[0].Data.Instant.Details.WindFromDirection
		speed := fc.Properties.TimeSeries[0].Data.Instant.Details.WindSpeed

		//	bearing := 142.5
		// to work

		// super fancy headwind or not calculator 💩
		headwind := false
		if direction < 90 || direction > 200 {
			headwind = true
		}

		description := ""
		if speed < 4 {
			description = "lett motvind"
		} else if speed < 8 {
			description = "grei motvind"
		} else if speed < 15 {
			description = "sterk motvind"
		} else {
			description = "kjør bil"
		}

		data := WindInfo{
			Headwind:    headwind,
			Description: description,
			Speed:       speed,
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)
}

type WindInfo struct {
	Headwind    bool
	Description string
	Speed       float64
}
