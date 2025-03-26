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
		In            interface{}
		ExpectedCalls []string
	}{
		{
			"struct with 1 string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non-string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{
				{33, "London"},
				{34, "Paris"},
			},
			[]string{"London", "Paris"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "London"},
				{34, "Paris"},
			},
			[]string{"London", "Paris"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.In, func(in string) {
				got = append(got, in)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("maps", func(t *testing.T) {
		aMap := map[string]string{"cow": "moo", "sheep": "baa"}
		var got []string
		walk(aMap, func(in string) {
			got = append(got, in)
		})
		assertContains(t, got, "moo")
		assertContains(t, got, "baa")
	})

	t.Run("channels", func(t *testing.T) {
		ch := make(chan Profile)

		go func() {
			ch <- Profile{33, "London"}
			ch <- Profile{34, "Paris"}
			close(ch)
		}()

		var got []string
		want := []string{"London", "Paris"}

		walk(ch, func(in string) {
			got = append(got, in)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("with functions", func(t *testing.T) {
		f := func() (Profile, Profile) {
			return Profile{33, "London"}, Profile{34, "Paris"}
		}

		var got []string
		want := []string{"London", "Paris"}
		walk(f, func(in string) {
			got = append(got, in)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := true
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected to find %q in %v but did not", needle, haystack)
	}
}
