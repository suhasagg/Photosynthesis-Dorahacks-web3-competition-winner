package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	tls := int64(123456789) // Your integer value; replace with your actual value

	fmt.Printf("Liquid stake amount: %d \n", tls) // Logging the message "Liquid stake amount: " followed by the `tls` value
	text := strconv.FormatInt(tls, 10)            // Converting `tls` to string type and storing it in `text`
	fmt.Printf("Liquid stake amount: %s \n", text) // Logging the message "Liquid stake amount: " followed by the `text` value

	// Convert the text to a byte slice because WriteFile requires a byte slice
	data := []byte(text + "\n") // Converting `text` to a byte slice and adding a newline character

	// Open the file in append mode, create it if it does not exist
	file, err := os.OpenFile("/media/usbHDD1/liquidstakeparameters", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err.Error())
		return
	}
	defer file.Close()

	// Write the data to the file
	_, err = file.Write(data)
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err.Error()) // Logging the error message if an error occurred
	}
}

