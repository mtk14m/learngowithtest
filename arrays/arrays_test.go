package arrays

import "testing"

func TestSomme(t *testing.T) {
	nombres := [5]int{1, 2, 3, 4, 5}

	resultat := Somme(nombres)
	attendu := 15

	if resultat != attendu {
		t.Errorf("reçu %d, attendu %d, donnée %v", resultat, attendu, nombres)
	}

}
