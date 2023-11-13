package main

import (
	"fmt"
	"log"
	"os"

	"github.com/debicky/customer-importer/customerimporter"
)

func main() {
	filename := os.Getenv("FILENAME")

	if filename == "" {
		log.Fatal("no input file provided")
	}

	customerEntries, err := customerimporter.ProcessCustomerData(filename)
	if err != nil {
		log.Fatalf("Error processing customer data: %v", err)
	}

	displayCustomerEntries(customerEntries)
}

func displayCustomerEntries(entries []customerimporter.DomainEntry) {
	fmt.Println("Displaying customer entries:")
	for _, entry := range entries {
		fmt.Printf("  - %s: %d\n", entry.Domain, entry.Count)
	}
}
