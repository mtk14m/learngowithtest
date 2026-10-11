package pointeur

import "testing"

func TestPortefeuille(t *testing.T) {

	verifieSolde := func(t testing.TB, portefeuille Portefeuille, attendu Bitcoin) {
		t.Helper()
		solde := portefeuille.Solde()

		if solde != attendu {
			t.Errorf("solde %s, attendu %s", solde, attendu)
		}
	}

	verifieErreur := func(t testing.TB, erreur error, attendu error) {
		t.Helper()
		if erreur != attendu {
			t.Errorf("erreur attendu %q, reçu %q", attendu, erreur)
		}
	}

	t.Run("Deposer", func(t *testing.T) {
		portefeuille := Portefeuille{}
		portefeuille.Deposer(10)
		attendu := Bitcoin(10)
		verifieSolde(t, portefeuille, attendu)
	})

	t.Run("Retirer", func(t *testing.T) {
		portefeuille := Portefeuille{}
		portefeuille.Deposer(20)
		portefeuille.Retirer(10)
		attendu := Bitcoin(10)
		verifieSolde(t, portefeuille, attendu)
	})

	t.Run("Retirer avec des fonds insuffisants", func(t *testing.T) {
		soldeInitial := Bitcoin(20)
		portefeuille := Portefeuille{soldeInitial}
		erreur := portefeuille.Retirer(Bitcoin(100))

		verifieSolde(t, portefeuille, soldeInitial)
		verifieErreur(t, erreur, ErrSoldeInsuffisant)
	})
}
