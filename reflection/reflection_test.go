package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "Struct with one string field",
			Input: struct {
				Name string
			}{"Chris"},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "Struct with two string fields",
			Input: struct {
				Name string
				City string
			}{"Dave", "London"},
			ExpectedCalls: []string{"Dave", "London"},
		},
		{
			Name: "Struct with two string fields",
			Input: struct {
				Name     string
				LuckyNum int
			}{"Dave", 6},
			ExpectedCalls: []string{"Dave"},
		},
		{
			Name: "Struct with nested fields",
			Input: Person{
				"Dave",
				Profile{38, "London"}},
			ExpectedCalls: []string{"Dave", "London"},
		},
		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("Got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
