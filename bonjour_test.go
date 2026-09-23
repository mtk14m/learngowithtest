package main

import "testing"

func TestBonjour(t *testing.T) {
	name := "Chris"
	resultat := Bonjour(name)
	attendu := "Bonjour " + name

	if resultat != attendu {
		t.Errorf("reçu %q, attendu %q", resultat, attendu)
	}
}
