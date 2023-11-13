// Package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package customerimporter

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

// DomainEntry ...
type DomainEntry struct {
	Count  int
	Domain string
}

// ProcessCustomerData ...
func ProcessCustomerData(filename string) ([]DomainEntry, error) {
	csvData, err := readCSVFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading CSV data: %v", err)
	}

	emailCounts, err := CountEmailDomains(csvData)
	if err != nil {
		return nil, fmt.Errorf("failed to count email domains: %v", err)
	}

	return sortDomainEntries(emailCounts), nil
}

func readCSVFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

// CountEmailDomains ...
func CountEmailDomains(records [][]string) (map[string]int, error) {
	if records == nil {
		return nil, errors.New("empty CSV data")
	}

	domainCounts := make(map[string]int)

	for _, record := range records {
		if len(record) < 3 {
            return nil, fmt.Errorf("malformed record: %v", record)
        }
		domain := extractDomain(record[2])
		domainCounts[domain]++
	}

	return domainCounts, nil
}

func sortDomainEntries(domainCounts map[string]int) []DomainEntry {
	entries := make([]DomainEntry, 0, len(domainCounts))

	for domain, count := range domainCounts {
		entries = append(entries, DomainEntry{Domain: domain, Count: count})
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Domain < entries[j].Domain
	})

	return entries
}

func extractDomain(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}
