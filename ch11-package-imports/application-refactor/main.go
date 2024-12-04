package main

import (
	"application2/email"
	"application2/invoice"
)

func main() {
	// first reservation
	customerName := "Doe"
	customerEmail := "john.doe@example.com"
	var nights uint = 12
	emailContents := email.Contents("M", customerName, nights)
	email.Send(emailContents, customerEmail)
	invoice.Create(customerName, nights, 145.32)
}
