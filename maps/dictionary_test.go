package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run(`known word`, func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		got, _ := dictionary.Search("test")

		assertStrings(t, got, "this is just a test")
	})

	t.Run(`unknown word`, func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		_, err := dictionary.Search("cheese")
		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "just a test")
	want := "just a test"
	got, err := dictionary.Search(`test`)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if want != got {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf(`got %q want %q`, got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("expected to get an error.")
	}

	if got != want {
		t.Errorf(`got %q want %q`, got, want)
	}
}
