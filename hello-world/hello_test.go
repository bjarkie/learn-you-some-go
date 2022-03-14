package main

import "testing"

func TestHello(t *testing.T) {
	gotNoLanguage := Hello("Bjarki", "")
	wantNoLanguage := "Hello, Bjarki"

	if gotNoLanguage != wantNoLanguage {
		t.Errorf("got %q want %q", gotNoLanguage, wantNoLanguage)
	}

	got := Hello("Bjarki", "es")
	want := "Hola, Bjarki"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	gotJa := Hello("Bjarki", "ja")
	wantJa := "こんにちは, Bjarki"

	if gotJa != wantJa {
		t.Errorf("got %q want %q", gotJa, wantJa)
	}

	gotRu := Hello("Bjarki", "ru")
	wantRu := "Привет, Bjarki"

	if gotRu != wantRu {
		t.Errorf("got %q want %q", gotRu, wantRu)
	}

}