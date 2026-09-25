package collection

import "testing"

func TestSomme(t *testing.T) {

	t.Run("Sommer une collection de n numbres", func(t *testing.T) {
		nombres := []int{1, 2, 3}

		resultat := Somme(nombres)
		attendu := 6

		if resultat != attendu {
			t.Errorf("reçu %d, attendu %d, donnée %v", resultat, attendu, nombres)
		}
	})

}
