package dictionary

import "testing"

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
		assertNoError(t, err)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})

}
func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newdefinition := "new definition"

		err := dictionary.Update(word, newdefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newdefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExists)
	})

}

func TestDelete(t *testing.T) {
	t.Run("word existing", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}

		e := dictionary.Delete(word)

		_, err := dictionary.Search(word)

		assertNoError(t, e)

		if err != ErrNotFound {
			t.Errorf("Expected '%s' to be deleted", word)
		}
	})

	t.Run("no word", func(t *testing.T) {
		dictionary := Dictionary{}

		e := dictionary.Delete("test")

		assertError(t, e, ErrWordDoesNotExists)
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got '%s' want '%s'", got, definition)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an expected error", got)
	}
}
