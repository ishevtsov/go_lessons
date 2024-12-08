package main

import "fmt"

func main() {
	customers := [4]string{"John Doe", "Helmuth Verein", "Dany Beril", "Oliver Lump"}
	customersSlice := customers[0:1]
	fmt.Println(customersSlice)

	customersSlice2 := customers[2:4]
	fmt.Println(customersSlice2)
}
