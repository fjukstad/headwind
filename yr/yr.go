package yr

import "errors"
import "fmt"
import "net/http"
import "encoding/json"
import "io/ioutil"

// Forecast is a weather forecast for a specific location
type Forecast struct {
	Properties Properties `json:"properties"`
}

// Properties
type Properties struct {
	Meta Meta `json:"meta"`
}

/// Meta
type Meta struct {
	UpdatedAt string `json:"updated_at"`
	Units     Units  `json:"units"`
}

type Units struct {
	WindFromDirection string `json:"wind_from_direction"`
	WindSpeed         string `json:"wind_speed"`
}

// GetForecast gets the forecast for a specific location
func GetForecast(lat, long float64) (Forecast, error) {
	baseURL := "https://api.met.no/weatherapi/locationforecast/2.0/compact?lat=67.2829&lon=14.4151"

	client := &http.Client{}
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return Forecast{}, err
	}

	req.Header.Set("User-Agent", "github.com/fjukstad/headwind")

	resp, err := client.Do(req)
	if err != nil {
		return Forecast{}, err
	}

	if resp.StatusCode != 200 {
		return Forecast{}, errors.New("Call failed. Response HTTP" + resp.Status)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return Forecast{}, err
	}

	var forecast Forecast
	err = json.Unmarshal(body, &forecast)
	if err != nil {
		return Forecast{}, err
	}

	return forecast, nil

}
