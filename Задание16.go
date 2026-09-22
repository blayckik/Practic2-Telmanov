package main

import (
	"fmt"
	"time"
)

type SensorData struct {
	SensorID    string
	Temperature float64
	Humidity    float64
	Timestamp   time.Time
}

func calculate(data []SensorData) float64 {
	if len(data) == 0 {
		return 0.0
	}

	var totalTemp float64
	for _, entry := range data {
		totalTemp += entry.Temperature
	}

	return totalTemp / float64(len(data))
}

func main() {
	now := time.Now()

	records := []SensorData{
		{SensorID: "sensor-1", Temperature: 21.5, Humidity: 45.0, Timestamp: now.Add(-24 * time.Hour)},
		{SensorID: "sensor-1", Temperature: 19.0, Humidity: 50.0, Timestamp: now.Add(-20 * time.Hour)},
		{SensorID: "sensor-2", Temperature: 22.0, Humidity: 42.0, Timestamp: now.Add(-16 * time.Hour)},
		{SensorID: "sensor-1", Temperature: 24.5, Humidity: 40.0, Timestamp: now.Add(-12 * time.Hour)},
		{SensorID: "sensor-2", Temperature: 23.0, Humidity: 43.0, Timestamp: now.Add(-8 * time.Hour)},
		{SensorID: "sensor-1", Temperature: 20.0, Humidity: 48.0, Timestamp: now.Add(-4 * time.Hour)},
	}

	avgTemp := calculate(records)
	fmt.Println("Показания датчиков умного дома за сутки:")
	for _, record := range records {
		fmt.Printf("Датчик: %s | Температура: %.1f°C | Влажность: %.1f%% | Время: %s\n",
			record.SensorID, record.Temperature, record.Humidity, record.Timestamp.Format("15:04"))
	}

	fmt.Printf("\nСредняя температура за сутки: %.2f°C\n", avgTemp)
}