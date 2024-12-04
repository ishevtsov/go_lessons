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
	roomsOccupied := rand.Intn(totalRooms)
	roomsAvailable := totalRooms - roomsOccupied

	occupancyRate := occupancyRate(roomsOccupied, totalRooms)
	occupancyLevel := occupancyLevel(occupancyRate)

	fmt.Println("Hotel:", hotelName)
	fmt.Println("Number of rooms", totalRooms)
	fmt.Println("Rooms available", roomsAvailable)
	fmt.Println("                  Occupancy Level:", occupancyLevel)
	fmt.Printf("                  Occupancy Rate: %0.2f %%\n", occupancyRate)

	if roomsAvailable > 0 {
		fmt.Println("Rooms:")
		for i := 0; roomsAvailable > i; i++ {
			roomNumber := firstRoomNumber + i
			size := rand.Intn(6) + 1
			nights := rand.Intn(10) + 1
			printRoomDetails(roomNumber, size, nights)
		}
	} else {
		fmt.Println("No rooms available for tonight")
	}
}

func occupancyLevel(occupancyRate float32) string {
	switch {
	case occupancyRate > 70:
		return "High"
	case occupancyRate > 20:
		return "Medium"
	default:
		return "Low"
	}
}

func occupancyRate(roomsOccupied, totalRooms int) float32 {
	return (float32(roomsOccupied) / float32(totalRooms)) * 100
}

func printRoomDetails(roomNumber, size, nights int) {
	nightText := "nights"
	if nights == 1 {
		nightText = "night"
	}
	fmt.Println(roomNumber, ":", size, "people /", nights, nightText)
}
