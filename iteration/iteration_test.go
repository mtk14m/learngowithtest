package iteration

import (
	"fmt"
	"testing"
)

func TestRepeter(t *testing.T) {
	repete := Repeter("a", 5)
	attendu := "aaaaa"

	if repete != attendu {
		t.Errorf("attendu %q mais reçu %q", attendu, repete)
	}
}

func BenchmarkRepeter(b *testing.B) {
	for b.Loop() {
		Repeter("a", 5)
	}
}

func ExampleRepeter() {
	chaine := Repeter("a", 5)
	fmt.Println(chaine)
	// Output: aaaaa
}
