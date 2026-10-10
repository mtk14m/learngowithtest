package structure

import "math"

type Rectangle struct {
	Largeur float64
	Hauteur float64
}

type Cercle struct {
	rayon float64
}

type Forme interface {
	Aire() float64
}

func Perimetre(r Rectangle) float64 {
	return 2 * (r.Hauteur + r.Largeur)
}

func (c Cercle) Aire() float64 {
	return math.Pi * c.rayon * c.rayon
}

func (r Rectangle) Aire() float64 {
	return r.Hauteur * r.Largeur
}
