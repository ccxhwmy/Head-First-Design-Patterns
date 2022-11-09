package weather

import "testing"

func TestWeather(t *testing.T) {
	w := NewWeatherData()
	NewCurrentConditionsDisplay(w)
	w.setMeasurements(80, 65, 30.4)
}
