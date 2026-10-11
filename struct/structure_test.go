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
		nom   string
		forme Forme
		aAire float64
	}{
		{nom: "Rectangle", forme: Rectangle{Largeur: 12, Hauteur: 6}, aAire: 72.0},
		{nom: "Cercle", forme: Cercle{Rayon: 10}, aAire: 314.1592653589793},
		{nom: "Triangle", forme: Triangle{Base: 12, Hauteur: 6}, aAire: 36.0},
	}

	for _, tt := range testAires {
		//on va utiliser t.Run pour rendre les cas de tests bien plus explicite
		t.Run(tt.nom, func(t *testing.T) {
			resultat := tt.forme.Aire()
			if resultat != tt.aAire {
				t.Errorf("%#v reçu %g, attendu %g", tt.forme, resultat, tt.aAire)
			}
		})
	}
}
