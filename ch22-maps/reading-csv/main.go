package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type employee struct {
	name     string
	genre    string
	position string
}

func main() {
	employees := make(map[string]employee)
	file, err := os.Open("/home/shevtsov/code/Go/go_lessons/ch22-maps/reading-csv/data/employees.csv")
	if err != nil {
		log.Fatalf("can't open file %s", err)
	}

	defer file.Close()

	r := csv.NewReader(file)
	// skip csv header line
	_, err = r.Read()
	if err != nil {
		log.Fatal(err)
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		employeeId := record[0]
		employee := employee{
			name:     record[1],
			genre:    record[2],
			position: record[3],
		}
		employees[employeeId] = employee
	}

	fmt.Println(len(employees))

	employeeId := "V6555x"
	walter, ok := employees[employeeId]
	if ok {
		fmt.Println(walter)
	} else {
		fmt.Printf("No employee with ID: %s\n", employeeId)
	}

	for k, v := range employees {
		fmt.Printf("Key: %s - Value: %s\n", k, v)
	}
}
