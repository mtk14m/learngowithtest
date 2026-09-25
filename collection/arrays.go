package collection

func Somme(nombers []int) int {
	somme := 0
	for _, nombre := range nombers {
		somme += nombre
	}
	return somme
}
