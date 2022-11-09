package weather

type Observable interface {
	registerObserver(string, Observer)
	removeObserver(string)
	notifyObservers()
}

type Observer interface {
	update(interface{})
}

type DisplayElement interface {
	display()
}

type ObserverArgs interface {
}

type WeatherData struct {
	observers   map[string]Observer
	temperature float64
	humidity    float64
	pressure    float64
	changed     bool
}

func NewWeatherData() *WeatherData {
	return &WeatherData{
		observers: make(map[string]Observer),
	}
}

func (this *WeatherData) registerObserver(name string, o Observer) {
	this.observers[name] = o
}

func (this *WeatherData) removeObserver(name string) {
	delete(this.observers, name)
}

func (this *WeatherData) notifyObservers() {
	if this.changed == false {
		return
	}
	defer func() {
		this.changed = false
	}()
	for _, oo := range this.observers {
		oo.update(this)
	}
}

func (this *WeatherData) setChanged() {
	this.changed = true
}

func (this *WeatherData) measurementsChanged() {
	this.setChanged()
	this.notifyObservers()
}

func (this *WeatherData) setMeasurements(temperature float64, humidity float64, pressure float64) {
	this.temperature = temperature
	this.humidity = humidity
	this.pressure = pressure
	this.measurementsChanged()
}

func (this *WeatherData) getTemperature() float64 {
	return this.temperature
}

func (this *WeatherData) getHumidity() float64 {
	return this.humidity
}

func (this *WeatherData) getPressure() float64 {
	return this.pressure
}
