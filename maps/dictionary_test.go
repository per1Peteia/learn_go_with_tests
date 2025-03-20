package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("good search - known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("bad search - unknown word", func(t *testing.T) {
		_, err := dict.Search("not there")
		if err == nil {
			t.Fatal("expected error but got nil")
		}
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dict := Dictionary{}

	t.Run("good add", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		err := dict.Add(key, value)

		assertError(t, err, nil)
		assertKey(t, dict, key, value)
	})

	t.Run("bad add - key already exists", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dict := Dictionary{key: value}
		err := dict.Add(key, "another test")
		assertError(t, err, ErrKeyExists)
		assertKey(t, dict, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("good update", func(t *testing.T) {
		dict := Dictionary{"test": "this is a test"}
		key := "test"
		value := "this is an updated test"
		err := dict.Update(key, value)

		assertError(t, err, nil)
		assertKey(t, dict, key, value)
	})

	t.Run("bad update - new word", func(t *testing.T) {
		dict := Dictionary{}
		key := "test"
		value := "this is a test"

		err := dict.Update(key, value)
		assertError(t, err, ErrKeyDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("good delete", func(t *testing.T) {
		dict := Dictionary{"test": "this is a test"}
		key := "test"
		err := dict.Delete(key)
		assertError(t, err, nil)

		_, err = dict.Search(key)
		assertError(t, err, ErrNotFound)
	})

	t.Run("bad delete - key does not exist", func(t *testing.T) {
		dict := Dictionary{}
		key := "test"
		err := dict.Delete(key)
		assertError(t, err, ErrKeyDoesNotExist)
	})

}

func assertKey(t testing.TB, dict Dictionary, key, value string) {
	t.Helper()
	got, err := dict.Search(key)
	if err != nil {
		t.Fatal("should find key, but did not")
	}
	assertStrings(t, got, value)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}

}
