package weather

import "fmt"

type CurrentConditionsDisplay struct {
	temperature float64
	humidity    float64
	weatherData Observable
}

func NewCurrentConditionsDisplay(w Observable) *CurrentConditionsDisplay {
	current := &CurrentConditionsDisplay{
		weatherData: w,
	}
	w.registerObserver("CurrentConditionsDisplay", current)
	return current
}

func (this *CurrentConditionsDisplay) update(w interface{}) {
	weatherData := w.(*WeatherData)
	if weatherData == nil {
		return
	}
	this.temperature = weatherData.getTemperature()
	this.humidity = weatherData.getHumidity()
	this.display()
}

func (this *CurrentConditionsDisplay) display() {
	fmt.Println("Current conditions: ", this.temperature, "F degrees and", this.humidity, "% humidity")
}
