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
		os.Exit(1)
	}

	customerEntries, err := customerimporter.ProcessCustomerData(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	displayCustomerEntries(customerEntries)
	os.Exit(0)
}

func displayCustomerEntries(entries []customerimporter.DomainEntry) {
	fmt.Println("Displaying customer entries:")
	for _, entry := range entries {
		fmt.Printf("  - %s: %d\n", entry.Domain, entry.Count)
	}
}
