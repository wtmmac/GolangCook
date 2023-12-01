package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"

	assertStrings(t, got, want)
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		content := "this is just a test"

		err := dictionary.Add(key, content)

		assertError(t, err, nil)
		assertContent(t, dictionary, key, content)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "test"
		content := "this is just a test"
		dictionary := Dictionary{key: content}

		err := dictionary.Add(key, "new test")

		assertError(t, err, ErrWordExists)
		assertContent(t, dictionary, key, content)
	})
}

func assertContent(t testing.TB, dictionary Dictionary, key, content string) {
	t.Helper()

	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if content != got {
		t.Errorf("got %q want %q", got, content)
	}
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		key := "test"
		content := "this is just a test"
		dictionary := Dictionary{key: content}
		newContent := "new content"

		err := dictionary.Update(key, newContent)

		assertError(t, err, nil)
		assertContent(t, dictionary, key, newContent)
	})
	t.Run("new word", func(t *testing.T) {
		key := "test"
		content := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(key, content)

		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestDictionarySearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrNotFound)
	})
}

func TestDelete(t *testing.T) {
	key := "test"
	content := "test content"
	d := Dictionary{key: content}
	d.Delete(key)
	_, err := d.Search(key)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", key)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
