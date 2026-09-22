package main

import "fmt"

const (
	Single = "single"
	Double = "double"
	Suite  = "suite"
)


const (
	Free        = "free"
	Booked      = "booked"
	Maintenance = "maintenance"
)


type HotelRoom struct {
	Type  string
	State string
	Price float64
}


func bookRoom(rooms map[string]HotelRoom, roomNumber string) {
	room, ok := rooms[roomNumber]
	if ok {
		room.State = Booked
		rooms[roomNumber] = room
	} else {
		fmt.Printf("Номер %s не найден\n", roomNumber)
	}
}

func main() {

	rooms := map[string]HotelRoom{
		"101": {Type: Single, State: Free, Price: 3000.0},
		"102": {Type: Double, State: Free, Price: 4500.0},
		"201": {Type: Suite, State: Maintenance, Price: 8000.0},
	}

	fmt.Printf("Статус номера 101 до бронирования: %+v\n", rooms["101"])


	bookRoom(rooms, "101")

	fmt.Printf("Статус номера 101 после бронирования: %+v\n", rooms["101"])
}