package models

type WeatherStatus struct {
	Status string
	Water  uint
	Wind   uint
}

func (ws *WeatherStatus) Check() (status string) {
	if ws.Water < 5 && ws.Wind < 6 {
		status = "aman"
	} else if ws.Water >= 6 && ws.Water <= 8 && ws.Wind >= 7 && ws.Wind < 15 {
		status = "siaga"
	} else if ws.Water > 8 && ws.Wind > 15 {
		status = "bahaya"
	}
	return status
}
