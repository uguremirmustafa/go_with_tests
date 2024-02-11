package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", english)
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", english)
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Turkish", func(t *testing.T) {
		got := Hello("Ugur", turkish)
		want := "Merhaba, Ugur"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Ugur", french)
		want := "Bonjour, Ugur"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	/* t.Helper() is needed to tell the test suite that this method is a helper.
	By doing this when it fails the line number reported will be in our function call
	rather than inside our test helper. */
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
