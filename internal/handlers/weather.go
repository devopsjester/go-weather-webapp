package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"go-weather-webapp/pkg/weather"
)

// TemplateDir is the base directory for template files.
// This variable can be set via dependency injection or configuration.
var TemplateDir = "templates"

// PageData holds the template data.
type PageData struct {
	Location *weather.Location
	Weather  *weather.Weather
	Error    string
	BgImage  string
}

// HTTPClient is the injected dependency for making HTTP requests.
// It can be overridden in tests.
var HTTPClient weather.HTTPGetter = http.DefaultClient

// selectBackgroundImage returns a background image URL based on the current time and weather condition.
func selectBackgroundImage(condition string) string {
	hour := time.Now().Hour()
	var timeOfDay string

	switch {
	case hour < 12:
		timeOfDay = "morning"
	case hour < 18:
		timeOfDay = "afternoon"
	default:
		timeOfDay = "night"
	}

	var url string
	switch condition {
	case "clear":
		switch timeOfDay {
		case "morning":
			url = "https://picsum.photos/id/1015/1280/720" // clear morning sky
		case "afternoon":
			url = "https://picsum.photos/id/1016/1280/720" // clear afternoon sky
		case "night":
			url = "https://picsum.photos/id/1018/1280/720" // clear night sky
		}
	case "cloudy":
		switch timeOfDay {
		case "morning":
			url = "https://picsum.photos/id/1020/1280/720" // cloudy morning sky
		case "afternoon":
			url = "https://picsum.photos/id/1021/1280/720" // cloudy afternoon sky
		case "night":
			url = "https://picsum.photos/id/1022/1280/720" // cloudy night sky
		}
	default:
		url = "https://picsum.photos/1280/720"
	}

	fmt.Printf("It's %s and it's %s outside.\n", time.Now().Format("3:04pm"), condition)
	fmt.Println("Background image URL:", url)
	return url
}

// WeatherHandler handles weather requests.
func WeatherHandler(w http.ResponseWriter, r *http.Request, zipcode string) {
	// Determine the current working directory.
	// cwd, err := os.Getwd()
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Error determining working directory: %v", err), http.StatusInternalServerError)
	// 	return
	// }
	// // Build the absolute template path.
	// tmplPath := filepath.Join(cwd, TemplateDir, "index.html")
	// fmt.Println("Working Directory:", cwd)
	tmplPath := filepath.Join(TemplateDir, "index.html")
	fmt.Println("Template path:", tmplPath)

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error loading template from %s: %v", tmplPath, err), http.StatusInternalServerError)
		return
	}

	data := PageData{}

	// Get the location details using the provided zipcode.
	loc, err := weather.GetLocationDetailsWithClient(HTTPClient, zipcode)
	if err != nil {
		data.Error = err.Error()
	} else {
		// Retrieve weather details from the Open-Meteo API.
		report, weatherData, err := weather.GetWeatherReport(HTTPClient, loc)
		if err != nil {
			data.Error = err.Error()
		} else {
			data.Location = &loc
			data.Weather = &weatherData
			data.Weather.Location = report
			// Set the background image URL based on the weather condition.
			data.BgImage = selectBackgroundImage(weatherData.Condition)
		}
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, fmt.Sprintf("Error rendering template: %v", err), http.StatusInternalServerError)
		return
	}
}
