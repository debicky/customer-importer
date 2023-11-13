// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package customerimporter

import (
	"encoding/csv"
	"errors"
	"os"
	"sort"
	"strings"
)

type DomainEntry struct {
	Count  int
	Domain string
}

func ProcessCustomerData(filename string) ([]DomainEntry, error) {
	csvData, err := readCSVFile(filename)
	if err != nil {
		return nil, err
	}

	emailCounts, err := CountEmailDomains(csvData)
	if err != nil {
		return nil, err
	}

	return sortDomainEntries(emailCounts), nil
}

func readCSVFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

func CountEmailDomains(records [][]string) (map[string]int, error) {
	if records == nil {
		return nil, errors.New("empty CSV data")
	}

	domainCounts := make(map[string]int)

	for _, record := range records {
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
