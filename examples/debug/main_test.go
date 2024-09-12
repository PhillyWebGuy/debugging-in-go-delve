package main

import (
	"testing"
)

func TestFindPerson(t *testing.T) {
	people := []*person{
		{name: "Alicia", id: 1, age: 30},
		{name: "Alex", id: 2, age: 25},
		{name: "Erica", id: 3, age: 40},
	}

	tests := []struct {
		name     string
		expected *person
	}{
		{"Alicia", people[0]},
		{"Alex", people[1]},
		{"Erica", people[2]},
		{"Derek", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findPerson(people, tt.name)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestPopulate(t *testing.T) {
	people := make([]*person, 0, 6)
	people = populate(people)

	if len(people) != 6 {
		t.Errorf("expected 6 people, got %d", len(people))
	}

	names := map[string]bool{
		"Alicia": false,
		"Alex":   false,
		"Erica":  false,
		"Derek":  false,
		"Julian": false,
		"Bob":    false,
	}

	for _, p := range people {
		if _, exists := names[p.name]; !exists {
			t.Errorf("unexpected name %s", p.name)
		}
		names[p.name] = true
	}

	for name, found := range names {
		if !found {
			t.Errorf("expected name %s not found", name)
		}
	}
}
