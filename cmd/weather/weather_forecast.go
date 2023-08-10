// Package waether api.
package weather

// 当前状态.
var CurrentCondition string

// 当前位置.
var CurrentLocation string

// 预测天气.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
