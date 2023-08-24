package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (t TemperatureUnit) String() string {
	r := []string{"°C", "°F"}
	return r[t]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
// // => 21 °C
func (t Temperature) String() string {
	return fmt.Sprintf("%v %v", t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
// => mph km/h
func (s SpeedUnit) String() string {
	r := []string{
		"km/h", "mph",
	}
	return r[s]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
// Output: 18 km/h
func (s Speed) String() string {
	return fmt.Sprintf("%v %v", s.magnitude, s.unit)
}

// <location>: <temperature>, Wind <wind_direction> at <wind_speed>, <humidity>% Humidity
// => San Francisco: 57 °F, Wind NW at 19 mph, 60% Humidity
type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (m MeteorologyData) String() string {
	return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity", m.location, m.temperature, m.windDirection,
		m.windSpeed, m.humidity)
}

// Add a String method to MeteorologyData
