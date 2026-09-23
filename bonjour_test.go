package main

import "testing"

func TestBonjour(t *testing.T) {

	verifierMessageCorrect := func(t testing.TB, resultat, attendu string) {
		t.Helper()
		if resultat != attendu {
			t.Errorf("reçu %q, attendu %q", resultat, attendu)
		}
	}
	t.Run("Dire bonjour aux personnes", func(t *testing.T) {

		t.Run("en espagnol", func(t *testing.T) {
			resultat := Bonjour("Espagnol", "Chris")
			attendu := "Hola, Chris"

			verifierMessageCorrect(t, resultat, attendu)
		})

		t.Run("en français", func(t *testing.T) {
			resultat := Bonjour("Français", "Chris")
			attendu := "Bonjour, Chris"

			verifierMessageCorrect(t, resultat, attendu)
		})

		t.Run("par default en anglais", func(t *testing.T) {
			resultat := Bonjour("", "Chris")
			attendu := "Hello, Chris"

			verifierMessageCorrect(t, resultat, attendu)
		})
	})

	t.Run("'Dire bonjour tout le monde' quand une chaine vide est fournie", func(t *testing.T) {

		t.Run("en espagnol", func(t *testing.T) {
			resultat := Bonjour("Espagnol", "")
			attendu := "Hola, mundo"

			verifierMessageCorrect(t, resultat, attendu)
		})

		t.Run("en français", func(t *testing.T) {
			resultat := Bonjour("Français", "")
			attendu := "Bonjour, monde"

			verifierMessageCorrect(t, resultat, attendu)
		})

		t.Run("par default en anglais", func(t *testing.T) {
			resultat := Bonjour("", "")
			attendu := "Hello, world"

			verifierMessageCorrect(t, resultat, attendu)
		})
	})
}
