package collection

import (
	"reflect"
	"testing"
)

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

func TestSommeTout(t *testing.T) {

	resultat := SommeTout([]int{1, 2}, []int{0, 9})
	attendu := []int{3, 9}
	if !reflect.DeepEqual(resultat, attendu) {
		t.Errorf("reçu %v attendu %v", resultat, attendu)
	}
}

func TestSommeToutQueue(t *testing.T) {

	verifierSomme := func(t testing.TB, resultat, attendu []int) {
		t.Helper()
		if !reflect.DeepEqual(resultat, attendu) {
			t.Errorf("reçu %v attendu %v", resultat, attendu)
		}
	}

	t.Run("Faire les sommes de quelques slices", func(t *testing.T) {

		resultat := sommeToutQueue([]int{1, 2}, []int{0, 9})
		attendu := []int{2, 9}

		verifierSomme(t, resultat, attendu)
	})

	t.Run("Sommer en toute securité des slices vides", func(t *testing.T) {
		resultat := sommeToutQueue([]int{}, []int{0, 4, 5})
		attendu := []int{0, 9}

		verifierSomme(t, resultat, attendu)
	})
}
