package customerimporter

import (
	"reflect"
	"testing"
)

func TestReadCSVFile(t *testing.T) {
	testFileName := "../data/customers.csv" 

	records, err := readCSVFile(testFileName)
	if err != nil {
		t.Errorf("readCSVFile() error = %v", err)
		return
	}
	if len(records) == 0 {
		t.Errorf("Expected non-empty records, got empty")
	}
}

func TestCountEmailDomains(t *testing.T) {
	records := [][]string{
		{"Alice", "Smith", "alice@example.com"},
		{"Bob", "Jones", "bob@example.com"},
		{"Carol", "Johnson", "carol@example.org"},
		{"Dave", "Williams", "dave@example.com"},
	}

	expected := map[string]int{
		"example.com": 3,
		"example.org": 1,
	}

	result, err := CountEmailDomains(records)
	if err != nil {
		t.Errorf("CountEmailDomains() error = %v", err)
		return
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("CountEmailDomains() = %v, want %v", result, expected)
	}
}

func TestSortDomainEntries(t *testing.T) {
	domainCounts := map[string]int{
		"example.com": 3,
		"example.org": 1,
		"example.net": 2,
	}

	expected := []DomainEntry{
		{Domain: "example.com", Count: 3},
		{Domain: "example.net", Count: 2},
		{Domain: "example.org", Count: 1},
	}

	result := sortDomainEntries(domainCounts)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("sortDomainEntries() = %v, want %v", result, expected)
	}
}

func TestExtractDomain(t *testing.T) {
	email := "test@example.com"
	expected := "example.com"
	result := extractDomain(email)
	if result != expected {
		t.Errorf("extractDomain() = %v, want %v", result, expected)
	}
}
