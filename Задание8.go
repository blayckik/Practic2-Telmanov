package main

import "fmt"

type LogEntry struct {
	IP        string
	HTTPCode  int
	Timestamp string
}

func filter(logs []LogEntry) []LogEntry {
	var result []LogEntry

	for _, log := range logs {
		if log.HTTPCode >= 400 && log.HTTPCode <= 599 {
			result = append(result, log)
		}
	}

	return result
}

func main() {

	logs := []LogEntry{
		{IP: "192.168.1.10", HTTPCode: 200, Timestamp: "10:00:01"},
		{IP: "192.168.1.11", HTTPCode: 404, Timestamp: "10:00:05"},
		{IP: "192.168.1.12", HTTPCode: 500, Timestamp: "10:00:12"},
		{IP: "192.168.1.13", HTTPCode: 301, Timestamp: "10:00:15"},
	}


	errorsOnly := filter(logs)


	fmt.Println("Запросы с ошибками (4xx и 5xx):")
	for _, entry := range errorsOnly {
		fmt.Printf("IP: %s, Код: %d, Время: %s\n", entry.IP, entry.HTTPCode, entry.Timestamp)
	}
}