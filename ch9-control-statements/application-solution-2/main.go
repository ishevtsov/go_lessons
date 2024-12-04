package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const hotelName = "Gopher Paris Inn"
	const totalRooms = 134
	const firstRoomNumber = 110

	rand.Seed(time.Now().UTC().UnixNano())
	roomsOccupied := rand.Intn(totalRooms) + 1
	roomsAvailable := totalRooms - roomsOccupied

	occupancyRate := (float64(roomsOccupied) / float64(totalRooms)) * 100
	occupancyLevel := "Low"
	switch {
	case occupancyRate > 70:
		occupancyLevel = "High"
	case occupancyRate > 20:
		occupancyLevel = "Medium"
	}

	fmt.Println("Hotel:", hotelName)
	fmt.Println("Number of rooms", totalRooms)
	fmt.Println("Rooms available", roomsAvailable)
	fmt.Println("					Occupancy Level:", occupancyLevel)
	fmt.Printf("					Occupancy Rate: %0.2f%%\n", occupancyRate)

	if roomsAvailable > 0 {
		fmt.Println("Rooms:")
		for i := 0; i < roomsAvailable; i++ {
			roomNumber := firstRoomNumber + i
			size := rand.Intn(6) + 1
			nights := rand.Intn(10) + 1
			fmt.Println(roomNumber, ":", size, "people /", nights, " nights ")
		}
	} else {
		fmt.Println("No rooms available for tonight")
	}

}
