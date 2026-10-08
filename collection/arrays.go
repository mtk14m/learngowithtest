package collection

func Somme(nombers []int) int {
	somme := 0
	for _, nombre := range nombers {
		somme += nombre
	}
	return somme
}

func SommeTout(listeDeNombres ...[]int) []int {
	resultat := make([]int, 0, len(listeDeNombres))
	for _, nombres := range listeDeNombres {
		somme := 0
		for _, nombre := range nombres {
			somme += nombre
		}
		resultat = append(resultat, somme)
	}
	return resultat
}

func sommeToutQueue(listeDeNombres ...[]int) []int {
	var sommes []int
	for _, nombres := range listeDeNombres {
		if len(nombres) == 0 {
			sommes = append(sommes, 0)
		} else {
			queue := nombres[1:]
			sommes = append(sommes, Somme(queue))
		}
	}
	return sommes
}
