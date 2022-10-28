package controller

import (
	"math/rand"
	"net/http"
	"text/template"

	"github.com/rasyidev/dts-h8-assignment-iii/models"
)

func RenderHTML(w http.ResponseWriter, r *http.Request) {
	weather := models.WeatherStatus{
		Water: uint(rand.Intn(100)),
		Wind:  uint(rand.Intn(100)),
	}

	weather.Status = weather.Check()

	t := template.Must(template.ParseFiles("./templates/weather-status.html"))
	t.ExecuteTemplate(w, "weather-status.html", weather)
}
