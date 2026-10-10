package structure

import "math"

type Rectangle struct {
	Largeur float64
	Hauteur float64
}

type Cercle struct {
	Rayon float64
}

type Triangle struct {
	Base    float64
	Hauteur float64
}

type Forme interface {
	Aire() float64
}

func Perimetre(r Rectangle) float64 {
	return 2 * (r.Hauteur + r.Largeur)
}

func (c Cercle) Aire() float64 {
	return math.Pi * c.Rayon * c.Rayon
}

func (r Rectangle) Aire() float64 {
	return r.Hauteur * r.Largeur
}

func (t Triangle) Aire() float64 {
	return (t.Base * t.Hauteur) * 0.5
}
