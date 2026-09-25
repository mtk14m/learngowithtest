package collection

func Somme(nombers [5]int) int {
	somme := 0
	for _, nombre := range nombers {
		somme += nombre
	}
	return somme
}
