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
		resultat := Bonjour("Chris")
		attendu := "Bonjour, Chris"

		verifierMessageCorrect(t, resultat, attendu)
	})

	t.Run("'Dire bonjour tout le monde' quand une chaine vide est fournie", func(t *testing.T) {
		resultat := Bonjour("")
		attendu := "Bonjour, tout le monde"

		verifierMessageCorrect(t, resultat, attendu)
	})
}
