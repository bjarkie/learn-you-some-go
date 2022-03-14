package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello in Russain", func(t *testing.T) {
		gotRu := Hello("Bjarki", "ru")
		wantRu := "Привет, Bjarki"

		assertCorrectMessage(t, gotRu, wantRu)
	})
	t.Run("say hello in English if no language is selected", func(t *testing.T) {
		gotNoLanguage := Hello("Bjarki", "")
		wantNoLanguage := "Hello, Bjarki"

		assertCorrectMessage(t, gotNoLanguage, wantNoLanguage)
	})
	t.Run("say hello in Japanese ", func(t *testing.T) {
		gotJa := Hello("Bjarki", "ja")
		wantJa := "こんにちは, Bjarki"

		assertCorrectMessage(t, gotJa, wantJa)
	})
}