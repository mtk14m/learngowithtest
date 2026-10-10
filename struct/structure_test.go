package structure

import "testing"

func TestPerimetre(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	resultat := Perimetre(rectangle)
	attendu := 40.0

	if resultat != attendu {
		t.Errorf("reçu %g, attendu %g", resultat, attendu)
	}
}

func TestAire(t *testing.T) {

	verifierLAire := func(t testing.TB, forme Forme, attendu float64) {
		t.Helper()
		resultat := forme.Aire()

		if resultat != attendu {
			t.Errorf("reçu %g, attendu %g", resultat, attendu)
		}
	}
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		//resultat := rectangle.Aire()
		attendu := 72.0

		verifierLAire(t, rectangle, attendu)
	})

	t.Run("cercle", func(t *testing.T) {
		cercle := Cercle{10}
		//resultat := cercle.Aire()
		attendu := 314.1592653589793

		verifierLAire(t, cercle, attendu)
	})
}
