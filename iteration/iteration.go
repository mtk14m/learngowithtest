package iteration

import "strings"

func Repeter(char string, nombreDeRepetiton int) string {
	var repete strings.Builder
	for i := 0; i < nombreDeRepetiton; i++ {
		repete.WriteString(char)
	}
	return repete.String()
}
