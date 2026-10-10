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
	testAires := []struct {
		forme   Forme
		attendu float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Cercle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range testAires {
		resultat := tt.forme.Aire()
		if resultat != tt.attendu {
			t.Errorf("reçu %g, attendu %g", resultat, tt.attendu)
		}
	}
}
